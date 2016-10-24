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

type Int32 struct {
	Value int32
	Valid bool
}

func NoneInt32() Int32 {
	return Int32{}
}

func SomeInt32(v int32) Int32 {
	return Int32{Value: v, Valid: true}
}

func (o Int32) Ptr() *int32 {
	if !o.Valid {
		return nil
	}
	return &o.Value
}

func (o Int32) String() string {
	if !o.Valid {
		return NONE
	}
	return fmt.Sprintf("%d", o.Value)
}

func (o Int32) GoString() string {
	return fmt.Sprintf("option.Int32(%s)", o.String())
}

func (o Int32) interfacePtr() interface{} {
	return interface{}(o.Ptr())
}
