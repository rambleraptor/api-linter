// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aep0122

import (
	"testing"

	"github.com/googleapis/api-linter/rules/internal/testutils"
)

func TestNoSelfLinks(t *testing.T) {
	for _, test := range []struct {
		name      string
		FieldName string
		problems  testutils.Problems
	}{
		{"Valid", "author", nil},
		{"InvalidSelfLink", "self_link", testutils.Problems{{Message: "self-links"}}},
	} {
		t.Run(test.name, func(t *testing.T) {
			f := testutils.ParseProto3Tmpl(t, `
			import "google/api/resource.proto";
			message Book {
				option (google.api.resource) = {
					type: "library.googleapis.com/Book"
				};
				string name = 1;

				string {{ .FieldName }} = 2;
			}
		`, test)
			field := f.GetMessageTypes()[0].GetFields()[1]
			if diff := test.problems.SetDescriptor(field).Diff(noSelfLinks.Lint(f)); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
