/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/apis/extensions/v1beta1"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/resource"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/strategicpatch"
)

// GetOptions is the start of the data required to perform the operation.  As new fields are added, add them here instead of
// referencing the cmd.Flags()

const (
	link_example = `# Link a secret to a deploymentInfo
kubectl link secret/mysecret deploymentInfo/myapp`
)

// NewCmdGet creates a command object for the generic "get" action, which
// retrieves one or more resources from a server.
func NewCmdLink(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "link [(-o|--output=)json|yaml|wide|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=...] (TYPE [NAME | -l label] | TYPE/NAME ...) [flags]",
		Short:   "Link a secret into a pod-containing resource",
		Long:    "Link a secret into a pod-containing resource",
		Example: link_example,
		Run: func(cmd *cobra.Command, args []string) {
			err := RunLink(f, out, cmd, args)
			cmdutil.CheckErr(err)
		},
	}
	return cmd
}

// RunLink implements the Link command
func RunLink(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, args []string) error {
	cmdNamespace, _, err := f.DefaultNamespace()
	if err != nil {
		return err
	}
	mapper, typer := f.Object(cmdutil.GetIncludeThirdPartyAPIs(cmd))
	r := resource.NewBuilder(mapper, typer, resource.ClientMapperFunc(f.ClientForMapping), f.Decoder(true)).
		NamespaceParam(cmdNamespace).DefaultNamespace().
		ResourceTypeOrNameArgs(true, args...).
		Flatten().
		Do()

	err = r.Err()
	if err != nil {
		return err
	}
	infos, err := r.Infos()
	if err != nil {
		return err
	}
	if len(infos) != 2 {
		return fmt.Errorf("TODO %d", len(infos))
	}
	// secret
	secretInfo := infos[0]
	// deploymentInfo
	deploymentInfo := infos[1]

	fmt.Printf("From:\n%#v\nTo:\n%#v\n", secretInfo, deploymentInfo)

	deploymentObj, err := deploymentInfo.Mapping.ConvertToVersion(deploymentInfo.Object, deploymentInfo.Mapping.GroupVersionKind.GroupVersion())
	if err != nil {
		return err
	}

	deployment := deploymentObj.(*v1beta1.Deployment)

	oldData, err := json.Marshal(deployment)
	if err != nil {
		return err
	}

	// modify here

	secret, ok := secretInfo.Object.(*api.Secret)
	if !ok {
		return fmt.Errorf("expected secret type")
	}

	fmt.Printf("%#v\n\n", secret)

	for i := range deployment.Spec.Template.Spec.Containers {
		container := &deployment.Spec.Template.Spec.Containers[i]
		for key := range secret.Data {
			envvar := v1.EnvVar{
				Name: fmt.Sprintf("%s_%s", strings.ToUpper(secret.Name), strings.ToUpper(key)),
				ValueFrom: &v1.EnvVarSource{
					SecretKeyRef: &v1.SecretKeySelector{
						LocalObjectReference: v1.LocalObjectReference{
							Name: secret.Name,
						},
						Key: key,
					},
				},
			}
			container.Env = append(container.Env, envvar)
		}
	}

	fmt.Printf("%#v\n", deployment.Spec.Template.Spec.Containers[0].Env[0].ValueFrom.SecretKeyRef)

	newData, err := json.Marshal(deployment)
	if err != nil {
		return err
	}
	patchBytes, err := strategicpatch.CreateTwoWayMergePatch(oldData, newData, deployment)

	fmt.Printf("%s\n", string(patchBytes))

	createdPatch := err == nil
	if err != nil {
		glog.V(2).Infof("couldn't compute patch: %v", err)
	}

	mapping := deploymentInfo.ResourceMapping()
	client, err := f.ClientForMapping(mapping)
	if err != nil {
		return err
	}
	helper := resource.NewHelper(client, mapping)

	name, namespace := deploymentInfo.Name, deploymentInfo.Namespace
	var outputObj runtime.Object
	if createdPatch {
		outputObj, err = helper.Patch(namespace, name, api.StrategicMergePatchType, patchBytes)
	} else {
		outputObj, err = helper.Replace(namespace, name, false, deployment)
	}
	if err != nil {
		return err
	}
	_ = outputObj

	/*mapper, _ := f.Object(cmdutil.GetIncludeThirdPartyAPIs(o.cmd))
	outputFormat := cmdutil.GetFlagString(o.cmd, "output")
	if outputFormat != "" {
		return f.PrintObject(o.cmd, mapper, outputObj, o.out)
	}

	cmdutil.PrintSuccess(mapper, false, o.out, info.Mapping.Resource, info.Name, "linked")*/

	fmt.Println("yay!")

	return nil
}
