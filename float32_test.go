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

func TestFloat32(t *testing.T) {
	table := []testEntry{
		{NoneFloat32(), nil, "option.Float32"},
		{SomeFloat32(0), float32(0), "option.Float32"},
		{SomeFloat32(math.MaxFloat32), float32(math.MaxFloat32), "option.Float32"},
		{SomeFloat32(math.SmallestNonzeroFloat32), float32(math.SmallestNonzeroFloat32), "option.Float32"},
	}

	for idx, e := range table {
		verifyPtr(t, idx, e)
		verifyString(t, idx, e)
		verifyGoString(t, idx, e)
	}
}
