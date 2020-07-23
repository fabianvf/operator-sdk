/*
Copyright 2020 The Kubernetes Authors.
Modifications copyright 2020 The Operator-SDK Authors

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
	// TODO(asmacdo) clean up
	"fmt"
	// "os"
	// "path/filepath"
	// "strings"

	"github.com/spf13/pflag"
	// "k8s.io/apimachinery/pkg/util/validation"
	"sigs.k8s.io/kubebuilder/pkg/model/config"
	"sigs.k8s.io/kubebuilder/pkg/plugin"
	"sigs.k8s.io/kubebuilder/pkg/plugin/scaffold"

	"github.com/operator-framework/operator-sdk/internal/kubebuilder/cmdutil"
	// "github.com/operator-framework/operator-sdk/internal/kubebuilder/validation"
	// "github.com/operator-framework/operator-sdk/joe/pkg/plugin/v1/scaffolds"
	"github.com/operator-framework/operator-sdk/internal/plugins/ansible/scaffolds"
)

const (
	groupFlag      = "group"
	versionFlag    = "version"
	kindFlag       = "kind"
	crdVersionFlag = "crd-version"

	crdVersionV1      = "v1"
	crdVersionV1beta1 = "v1beta1"
)

type createAPIPlugin struct {
	config        *config.Config
	createOptions scaffolds.CreateOptions
}

var (
	_ plugin.CreateAPI   = &createAPIPlugin{}
	_ cmdutil.RunOptions = &createAPIPlugin{}
)

// TODO(asmacdo) document this
func (p *createAPIPlugin) UpdateContext(ctx *plugin.Context) {
	ctx.Description = `Scaffold a Kubernetes API in which the controller is an Ansible role or playbook.
`
	ctx.Examples = fmt.Sprintf(`  $ %s create api \
      --group=apps --version=v1alpha1 \
      --kind=AppService
`,
		ctx.CommandName,
	)
}

func (p *createAPIPlugin) BindFlags(fs *pflag.FlagSet) {
	fs.SortFlags = false

	fs.StringVar(&p.createOptions.GVK.Group, groupFlag, "", "resource group")
	fs.StringVar(&p.createOptions.GVK.Version, versionFlag, "", "resource version")
	fs.StringVar(&p.createOptions.GVK.Kind, kindFlag, "", "resource kind")
	fs.StringVar(&p.createOptions.CRDVersion, crdVersionFlag, crdVersionV1, "crd version to generate")
}

func (p *createAPIPlugin) InjectConfig(c *config.Config) {
	p.config = c
}

func (p *createAPIPlugin) Run() error {
	return cmdutil.Run(p)
}

// TODO(asmacdo) validate
func (p *createAPIPlugin) Validate() error {
	// TODO(asmacdo) must not be empty
	return nil
}

func (p *createAPIPlugin) GetScaffolder() (scaffold.Scaffolder, error) {
	return scaffolds.NewCreateAPIScaffolder(p.config, p.createOptions), nil
}

func (p *createAPIPlugin) PostScaffold() error {
	return nil
}
