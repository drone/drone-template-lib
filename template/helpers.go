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
	"net/url"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/aymerick/raymond"
)

var (
	funcs = map[string]interface{}{
		"duration":       toDuration,
		"datetime":       toDatetime,
		"success":        isSuccess,
		"failure":        isFailure,
		"truncate":       truncate,
		"urlencode":      urlencode,
		"since":          since,
		"uppercasefirst": uppercaseFirst,
		"uppercase":      strings.ToUpper,
		"lowercase":      strings.ToLower,
		"trim":           strings.TrimSpace,
		"title":          strings.Title,
	}
)

func init() {
	raymond.RegisterHelpers(funcs)
}

func toDuration(started, finished int64) string {
	return fmt.Sprint(time.Duration(finished-started) * time.Second)
}

func toDatetime(timestamp int64, layout, zone string) string {
	if len(zone) == 0 {
		return time.Unix(timestamp, 0).Format(layout)
	}

	loc, err := time.LoadLocation(zone)

	if err != nil {
		return time.Unix(timestamp, 0).Local().Format(layout)
	}

	return time.Unix(timestamp, 0).In(loc).Format(layout)
}

func isSuccess(conditional bool, options *raymond.Options) string {
	if !conditional {
		return options.Inverse()
	}

	switch options.ParamStr(0) {
	case "success":
		return options.Fn()
	default:
		return options.Inverse()
	}
}

func isFailure(conditional bool, options *raymond.Options) string {
	if !conditional {
		return options.Inverse()
	}

	switch options.ParamStr(0) {
	case "failure", "error", "killed":
		return options.Fn()
	default:
		return options.Inverse()
	}
}

func truncate(s string, len int) string {
	if utf8.RuneCountInString(s) <= len {
		return s
	}

	runes := []rune(s)

	return string(runes[:len])
}

func urlencode(options *raymond.Options) string {
	return url.QueryEscape(options.Fn())
}

func since(start int64) string {
	now := time.Unix(time.Now().Unix(), 0)
	return fmt.Sprint(now.Sub(time.Unix(start, 0)))
}

func uppercaseFirst(s string) string {
	a := []rune(s)

	a[0] = unicode.ToUpper(a[0])
	s = string(a)

	return s
}
