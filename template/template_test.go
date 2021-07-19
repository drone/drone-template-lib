// Copyright 2018 Drone.IO Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package template

import (
	"testing"
)

const template = `
<p>{{ safeRegexReplace "\n" hello_world "<br>" }}</p>
<p>{{ htmlLineBreaks lorem_ipsum }}</p>
`

const expected = `
<p>hello<br>world</p>
<p>lorem ipsum<br>sit dolor<br></p>
`

func TestRender(t *testing.T) {
	ctx := map[string]string{
		"hello_world": "hello\nworld",
		"lorem_ipsum": "lorem ipsum\nsit dolor\n",
	}

	actual, err := Render(template, ctx)
	if err != nil {
		t.Errorf("error, %v", err)
	}
	if actual != expected {
		t.Errorf("error, expected %s, got %s", expected, actual)
	}
}
