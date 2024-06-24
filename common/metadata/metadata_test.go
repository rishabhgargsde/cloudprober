// Copyright 2023 The Cloudprober Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package metadata implements metadata related utilities.
*/
package metadata

import (
	"regexp"
	"testing"

	"google3/third_party/golang/testify/assert/assert"
)

func TestUniqueID(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "firstRun"},
		{name: "fromCache"},
	}

	re := regexp.MustCompile("cloudprober-[a-z]{6}-[0-9]{4}")
	var last string
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniqueID()
			if last != "" {
				assert.Equal(t, last, got, "Unique id not same as last time")
				return
			}

			assert.True(t, re.MatchString(got), "Unique id doesn't match the regex")

			last = got
		})
	}
}
