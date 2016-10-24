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
	"errors"
	"fmt"
	"math"
	"reflect"
	"testing"
)

const tableLineOffset = 22

var table = []struct {
	Object   Interface
	Value    interface{}
	TypeName string
}{
	{NoneInt(), nil, "option.Int"},
	{SomeInt(0), int(0), "option.Int"},
	{NoneInt8(), nil, "option.Int8"},
	{SomeInt8(0), int8(0), "option.Int8"},
	{SomeInt8(math.MaxInt8), int8(math.MaxInt8), "option.Int8"},
	{SomeInt8(math.MinInt8), int8(math.MinInt8), "option.Int8"},
	{NoneInt16(), nil, "option.Int16"},
	{SomeInt16(0), int16(0), "option.Int16"},
	{SomeInt16(math.MaxInt16), int16(math.MaxInt16), "option.Int16"},
	{SomeInt16(math.MinInt16), int16(math.MinInt16), "option.Int16"},
	{NoneInt32(), nil, "option.Int32"},
	{SomeInt32(0), int32(0), "option.Int32"},
	{SomeInt32(math.MaxInt32), int32(math.MaxInt32), "option.Int32"},
	{SomeInt32(math.MinInt32), int32(math.MinInt32), "option.Int32"},
	{NoneInt64(), nil, "option.Int64"},
	{SomeInt64(0), int64(0), "option.Int64"},
	{SomeInt64(math.MaxInt64), int64(math.MaxInt64), "option.Int64"},
	{SomeInt64(math.MinInt64), int64(math.MinInt64), "option.Int64"},
	{NoneUint(), nil, "option.Uint"},
	{SomeUint(0), uint(0), "option.Uint"},
	{NoneUint8(), nil, "option.Uint8"},
	{SomeUint8(0), uint8(0), "option.Uint8"},
	{SomeUint8(math.MaxUint8), uint8(math.MaxUint8), "option.Uint8"},
	{NoneUint16(), nil, "option.Uint16"},
	{SomeUint16(0), uint16(0), "option.Uint16"},
	{SomeUint16(math.MaxUint16), uint16(math.MaxUint16), "option.Uint16"},
	{NoneUint32(), nil, "option.Uint32"},
	{SomeUint32(0), uint32(0), "option.Uint32"},
	{SomeUint32(math.MaxUint32), uint32(math.MaxUint32), "option.Uint32"},
	{NoneUint64(), nil, "option.Uint64"},
	{SomeUint64(0), uint64(0), "option.Uint64"},
	{SomeUint64(math.MaxUint64), uint64(math.MaxUint64), "option.Uint64"},
	{NoneFloat32(), nil, "option.Float32"},
	{SomeFloat32(0), float32(0), "option.Float32"},
	{SomeFloat32(math.MaxFloat32), float32(math.MaxFloat32), "option.Float32"},
	{SomeFloat32(math.SmallestNonzeroFloat32), float32(math.SmallestNonzeroFloat32), "option.Float32"},
	{NoneFloat64(), nil, "option.Float64"},
	{SomeFloat64(0), float64(0), "option.Float64"},
	{SomeFloat64(math.MaxFloat64), float64(math.MaxFloat64), "option.Float64"},
	{SomeFloat64(math.SmallestNonzeroFloat64), float64(math.SmallestNonzeroFloat64), "option.Float64"},
	{NoneBool(), nil, "option.Bool"},
	{SomeBool(true), true, "option.Bool"},
	{SomeBool(false), false, "option.Bool"},
	{NoneString(), nil, "option.String"},
	{SomeString(""), "", "option.String"},
	{SomeString("something here"), "something here", "option.String"},
	{NoneError(), nil, "option.Error"},
	{SomeError(errors.New("")), errors.New(""), "option.Error"},
	{SomeError(errors.New("something here")), errors.New("something here"), "option.Error"},
}

func TestTableAsInterfaces(t *testing.T) {
	for idx, e := range table {
		line := tableLineOffset + idx

		var iptrValue interface{}
		iptr := reflect.ValueOf(e.Object.interfacePtr())

		if iptr.IsNil() {
			iptrValue = nil
		} else {
			iptrValue = iptr.Elem().Interface()
		}

		if !reflect.DeepEqual(iptrValue, e.Value) {
			t.Errorf("Line %d: expected .interfacePtr() to return %#v, got %#v instead", line, e.Value, iptrValue)
		}

		str := e.Object.String()
		if iptrValue == nil {
			if str != NONE {
				t.Errorf("Line %d: expected .String() to return %v, got %s instead", line, NONE, str)
			}
		} else {
			if str != fmt.Sprintf("%v", iptrValue) {
				t.Errorf("Line %d: expected .String() to return %v, got %s instead", line, iptrValue, str)
			}
		}

		goStr := e.Object.GoString()
		if iptrValue == nil {
			exp := fmt.Sprintf("%s(%s)", e.TypeName, NONE)
			if goStr != exp {
				t.Errorf("Line %d: expected .GoString() to return %#v, got %#v instead", line, exp, goStr)
			}
		} else {
			exp := fmt.Sprintf("%s(%#v)", e.TypeName, iptrValue)
			if goStr != exp {
				t.Errorf("Line %d: expected .GoString() to return %#v, got %#v instead", line, exp, goStr)
			}
		}
	}
}
