// Copyright 2017 Cisco Systems, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hostagent

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func doTestOpflex(t *testing.T, agent *testHostAgent) {
	tempdir, err := ioutil.TempDir("", "hostagent_test_")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tempdir)

	agent.config.OpFlexConfigPath = tempdir

	agent.writeOpflexConfig()
	_, err = os.Stat(filepath.Join(tempdir, "01-base.conf"))
	assert.Nil(t, err, "base")
	_, err = os.Stat(filepath.Join(tempdir, "10-renderer.conf"))
	assert.Nil(t, err, "renderer")
}

func TestOpflexConfigVxlan(t *testing.T) {
	agent := testAgent()
	agent.config = &HostAgentConfig{
		HostAgentNodeConfig: HostAgentNodeConfig{
			VxlanIface:  "eth1.4093",
			UplinkIface: "eth1",
		},
		EncapType:    "vxlan",
		AciInfraVlan: 4093,
		NodeName:     "node1",
	}
	doTestOpflex(t, agent)
}

func TestOpflexConfigVlan(t *testing.T) {
	agent := testAgent()
	agent.config = &HostAgentConfig{
		HostAgentNodeConfig: HostAgentNodeConfig{
			VxlanIface:  "eth1.4093",
			UplinkIface: "eth1",
		},
		EncapType:    "vlan",
		AciInfraVlan: 4093,
		NodeName:     "node1",
	}
	doTestOpflex(t, agent)
}
