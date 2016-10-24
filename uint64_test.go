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

func TestUint64(t *testing.T) {
	table := []testEntry{
		{NoneUint64(), nil, "option.Uint64"},
		{SomeUint64(0), uint64(0), "option.Uint64"},
		{SomeUint64(math.MaxUint64), uint64(math.MaxUint64), "option.Uint64"},
	}

	for idx, e := range table {
		verifyTestEntry(t, idx, e)
	}
}
