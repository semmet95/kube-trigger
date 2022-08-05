/*
Copyright 2022 The KubeVela Authors.

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

package registry

import (
	"github.com/kubevela/kube-trigger/pkg/action/builtin/bumpapplicationrevision"
	"github.com/kubevela/kube-trigger/pkg/action/builtin/patchk8sobjects"
	"github.com/kubevela/kube-trigger/pkg/action/types"
)

// RegisterBuiltinActions registers builtin actions into registry.
// If you developed another action, put it here to make it available.
func RegisterBuiltinActions(reg *Registry) {
	registerFromInstance(reg, &patchk8sobjects.PatchK8sObjects{})
	registerFromInstance(reg, &bumpapplicationrevision.BumpApplicationRevision{})
}

func registerFromInstance(reg *Registry, act types.Action) {
	ins := act
	insMeta := types.ActionMeta{
		Type: ins.Type(),
	}
	reg.RegisterType(insMeta, ins)
}
