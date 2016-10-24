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

type Bool struct {
	Value bool
	Valid bool
}

func NoneBool() Bool {
	return Bool{}
}

func SomeBool(v bool) Bool {
	return Bool{Value: v, Valid: true}
}

func (o Bool) Ptr() *bool {
	if !o.Valid {
		return nil
	}
	return &o.Value
}

func (o Bool) String() string {
	if !o.Valid {
		return NONE
	}

	return fmt.Sprintf("%v", o.Value)
}

func (o Bool) GoString() string {
	if !o.Valid {
		return fmt.Sprintf("option.Bool(%s)", NONE)
	}

	return fmt.Sprintf("option.Bool(%#v)", o.Value)
}

func (o Bool) interfacePtr() interface{} {
	return interface{}(o.Ptr())
}
