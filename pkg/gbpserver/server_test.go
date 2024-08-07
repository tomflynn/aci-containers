/***
Copyright 2018 Cisco Systems Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gbpserver

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestLoginHandlerServeHTTP(t *testing.T) {
	handler := &loginHandler{}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", http.NoBody)
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}
func TestNfhServeHTTP(t *testing.T) {
	n := &nfh{}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", http.NoBody)
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	recorder := httptest.NewRecorder()

	n.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}
func TestUTReadMsg(t *testing.T) {
	s := &Server{
		rxCh: make(chan *inputMsg),
	}

	op, data, err := s.UTReadMsg(time.Second)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != "timeout" {
		t.Errorf("Expected error 'timeout', got '%v'", err)
	}
	if op != 0 {
		t.Errorf("Expected op 0, got %d", op)
	}
	if data != nil {
		t.Errorf("Expected data nil, got %v", data)
	}
}
func TestAddNetPol(t *testing.T) {
	s := &Server{
		rxCh: make(chan *inputMsg),
	}

	np := NetworkPolicy{}

	expectedMsg := &inputMsg{
		op:   OpaddNetPol,
		data: &np,
	}

	go s.AddNetPol(np)

	select {
	case msg := <-s.rxCh:
		if !reflect.DeepEqual(msg, expectedMsg) {
			t.Errorf("Expected message %v, got %v", expectedMsg, msg)
		}
	case <-time.After(time.Second):
		t.Errorf("Timeout waiting for message")
	}
}
func TestDelEPServeHTTP(t *testing.T) {
	s := &Server{
		rxCh: make(chan *inputMsg),
	}

	ep := Endpoint{}

	expectedMsg := &inputMsg{
		op:   OpdelEP,
		data: &ep,
	}

	go s.DelEP(ep)

	select {
	case msg := <-s.rxCh:
		if !reflect.DeepEqual(msg, expectedMsg) {
			t.Errorf("Expected message %v, got %v", expectedMsg, msg)
		}
	case <-time.After(time.Second):
		t.Errorf("Timeout waiting for message")
	}
}
