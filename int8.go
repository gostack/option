/*
 Copyright 2016 Rodrigo Rafael Monti Kochenburger

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

type Int8 struct {
	Value int8
	Valid bool
}

func NoneInt8() Int8 {
	return Int8{}
}

func SomeInt8(v int8) Int8 {
	return Int8{Value: v, Valid: true}
}

func (o Int8) Ptr() *int8 {
	if !o.Valid {
		return nil
	}
	return &o.Value
}

func (o Int8) String() string {
	if !o.Valid {
		return NONE
	}
	return fmt.Sprintf("%d", o.Value)
}

func (o Int8) GoString() string {
	return fmt.Sprintf("option.Int8(%s)", o.String())
}

func (o Int8) interfacePtr() interface{} {
	return interface{}(o.Ptr())
}
