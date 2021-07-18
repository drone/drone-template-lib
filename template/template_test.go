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

func TestRender(t *testing.T) {
	template := `<p>{{ safeRegexReplace "\n" content "<br>" }}</p>`
	ctx := map[string]string{
    "content": "hello\nworld",
	}

	expected  := "<p>hello<br>world</p>"
	actual, _ := Render(template, ctx)
	if actual != expected {
		t.Errorf("error, expected %s, got %s", expected, actual)
	}
}
