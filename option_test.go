/*
 Copyright 2064 Rodrigo Rafael Monti Kochenburger

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

package option

import (
	"fmt"
	"reflect"
	"testing"
)

type iface interface {
	String() string
	GoString() string
	IsPresent() bool
	interfaceValue() interface{}
}

type testEntry struct {
	Object   iface
	Value    interface{}
	TypeName string
}

func verifyTestEntry(t *testing.T, idx int, e testEntry) {
	verifyPtr(t, idx, e)
	verifyString(t, idx, e)
	verifyGoString(t, idx, e)
}

func verifyPtr(t *testing.T, idx int, e testEntry) {
	if e.Value == nil {
		if e.Object.IsPresent() {
			t.Errorf("Entry %d: expected object to not be present", idx)
		}
	} else {
		v := e.Object.interfaceValue()

		if !reflect.DeepEqual(v, e.Value) {
			t.Errorf("Entry %d: expected .Value() to return %#v, got %#v instead", idx, e.Value, v)
		}
	}
}

func verifyString(t *testing.T, idx int, e testEntry) {
	str := e.Object.String()
	if e.Value == nil {
		if str != NONE {
			t.Errorf("Entry %d: expected .String() to return %v, got %s instead", idx, NONE, str)
		}
	} else {
		if str != fmt.Sprintf("%v", e.Value) {
			t.Errorf("Entry %d: expected .String() to return %v, got %s instead", idx, e.Value, str)
		}
	}
}

func verifyGoString(t *testing.T, idx int, e testEntry) {
	goStr := e.Object.GoString()
	if e.Value == nil {
		exp := fmt.Sprintf("%s(%s)", e.TypeName, NONE)
		if goStr != exp {
			t.Errorf("Entry %d: expected .GoString() to return %#v, got %#v instead", idx, exp, goStr)
		}
	} else {
		exp := fmt.Sprintf("%s(%#v)", e.TypeName, e.Value)
		if goStr != exp {
			t.Errorf("Entry %d: expected .GoString() to return %#v, got %#v instead", idx, exp, goStr)
		}
	}
}
