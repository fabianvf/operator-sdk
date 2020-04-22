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
	"io/ioutil"

	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
)

type ScorecardTest struct {
	Name        string            `yaml:"name"`                 // The container test name
	Image       string            `yaml:"image"`                // The container image name
	Entrypoint  string            `yaml:"entrypoint,omitempty"` // An optional entrypoint passed to the test image
	Labels      map[string]string `yaml:"labels"`               // User defined labels used to filter tests
	Description string            `yaml:"description"`          // User readable test description
	TestPod     *v1.Pod           `yaml:"-"`                    // Pod that ran the test
}

// Config represents the set of test configurations which scorecard
// would run based on user input
type Config struct {
	Tests []ScorecardTest `yaml:"tests"`
}

// LoadConfig will find and return the scorecard config, the config file
// can be passed in via command line flag or from a bundle location or
// bundle image
func LoadConfig(configFilePath string) (Config, error) {
	c := Config{}

	// TODO handle getting config from bundle (ondisk or image)
	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return c, err
	}

	if err := yaml.Unmarshal(yamlFile, &c); err != nil {
		return c, err
	}

	return c, nil
}
