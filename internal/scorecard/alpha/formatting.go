// Copyright 2020 The Operator-SDK Authors
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

package alpha

import (
	"context"
	"encoding/json"

	"github.com/operator-framework/operator-sdk/pkg/apis/scorecard/v1alpha3"
	v1 "k8s.io/api/core/v1"
)

// getTestResult fetches the test pod log and converts it into
// Test format
func (r PodTestRunner) getTestResult(ctx context.Context, p *v1.Pod, test Test) (output *v1alpha3.TestResult) {

	logBytes, err := getPodLog(ctx, r.Client, p)
	if err != nil {
		return testResultError(err, test)
	}
	// marshal pod log into TestResult
	err = json.Unmarshal(logBytes, &output)
	if err != nil {
		return testResultError(err, test)
	}
	return output
}

// ListTests lists the scorecard tests as configured that would be
// run based on user selection
func (o Scorecard) ListTests() (output v1alpha3.Test, err error) {
	tests := o.selectTests()
	if len(tests) == 0 {
		output.Status.Results = make([]v1alpha3.TestResult, 0)
		return output, err
	}

	output.Status.Results = make([]v1alpha3.TestResult, 0)

	for _, test := range tests {
		testResult := v1alpha3.TestResult{}
		testResult.Name = test.Name
		output.Spec.Labels = test.Labels
		output.Status.Results = append(output.Status.Results, testResult)
	}

	return output, err
}

func testResultError(err error, test Test) *v1alpha3.TestResult {
	r := v1alpha3.TestResult{}
	r.Name = test.Name
	r.State = v1alpha3.FailState
	r.Errors = []string{err.Error()}
	return &r
}
