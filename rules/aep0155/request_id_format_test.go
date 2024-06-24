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

package aep0155

import (
	"testing"

	"github.com/googleapis/api-linter/rules/internal/testutils"
)

func TestRequestIdFormat(t *testing.T) {
	for _, test := range []struct {
		name, FieldName, Type string
		problems                          testutils.Problems
	}{
		{
			name:       "ValidRequestIdFormat",
			FieldName:  "request_id",
			Type:       "aep.api.IdempotencyKey",
		},
		{
			name:      "SkipNonRequestId",
			FieldName: "other",
			Type:      "string",
		},
		{
			name:      "InvalidType",
			FieldName: "request_id",
			Type:      "string",
			problems:  testutils.Problems{{Message: "format = UUID4"}},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			f := testutils.ParseProto3Tmpl(t, `
				import "aep/api/idempotency_key.proto";
				message Person {
					{{.Type}} {{.FieldName}} = 2;
				}
				message Foo {}
			`, test)
			field := f.GetMessageTypes()[0].GetFields()[0]
			if diff := test.problems.SetDescriptor(field).Diff(requestIdFormat.Lint(f)); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
