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

func TestInt32(t *testing.T) {
	table := []testEntry{
		{NoneInt32(), nil, "option.Int32"},
		{SomeInt32(0), int32(0), "option.Int32"},
		{SomeInt32(math.MaxInt32), int32(math.MaxInt32), "option.Int32"},
		{SomeInt32(math.MinInt32), int32(math.MinInt32), "option.Int32"},
	}

	for idx, e := range table {
		verifyPtr(t, idx, e)
		verifyString(t, idx, e)
		verifyGoString(t, idx, e)
	}
}
