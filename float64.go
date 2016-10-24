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
)

type Float64 struct {
	Value float64
	Valid bool
}

func NoneFloat64() Float64 {
	return Float64{}
}

func SomeFloat64(v float64) Float64 {
	return Float64{Value: v, Valid: true}
}

func (o Float64) Ptr() *float64 {
	if !o.Valid {
		return nil
	}
	return &o.Value
}

func (o Float64) String() string {
	if !o.Valid {
		return NONE
	}

	return fmt.Sprintf("%v", o.Value)
}

func (o Float64) GoString() string {
	if !o.Valid {
		return fmt.Sprintf("option.Float64(%s)", NONE)
	}

	return fmt.Sprintf("option.Float64(%#v)", o.Value)
}

func (o Float64) interfacePtr() interface{} {
	return interface{}(o.Ptr())
}
