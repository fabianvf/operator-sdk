/*
TODO(asmacdo) sdk licence header
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ansible

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder/pkg/model/config"
	"sigs.k8s.io/kubebuilder/pkg/plugin"
	"sigs.k8s.io/kubebuilder/pkg/plugin/scaffold"

	"github.com/operator-framework/operator-sdk/joe/pkg/notinternal/kubebuilder/cmdutil"
	"github.com/operator-framework/operator-sdk/joe/pkg/notinternal/kubebuilder/validation"
	// "github.com/operator-framework/operator-sdk/joe/pkg/plugin/v1/scaffolds"
	"github.com/operator-framework/operator-sdk/internal/plugins/ansible/scaffolds"
)

type initPlugin struct {
	config *config.Config
	// apiPlugin     createAPIPlugin

	// TODO(asmacdo) should be in CreateAPI plugin
	generatePlaybook bool

	doAPIScaffold bool

	// For help text.
	commandName string
}

var (
	_ plugin.Init        = &initPlugin{}
	_ cmdutil.RunOptions = &initPlugin{}
)

// TODO(asmacdo) all this got to change
func (p *initPlugin) UpdateContext(ctx *plugin.Context) {
	ctx.Description = `Initialize a new Helm-based operator project.

Writes the following files:
- a helm-charts directory with the chart(s) to build releases from
- a watches.yaml file that defines the mapping between your API and a Helm chart
- a PROJECT file with the domain and repo
- a Makefile to build the project
- a Kustomization.yaml for customizating manifests
- a Patch file for customizing image for manager manifests
- a Patch file for enabling prometheus metrics
`
	ctx.Examples = fmt.Sprintf(`  $ %s init --plugins=%s \
      --domain=example.com \
      --group=apps --version=v1alpha1 \
      --kind=AppService

  $ %s init --plugins=%s \
      --domain=example.com \
      --group=apps --version=v1alpha1 \
      --kind=AppService \
      --helm-chart=myrepo/app

  $ %s init --plugins=%s \
      --domain=example.com \
      --helm-chart=myrepo/app

  $ %s init --plugins=%s \
      --domain=example.com \
      --helm-chart=myrepo/app \
      --helm-chart-version=1.2.3

  $ %s init --plugins=%s \
      --domain=example.com \
      --helm-chart=app \
      --helm-chart-repo=https://charts.mycompany.com/

  $ %s init --plugins=%s \
      --domain=example.com \
      --helm-chart=app \
      --helm-chart-repo=https://charts.mycompany.com/ \
      --helm-chart-version=1.2.3

  $ %s init --plugins=%s \
      --domain=example.com \
      --helm-chart=/path/to/local/chart-directories/app/

  $ %s init --plugins=%s \
      --domain=example.com \
      --helm-chart=/path/to/local/chart-archives/app-1.2.3.tgz
`,
		ctx.CommandName, plugin.KeyFor(Plugin{}),
		ctx.CommandName, plugin.KeyFor(Plugin{}),
		ctx.CommandName, plugin.KeyFor(Plugin{}),
		ctx.CommandName, plugin.KeyFor(Plugin{}),
		ctx.CommandName, plugin.KeyFor(Plugin{}),
		ctx.CommandName, plugin.KeyFor(Plugin{}),
		ctx.CommandName, plugin.KeyFor(Plugin{}),
		ctx.CommandName, plugin.KeyFor(Plugin{}),
	)

	p.commandName = ctx.CommandName
}

func (p *initPlugin) BindFlags(fs *pflag.FlagSet) {
	fs.SortFlags = false
	fs.StringVar(&p.config.Domain, "domain", "my.domain", "domain for groups")

	// p.apiPlugin.BindFlags(fs)
	// TODO(asmacdo) move this into api plugin?
	fs.BoolVarP(&p.generatePlaybook, "generate-playbook", "", false, "Generate a playbook skeleton. (Only used for TODO(asmacdo)--type ansible)")
}

func (p *initPlugin) InjectConfig(c *config.Config) {
	// v3 project configs get a 'layout' value.
	c.Layout = plugin.KeyFor(Plugin{})
	p.config = c
}

func (p *initPlugin) Run() error {
	return cmdutil.Run(p)
}

func (p *initPlugin) Validate() error {
	fmt.Println("initPlugin.Validate() internal/plugins/ansible/init.go")
	// Check if the project name is a valid namespace according to k8s
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error to get the current path: %v", err)
	}
	projectName := filepath.Base(dir)
	if err := validation.IsDNS1123Label(strings.ToLower(projectName)); err != nil {
		return fmt.Errorf("project name (%s) is invalid: %v", projectName, err)
	}

	// TODO(asmacdo)
	// defaultOpts := chartutil.CreateOptions{CRDVersion: "v1"}
	// if !p.apiPlugin.gvk.Empty() || p.apiPlugin.createOptions != defaultOpts {
	// 	p.doAPIScaffold = true
	// 	return p.apiPlugin.Validate()
	// }
	return nil
}

func (p *initPlugin) GetScaffolder() (scaffold.Scaffolder, error) {
	// p.apiPlugin.config = p.config
	var (
		apiScaffolder scaffold.Scaffolder
		// err           error
	)
	// if p.doAPIScaffold {
	// 	apiScaffolder, err = p.apiPlugin.GetScaffolder()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	return scaffolds.NewInitScaffolder(p.config, apiScaffolder, p.generatePlaybook), nil
}

func (p *initPlugin) PostScaffold() error {
	if !p.doAPIScaffold {
		fmt.Printf("Next: define a resource with:\n$ %s create api\n", p.commandName)
	}
	return nil
}
