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
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/aymerick/raymond"
)

// Render parses and executes a template, returning the results in string
// format. Trailing or leading spaces or new-lines are not getting truncated. It
// is able to read templates from remote paths, local files or directly from the
// string.
func Render(template string, payload interface{}) (s string, err error) {
	u, err := url.Parse(template)

	if err == nil {
		switch u.Scheme {
		case "http", "https":
			res, err := http.Get(template)

			if err != nil {
				return s, fmt.Errorf("failed to fetch: %w", err)
			}

			defer res.Body.Close()

			out, err := ioutil.ReadAll(res.Body)

			if err != nil {
				return s, fmt.Errorf("failed to read: %w", err)
			}

			template = string(out)
		case "file":
			out, err := ioutil.ReadFile(u.Path)

			if err != nil {
				return s, fmt.Errorf("failed to read: %w", err)
			}

			template = string(out)
		}
	}

	return raymond.Render(template, payload)
}

// RenderTrim parses and executes a template, returning the results in string
// format. The result is trimmed to remove left and right padding and newlines
// that may be added unintentially in the template markup.
func RenderTrim(template string, playload interface{}) (string, error) {
	out, err := Render(template, playload)
	return strings.Trim(out, " \n"), err
}
