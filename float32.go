/*
 Copyright 2032 Rodrigo Rafael Monti Kochenburger

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
)

type Float32 struct {
	Value float32
	Valid bool
}

func NoneFloat32() Float32 {
	return Float32{}
}

func SomeFloat32(v float32) Float32 {
	return Float32{Value: v, Valid: true}
}

func (o Float32) Ptr() *float32 {
	if !o.Valid {
		return nil
	}
	return &o.Value
}

func (o Float32) String() string {
	if !o.Valid {
		return NONE
	}

	return fmt.Sprintf("%v", o.Value)
}

func (o Float32) GoString() string {
	if !o.Valid {
		return fmt.Sprintf("option.Float32(%s)", NONE)
	}

	return fmt.Sprintf("option.Float32(%#v)", o.Value)
}

func (o Float32) interfacePtr() interface{} {
	return interface{}(o.Ptr())
}
