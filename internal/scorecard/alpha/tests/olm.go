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

package tests

import (
	"bytes"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"github.com/operator-framework/api/pkg/manifests"
	"github.com/operator-framework/operator-registry/pkg/lib/bundle"
	"github.com/operator-framework/operator-registry/pkg/registry"
	scapiv1alpha2 "github.com/operator-framework/operator-sdk/pkg/apis/scorecard/v1alpha2"
)

const (
	OLMBundleValidationTest   = "olm-bundle-validation"
	OLMCRDsHaveValidationTest = "olm-crds-have-validation"
	OLMCRDsHaveResourcesTest  = "olm-crds-have-resources"
	OLMSpecDescriptorsTest    = "olm-spec-descriptors"
	OLMStatusDescriptorsTest  = "olm-status-descriptors"
)

// BundleValidationTest validates an on-disk bundle
func BundleValidationTest(dir string) scapiv1alpha2.ScorecardTestResult {
	r := scapiv1alpha2.ScorecardTestResult{}
	r.Name = OLMBundleValidationTest
	r.Description = "Validates bundle contents"
	r.State = scapiv1alpha2.PassState
	r.Errors = []string{}
	r.Suggestions = []string{}

	defaultOutput := logrus.StandardLogger().Out
	defer logrus.SetOutput(defaultOutput)

	// Log output from the test will be captured in this buffer
	buf := &bytes.Buffer{}
	logger := logrus.WithField("name", "bundle-test")
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(buf)

	val := bundle.NewImageValidator("", logger)

	// Validate bundle format.
	if err := val.ValidateBundleFormat(dir); err != nil {
		r.State = scapiv1alpha2.FailState
		r.Errors = append(r.Errors, err.Error())
	}

	// Validate bundle content.
	manifestsDir := filepath.Join(dir, bundle.ManifestsDir)
	_, _, validationResults := manifests.GetManifestsDir(dir)
	for _, result := range validationResults {
		for _, e := range result.Errors {
			r.Errors = append(r.Errors, e.Error())
			r.State = scapiv1alpha2.FailState
		}

		for _, w := range result.Warnings {
			r.Suggestions = append(r.Suggestions, w.Error())
		}
	}

	if err := val.ValidateBundleContent(manifestsDir); err != nil {
		r.State = scapiv1alpha2.FailState
		r.Errors = append(r.Errors, err.Error())
	}

	r.Log = buf.String()
	return r
}

// CRDsHaveValidationTest verifies all CRDs have a validation section
func CRDsHaveValidationTest(bundle registry.Bundle) scapiv1alpha2.ScorecardTestResult {
	r := scapiv1alpha2.ScorecardTestResult{}
	r.Name = OLMCRDsHaveValidationTest
	r.Description = "All CRDs have an OpenAPI validation subsection"
	r.State = scapiv1alpha2.PassState
	r.Errors = make([]string, 0)
	r.Suggestions = make([]string, 0)
	return r
}

// CRDsHaveResourcesTest verifies CRDs have resources listed in its owned CRDs section
func CRDsHaveResourcesTest(bundle registry.Bundle) scapiv1alpha2.ScorecardTestResult {
	r := scapiv1alpha2.ScorecardTestResult{}
	r.Name = OLMCRDsHaveResourcesTest
	r.Description = "All Owned CRDs contain a resources subsection"
	r.State = scapiv1alpha2.PassState
	r.Errors = make([]string, 0)
	r.Suggestions = make([]string, 0)

	return r
}

// SpecDescriptorsTest verifies all spec fields have descriptors
func SpecDescriptorsTest(bundle registry.Bundle) scapiv1alpha2.ScorecardTestResult {
	r := scapiv1alpha2.ScorecardTestResult{}
	r.Name = OLMSpecDescriptorsTest
	r.Description = "All spec fields have matching descriptors in the CSV"
	r.State = scapiv1alpha2.PassState
	r.Errors = make([]string, 0)
	r.Suggestions = make([]string, 0)
	return r
}

// StatusDescriptorsTest verifies all CRDs have status descriptors
func StatusDescriptorsTest(bundle registry.Bundle) scapiv1alpha2.ScorecardTestResult {
	r := scapiv1alpha2.ScorecardTestResult{}
	r.Name = OLMStatusDescriptorsTest
	r.Description = "All status fields have matching descriptors in the CSV"
	r.State = scapiv1alpha2.PassState
	r.Errors = make([]string, 0)
	r.Suggestions = make([]string, 0)
	return r
}
