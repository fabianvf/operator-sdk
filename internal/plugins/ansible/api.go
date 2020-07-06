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
	// TODO(asmacdo) clean up
	// "fmt"
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
	groupFlag            = "group"
	versionFlag          = "version"
	kindFlag             = "kind"
	helmChartFlag        = "helm-chart"
	helmChartRepoFlag    = "helm-chart-repo"
	helmChartVersionFlag = "helm-chart-version"
	crdVersionFlag       = "crd-version"

	crdVersionV1      = "v1"
	crdVersionV1beta1 = "v1beta1"
)

type createAPIPlugin struct {
	config *config.Config

	group   string
	version string
	kind    string

	// CRDVersion is the version of the `apiextensions.k8s.io` API which will be used to generate the CRD.
	CRDVersion string

	// For help text.
	commandName string
}

var (
	_ plugin.CreateAPI   = &createAPIPlugin{}
	_ cmdutil.RunOptions = &createAPIPlugin{}
)

// TODO(asmacdo) document this
func (p *createAPIPlugin) UpdateContext(ctx *plugin.Context) {
	p.commandName = ctx.CommandName
}

func (p *createAPIPlugin) BindFlags(fs *pflag.FlagSet) {
	fs.SortFlags = false
	// TODO(asmacdo) short flags?
	fs.StringVar(&p.group, "group", "", "TODO(asmacdo) help text")
	fs.StringVar(&p.version, "version", "", "TODO(asmacdo) help text")
	fs.StringVar(&p.kind, "kind", "", "TODO(asmacdo) help text")
	fs.StringVar(&p.CRDVersion, crdVersionFlag, crdVersionV1, "crd versionFlag to generate")
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
	return scaffolds.NewCreateAPIScaffolder(p.config, p.group, p.version, p.kind, p.CRDVersion), nil
}

func (p *createAPIPlugin) PostScaffold() error {
	return nil
}
