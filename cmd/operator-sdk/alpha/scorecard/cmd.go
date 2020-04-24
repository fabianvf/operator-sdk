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

package scorecard

import (
	"fmt"

	"time"

	scorecard "github.com/operator-framework/operator-sdk/internal/scorecard/alpha"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/labels"
)

func NewCmd() *cobra.Command {
	var (
		config         string
		output         string
		bundle         string
		selector       string
		kubeconfig     string
		namespace      string
		serviceAccount string
		list           bool
		skipCleanup    bool
		waitTime       time.Duration
	)
	scorecardCmd := &cobra.Command{
		Use:    "scorecard",
		Short:  "Runs scorecard",
		Long:   `Has flags to configure dsl, bundle, and selector.`,
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {

			var err error
			o := scorecard.Options{
				ServiceAccount: serviceAccount,
				Namespace:      namespace,
				BundlePath:     bundle,
				OutputFormat:   output,
				Cleanup:        skipCleanup,
				WaitTime:       waitTime,
			}
			o.Client, err = scorecard.GetKubeClient(kubeconfig)
			if err != nil {
				return fmt.Errorf("could not get Kube connection %s", err.Error())
			}
			o.Config, err = scorecard.LoadConfig(config)
			if err != nil {
				return fmt.Errorf("could not find config file %s", err.Error())
			}

			if bundle == "" {
				return fmt.Errorf("bundle flag required")
			}

			o.Selector, err = labels.Parse(selector)
			if err != nil {
				return fmt.Errorf("could not parse selector %s", err.Error())
			}

			if list {
				return scorecard.ListTests(o)
			}

			return scorecard.RunTests(o)
		},
	}

	scorecardCmd.Flags().StringVarP(&config, "config", "c", "",
		"path to a new to be defined DSL yaml formatted file that configures what tests get executed")
	scorecardCmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "kubeconfig path")

	scorecardCmd.Flags().StringVar(&bundle, "bundle", "", "path to the operator bundle contents on disk")
	scorecardCmd.Flags().StringVarP(&selector, "selector", "l", "", "label selector to determine which tests are run")
	scorecardCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "namespace to run the test images in")
	scorecardCmd.Flags().StringVarP(&output, "output", "o", "text", "Output format for results.  Valid values: text, json")
	scorecardCmd.Flags().StringVarP(&serviceAccount, "service-account", "s", "default", "service account to use for tests")
	scorecardCmd.Flags().BoolVarP(&list, "list", "L", false, "option to enable listing which tests are run")
	scorecardCmd.Flags().BoolVarP(&skipCleanup, "skip-cleanup", "x", true, "option to disable resource cleanup after tests are run")
	scorecardCmd.Flags().DurationVarP(&waitTime, "wait-time", "w", time.Duration(30*time.Second), "time in seconds to wait for tests to complete")

	return scorecardCmd
}
