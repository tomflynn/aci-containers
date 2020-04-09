// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package packagestest

import (
	"archive/zip"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/tools/internal/gocommand"
	"golang.org/x/tools/internal/packagesinternal"
)

// Modules is the exporter that produces module layouts.
// Each "repository" is put in it's own module, and the module file generated
// will have replace directives for all other modules.
// Given the two files
//     golang.org/repoa#a/a.go
//     golang.org/repob#b/b.go
// You would get the directory layout
//     /sometemporarydirectory
//     ├── repoa
//     │   ├── a
//     │   │   └── a.go
//     │   └── go.mod
//     └── repob
//         ├── b
//         │   └── b.go
//         └── go.mod
// and the working directory would be
//     /sometemporarydirectory/repoa
var Modules = modules{}

type modules struct{}

type moduleAtVersion struct {
	module  string
	version string
}

func (modules) Name() string {
	return "Modules"
}

func (modules) Filename(exported *Exported, module, fragment string) string {
	if module == exported.primary {
		return filepath.Join(primaryDir(exported), fragment)
	}
	return filepath.Join(moduleDir(exported, module), fragment)
}

func (modules) Finalize(exported *Exported) error {
	// Write out the primary module. This module can use symlinks and
	// other weird stuff, and will be the working dir for the go command.
	// It depends on all the other modules.
	primaryDir := primaryDir(exported)
	if err := os.MkdirAll(primaryDir, 0755); err != nil {
		return err
	}
	exported.Config.Dir = primaryDir
	if exported.written[exported.primary] == nil {
		exported.written[exported.primary] = make(map[string]string)
	}

	// Create a map of modulepath -> {module, version} for modulepaths
	// that are of the form `repoa/mod1@v1.1.0`.
	versions := make(map[string]moduleAtVersion)
	for module := range exported.written {
		if splt := strings.Split(module, "@"); len(splt) > 1 {
			versions[module] = moduleAtVersion{
				module:  splt[0],
				version: splt[1],
			}
		}
	}

	// If the primary module already has a go.mod, write the contents to a temp
	// go.mod for now and then we will reset it when we are getting all the markers.
	if gomod := exported.written[exported.primary]["go.mod"]; gomod != "" {
		contents, err := ioutil.ReadFile(gomod)
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(gomod+".temp", contents, 0644); err != nil {
			return err
		}
	}

	exported.written[exported.primary]["go.mod"] = filepath.Join(primaryDir, "go.mod")
	primaryGomod := "module " + exported.primary + "\nrequire (\n"
	for other := range exported.written {
		if other == exported.primary {
			continue
		}
		version := moduleVersion(other)
		// If other is of the form `repo1/mod1@v1.1.0`,
		// then we need to extract the module and the version.
		if v, ok := versions[other]; ok {
			other = v.module
			version = v.version
		}
		primaryGomod += fmt.Sprintf("\t%v %v\n", other, version)
	}
	primaryGomod += ")\n"
	if err := ioutil.WriteFile(filepath.Join(primaryDir, "go.mod"), []byte(primaryGomod), 0644); err != nil {
		return err
	}

	// Create the mod cache so we can rename it later, even if we don't need it.
	if err := os.MkdirAll(modCache(exported), 0755); err != nil {
		return err
	}

	// Write out the go.mod files for the other modules.
	for module, files := range exported.written {
		if module == exported.primary {
			continue
		}
		dir := moduleDir(exported, module)
		modfile := filepath.Join(dir, "go.mod")
		// If other is of the form `repo1/mod1@v1.1.0`,
		// then we need to extract the module name without the version.
		if v, ok := versions[module]; ok {
			module = v.module
		}
		if err := ioutil.WriteFile(modfile, []byte("module "+module+"\n"), 0644); err != nil {
			return err
		}
		files["go.mod"] = modfile
	}

	// Zip up all the secondary modules into the proxy dir.
	proxyDir := filepath.Join(exported.temp, "modproxy")
	for module, files := range exported.written {
		if module == exported.primary {
			continue
		}
		version := moduleVersion(module)
		// If other is of the form `repo1/mod1@v1.1.0`,
		// then we need to extract the module and the version.
		if v, ok := versions[module]; ok {
			module = v.module
			version = v.version
		}
		dir := filepath.Join(proxyDir, module, "@v")
		if err := writeModuleProxy(dir, module, version, files); err != nil {
			return fmt.Errorf("creating module proxy dir for %v: %v", module, err)
		}
	}

	// Discard the original mod cache dir, which contained the files written
	// for us by Export.
	if err := os.Rename(modCache(exported), modCache(exported)+".orig"); err != nil {
		return err
	}
	exported.Config.Env = append(exported.Config.Env,
		"GO111MODULE=on",
		"GOPATH="+filepath.Join(exported.temp, "modcache"),
		"GOPROXY="+proxyDirToURL(proxyDir),
		"GOSUMDB=off",
	)
	gocmdRunner := &gocommand.Runner{}
	packagesinternal.SetGoCmdRunner(exported.Config, gocmdRunner)

	// Run go mod download to recreate the mod cache dir with all the extra
	// stuff in cache. All the files created by Export should be recreated.
	inv := gocommand.Invocation{
		Verb:       "mod",
		Args:       []string{"download"},
		Env:        exported.Config.Env,
		BuildFlags: exported.Config.BuildFlags,
		WorkingDir: exported.Config.Dir,
	}
	if _, err := gocmdRunner.Run(context.Background(), inv); err != nil {
		return err
	}
	return nil
}

// writeModuleProxy creates a directory in the proxy dir for a module.
func writeModuleProxy(dir, module, ver string, files map[string]string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// the modproxy checks for versions by looking at the "list" file,
	// since we are supporting multiple versions, create the file if it does not exist or
	// append the version number to the preexisting file.
	f, err := os.OpenFile(filepath.Join(dir, "list"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := f.WriteString(ver + "\n"); err != nil {
		return err
	}

	// go.mod, copied from the file written in Finalize.
	modContents, err := ioutil.ReadFile(files["go.mod"])
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(dir, ver+".mod"), modContents, 0644); err != nil {
		return err
	}

	// info file, just the bare bones.
	infoContents := []byte(fmt.Sprintf(`{"Version": "%v", "Time":"2017-12-14T13:08:43Z"}`, ver))
	if err := ioutil.WriteFile(filepath.Join(dir, ver+".info"), infoContents, 0644); err != nil {
		return err
	}

	// zip of all the source files.
	f, err = os.OpenFile(filepath.Join(dir, ver+".zip"), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	z := zip.NewWriter(f)
	for name, path := range files {
		zf, err := z.Create(module + "@" + ver + "/" + name)
		if err != nil {
			return err
		}
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		if _, err := zf.Write(contents); err != nil {
			return err
		}
	}
	if err := z.Close(); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func modCache(exported *Exported) string {
	return filepath.Join(exported.temp, "modcache/pkg/mod")
}

func primaryDir(exported *Exported) string {
	return filepath.Join(exported.temp, path.Base(exported.primary))
}

func moduleDir(exported *Exported, module string) string {
	if strings.Contains(module, "@") {
		return filepath.Join(modCache(exported), module)
	}
	return filepath.Join(modCache(exported), path.Dir(module), path.Base(module)+"@"+moduleVersion(module))
}

var versionSuffixRE = regexp.MustCompile(`v\d+`)

func moduleVersion(module string) string {
	if versionSuffixRE.MatchString(path.Base(module)) {
		return path.Base(module) + ".0.0"
	}
	return "v1.0.0"
}
