// Copyright 2020 Google LLC
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

package aep0164

import (
	"fmt"
	"strings"

	"github.com/googleapis/api-linter/lint"
	"github.com/googleapis/api-linter/locations"
	"github.com/googleapis/api-linter/rules/internal/utils"
	"github.com/jhump/protoreflect/desc"
)

// Undelete messages should use google.longrunning.Operation
// or the resource itself as the response message.
var responseMessageName = &lint.MethodRule{
	Name:   lint.NewRuleName(164, "response-message-name"),
	OnlyIf: isUndeleteMethod,
	LintMethod: func(m *desc.MethodDescriptor) []lint.Problem {
		// Rule check: Establish that for methods such as `UndeleteFoo`, the response
		// message is `Foo` or `google.longrunning.Operation`.
		want := strings.Replace(m.GetName(), "Undelete", "", 1)
		got := m.GetOutputType().GetName()

		// If the return type is an LRO, use the annotated response type instead.
		if utils.IsOperation(m.GetOutputType()) {
			got = utils.GetOperationInfo(m).GetResponseType()
		}

		// Return a problem if we did not get the expected return name.
		//
		// Note: If `got` is empty string, this is an unannotated LRO.
		// The AEP-151 rule will whine about that, and this rule should not as it
		// would be confusing.
		if got != want && got != "" {
			return []lint.Problem{{
				Message: fmt.Sprintf(
					"Undelete RPCs should have response message type %q, not %q.",
					want,
					got,
				),
				Suggestion: want,
				Descriptor: m,
				Location:   locations.MethodResponseType(m),
			}}
		}
		return nil
	},
}
