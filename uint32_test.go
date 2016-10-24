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
	"math"
	"testing"
)

func TestUint32(t *testing.T) {
	table := []testEntry{
		{NoneUint32(), nil, "option.Uint32"},
		{SomeUint32(0), uint32(0), "option.Uint32"},
		{SomeUint32(math.MaxUint32), uint32(math.MaxUint32), "option.Uint32"},
	}

	for idx, e := range table {
		verifyPtr(t, idx, e)
		verifyString(t, idx, e)
		verifyGoString(t, idx, e)
	}
}
