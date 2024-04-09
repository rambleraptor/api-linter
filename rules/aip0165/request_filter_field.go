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

package aep0165

import (
	"github.com/googleapis/api-linter/lint"
	"github.com/googleapis/api-linter/rules/internal/utils"
)

// The Purge request message should have filter field.
var requestFilterField = &lint.MessageRule{
	Name:        lint.NewRuleName(165, "request-filter-field"),
	OnlyIf:      isPurgeRequestMessage,
	LintMessage: utils.LintFieldPresentAndSingularString("filter"),
}
