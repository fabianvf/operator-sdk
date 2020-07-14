/*
Copyright 2019 The Kubernetes Authors.

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

package scaffolds

import (
	"fmt"
	// "os"

	"sigs.k8s.io/kubebuilder/pkg/model"
	"sigs.k8s.io/kubebuilder/pkg/model/config"
	"sigs.k8s.io/kubebuilder/pkg/plugin/scaffold"

	// TODO(amacdo) convenience FROM new/cmd.go
	// "github.com/operator-framework/operator-sdk/internal/flags/apiflags"
	// "github.com/operator-framework/operator-sdk/internal/genutil"
	// "github.com/operator-framework/operator-sdk/internal/scaffold"
	// "github.com/operator-framework/operator-sdk/internal/scaffold/ansible"
	"github.com/operator-framework/operator-sdk/internal/plugins/ansible/templates"
	// "github.com/operator-framework/operator-sdk/internal/plugins/ansible/templates/manager"
	// "github.com/operator-framework/operator-sdk/internal/scaffold/helm"
	// "github.com/operator-framework/operator-sdk/internal/scaffold/input"
	// "github.com/operator-framework/operator-sdk/internal/util/projutil"

	"github.com/operator-framework/operator-sdk/internal/kubebuilder/machinery"
	// "github.com/operator-framework/operator-sdk/joe/pkg/plugin/internal/chartutil"
	// 	"github.com/operator-framework/operator-sdk/joe/pkg/plugin/v1/scaffolds/internal/templates"
	// 	"github.com/operator-framework/operator-sdk/joe/pkg/plugin/v1/scaffolds/internal/templates/manager"
	// 	"github.com/operator-framework/operator-sdk/joe/pkg/plugin/v1/scaffolds/internal/templates/metricsauth"
	// 	"github.com/operator-framework/operator-sdk/joe/pkg/plugin/v1/scaffolds/internal/templates/prometheus"
)

const (
	// KustomizeVersion is the kubernetes-sigs/kustomize version to be used in the project
	KustomizeVersion = "v3.5.4"

	imageName = "controller:latest"
)

var _ scaffold.Scaffolder = &initScaffolder{}

type initScaffolder struct {
	config           *config.Config
	apiScaffolder    scaffold.Scaffolder
	generatePlaybook bool
	group            string
	version          string
	kind             string
}

// NewInitScaffolder returns a new Scaffolder for project initialization operations
func NewInitScaffolder(config *config.Config, apiScaffolder scaffold.Scaffolder, generatePlaybook bool, group string, version string, kind string) scaffold.Scaffolder {
	// TODO(asmacdo) not pleased that generatePlaybook ended up in here.
	return &initScaffolder{
		config:           config,
		apiScaffolder:    apiScaffolder,
		generatePlaybook: generatePlaybook,
		group:            group,
		version:          version,
		kind:             kind,
	}
}

func (s *initScaffolder) newUniverse() *model.Universe {
	return model.NewUniverse(
		model.WithConfig(s.config),
	)
}

// Scaffold implements Scaffolder
func (s *initScaffolder) Scaffold() error {
	switch {
	case s.config.IsV3():
		if err := s.scaffold(); err != nil {
			return err
		}
		// if s.apiScaffolder != nil {
		// 	return s.apiScaffolder.Scaffold()
		// }
		return nil
	default:
		return fmt.Errorf("unknown project version %v", s.config.Version)
	}
}

func (s *initScaffolder) scaffold() error {
	fmt.Println("HELLLO internal/plugins/ansible/scaffolds/init.go")
	// TODO(asmacdo) create roles dir
	// if err := os.MkdirAll(chartutil.HelmChartsDir, 0755); err != nil {
	// 	return err
	// }
	return machinery.NewScaffold().Execute(
		s.newUniverse(),
		// &manager.Config{Image: imageName},
		&templates.Dockerfile{GeneratePlaybook: s.generatePlaybook},

		// &ansible.RolesReadme{Resource: *resource},
		// &ansible.RolesMetaMain{Resource: *resource},
		// &roleFiles,
		// &roleTemplates,
		// &ansible.RolesVarsMain{Resource: *resource},
		// &ansible.MoleculeTestLocalConverge{Resource: *resource},
		// &ansible.RolesDefaultsMain{Resource: *resource},
		// &ansible.RolesTasksMain{Resource: *resource},
		// &ansible.MoleculeDefaultMolecule{},
		// &ansible.MoleculeDefaultPrepare{},
		// &ansible.MoleculeDefaultConverge{
		// 	GeneratePlaybook: generatePlaybook,
		// 	Resource:         *resource,
		// },
		// &ansible.MoleculeDefaultVerify{},
		// &ansible.RolesHandlersMain{Resource: *resource},
		// &ansible.Watches{
		// 	GeneratePlaybook: generatePlaybook,
		// 	Resource:         *resource,
		// },
		// &ansible.DeployOperator{},
		// &ansible.Travis{},
		&templates.RequirementsYml{},
		// &ansible.MoleculeTestLocalMolecule{},
		// &ansible.MoleculeTestLocalPrepare{},
		// &ansible.MoleculeTestLocalVerify{},
		// &ansible.MoleculeClusterMolecule{Resource: *resource},
		// &ansible.MoleculeClusterCreate{},
		// &ansible.MoleculeClusterPrepare{Resource: *resource},
		// &ansible.MoleculeClusterConverge{},
		// &ansible.MoleculeClusterVerify{Resource: *resource},
		// &ansible.MoleculeClusterDestroy{Resource: *resource},
		// &ansible.MoleculeTemplatesOperator{},

		// &templates.GitIgnore{},
		// &templates.AuthProxyRole{},
		// &templates.AuthProxyRoleBinding{},
		// &metricsauth.AuthProxyPatch{},
		// &metricsauth.AuthProxyService{},
		// &metricsauth.ClientClusterRole{},
		// &manager.Config{Image: imageName},
		// &templates.Makefile{
		// 	Image:            imageName,
		// 	KustomizeVersion: KustomizeVersion,
		// },
		// &templates.Dockerfile{},
		// &templates.Kustomize{},
		// &templates.ManagerRoleBinding{},
		// &templates.LeaderElectionRole{},
		// &templates.LeaderElectionRoleBinding{},
		// &templates.KustomizeRBAC{},
		// &templates.Watches{},
		// &manager.Kustomization{},
		// &prometheus.Kustomization{},
		// &prometheus.ServiceMonitor{},
	)
}

// func doAnsibleScaffold() error {
// 	cfg := &input.Config{
// 		AbsProjectPath: filepath.Join(projutil.MustGetwd(), projectName),
// 		ProjectName:    projectName,
// 	}
//
// 	resource, err := scaffold.NewResource(apiFlags.APIVersion, apiFlags.Kind)
// 	if err != nil {
// 		return fmt.Errorf("invalid apiVersion and kind: %v", err)
// 	}
//
// 	roleFiles := ansible.RolesFiles{Resource: *resource}
// 	roleTemplates := ansible.RolesTemplates{Resource: *resource}
//
// 	s := &scaffold.Scaffold{}
// 	err = s.Execute(cfg,
// 		&scaffold.ServiceAccount{},
// 		&scaffold.Role{},
// 		&scaffold.RoleBinding{},
// 		&scaffold.CR{Resource: resource},
// 		&ansible.BuildDockerfile{GeneratePlaybook: generatePlaybook},
// 		&ansible.RolesReadme{Resource: *resource},
// 		&ansible.RolesMetaMain{Resource: *resource},
// 		&roleFiles,
// 		&roleTemplates,
// 		&ansible.RolesVarsMain{Resource: *resource},
// 		&ansible.MoleculeTestLocalConverge{Resource: *resource},
// 		&ansible.RolesDefaultsMain{Resource: *resource},
// 		&ansible.RolesTasksMain{Resource: *resource},
// 		&ansible.MoleculeDefaultMolecule{},
// 		&ansible.MoleculeDefaultPrepare{},
// 		&ansible.MoleculeDefaultConverge{
// 			GeneratePlaybook: generatePlaybook,
// 			Resource:         *resource,
// 		},
// 		&ansible.MoleculeDefaultVerify{},
// 		&ansible.RolesHandlersMain{Resource: *resource},
// 		&ansible.Watches{
// 			GeneratePlaybook: generatePlaybook,
// 			Resource:         *resource,
// 		},
// 		&ansible.DeployOperator{},
// 		&ansible.Travis{},
// 		&ansible.RequirementsYml{},
// 		&ansible.MoleculeTestLocalMolecule{},
// 		&ansible.MoleculeTestLocalPrepare{},
// 		&ansible.MoleculeTestLocalVerify{},
// 		&ansible.MoleculeClusterMolecule{Resource: *resource},
// 		&ansible.MoleculeClusterCreate{},
// 		&ansible.MoleculeClusterPrepare{Resource: *resource},
// 		&ansible.MoleculeClusterConverge{},
// 		&ansible.MoleculeClusterVerify{Resource: *resource},
// 		&ansible.MoleculeClusterDestroy{Resource: *resource},
// 		&ansible.MoleculeTemplatesOperator{},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("new ansible scaffold failed: %v", err)
// 	}
//
// 	if err = genutil.GenerateCRDNonGo(projectName, *resource, apiFlags.CrdVersion); err != nil {
// 		return err
// 	}
//
// 	// Remove placeholders from empty directories
// 	err = os.Remove(filepath.Join(s.AbsProjectPath, roleFiles.Path))
// 	if err != nil {
// 		return fmt.Errorf("new ansible scaffold failed: %v", err)
// 	}
// 	err = os.Remove(filepath.Join(s.AbsProjectPath, roleTemplates.Path))
// 	if err != nil {
// 		return fmt.Errorf("new ansible scaffold failed: %v", err)
// 	}
//
// 	// Decide on playbook.
// 	if generatePlaybook {
// 		log.Infof("Generating %s playbook.", strings.Title(operatorType))
//
// 		err := s.Execute(cfg,
// 			&ansible.Playbook{Resource: *resource},
// 		)
// 		if err != nil {
// 			return fmt.Errorf("new ansible playbook scaffold failed: %v", err)
// 		}
// 	}
//
// 	// update deploy/role.yaml for the given resource r.
// 	if err := scaffold.UpdateRoleForResource(resource, cfg.AbsProjectPath); err != nil {
// 		return fmt.Errorf("failed to update the RBAC manifest for the resource (%v, %v): %v",
// 			resource.APIVersion, resource.Kind, err)
// 	}
// 	return nil
// }
