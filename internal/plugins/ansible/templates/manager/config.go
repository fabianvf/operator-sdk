// TODO(asmacdo) licence

package manager

import (
	"fmt"
	"os"
	"path/filepath"

	"sigs.k8s.io/kubebuilder/pkg/model/file"
)

var _ file.Template = &Config{}

// Config scaffolds yaml config for the manager.
type Config struct {
	file.TemplateMixin

	// Image is controller manager image name
	Image string

	// OperatorName will be used to create the pods
	OperatorName string
}

// SetTemplateDefaults implements input.Template
func (f *Config) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("config", "manager", "manager.yaml")
	}

	f.TemplateBody = configTemplate

	if f.OperatorName == "" {
		dir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("error to get the current path: %v", err)
		}
		f.OperatorName = filepath.Base(dir)
	}
	return nil
}

// todo(asmacdo): add the arg --enable-leader-election for the manager
// More info: https://github.com/operator-framework/operator-sdk/issues/3356

// todo(asmacdo): add the arg --metrics-addr for the manager
// More info: https://github.com/operator-framework/operator-sdk/issues/3358

const configTemplate = `apiVersion: v1
kind: Namespace
metadata:
  labels:
    control-plane: controller-manager
  name: system
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: controller-manager
  namespace: system
  labels:
    control-plane: controller-manager
spec:
  selector:
    matchLabels:
      control-plane: controller-manager
  replicas: 1
  template:
    metadata:
      labels:
        control-plane: controller-manager
    spec:
      containers:
      - args:
        - manager
        image: {{ .Image }}
        name: manager
				# resources:
        #   limits:
        #     cpu: 100m
        #     memory: 90Mi
        #   requests:
        #     cpu: 100m
        #     memory: 60Mi
        env:
          - name: WATCH_NAMESPACE
            value: ""
          - name: POD_NAME
            valueFrom:
              fieldRef:
                fieldPath: metadata.name
          - name: OPERATOR_NAME
            value: {{ .OperatorName }}
      terminationGracePeriodSeconds: 10
`

// -----------------------------------HELM
// package manager
//
// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
//
// 	"sigs.k8s.io/kubebuilder/pkg/model/file"
// )
//
// var _ file.Template = &Config{}
//
// // Config scaffolds yaml config for the manager.
// type Config struct {
// 	file.TemplateMixin
//
// 	// Image is controller manager image name
// 	Image string
//
// 	// OperatorName will be used to create the pods
// 	OperatorName string
// }
//
// // SetTemplateDefaults implements input.Template
// func (f *Config) SetTemplateDefaults() error {
// 	if f.Path == "" {
// 		f.Path = filepath.Join("config", "manager", "manager.yaml")
// 	}
//
// 	f.TemplateBody = configTemplate
//
// 	if f.OperatorName == "" {
// 		dir, err := os.Getwd()
// 		if err != nil {
// 			return fmt.Errorf("error to get the current path: %v", err)
// 		}
// 		f.OperatorName = filepath.Base(dir)
// 	}
// 	return nil
// }
//
// // todo(camilamacedo86): add the arg --enable-leader-election for the manager
// // More info: https://github.com/operator-framework/operator-sdk/issues/3356
//
// // todo(camilamacedo86): add the arg --metrics-addr for the manager
// // More info: https://github.com/operator-framework/operator-sdk/issues/3358
//
// const configTemplate = `apiVersion: v1
// kind: Namespace
// metadata:
//   labels:
//     control-plane: controller-manager
//   name: system
// ---
// apiVersion: apps/v1
// kind: Deployment
// metadata:
//   name: controller-manager
//   namespace: system
//   labels:
//     control-plane: controller-manager
// spec:
//   selector:
//     matchLabels:
//       control-plane: controller-manager
//   replicas: 1
//   template:
//     metadata:
//       labels:
//         control-plane: controller-manager
//     spec:
//       containers:
//       - args:
//         - manager
//         image: {{ .Image }}
//         name: manager
//         resources:
//           limits:
//             cpu: 100m
//             memory: 90Mi
//           requests:
//             cpu: 100m
//             memory: 60Mi
//         env:
//           - name: WATCH_NAMESPACE
//             value: ""
//           - name: POD_NAME
//             valueFrom:
//               fieldRef:
//                 fieldPath: metadata.name
//           - name: OPERATOR_NAME
//             value: {{ .OperatorName }}
//       terminationGracePeriodSeconds: 10
// `
