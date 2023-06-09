/*
Copyright AppsCode Inc. and Contributors

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

package v1alpha1

import (
	"strconv"

	"k8s.io/apimachinery/pkg/labels"
	"kmodules.xyz/client-go/apiextensions"
	"x-helm.dev/apimachinery/apis"
	"x-helm.dev/apimachinery/crds"
)

func (_ Product) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(GroupVersion.WithResource(ResourceProducts))
}

func (prod *Product) ResetLabels(prodID, prodKey, phase string, ownerID int64) {
	labelMap := map[string]string{
		apis.LabelProductID:      prodID,
		apis.LabelProductKey:     prodKey,
		apis.LabelProductPhase:   phase,
		apis.LabelProductOwnerID: strconv.FormatInt(ownerID, 10),
	}
	prod.ObjectMeta.SetLabels(labelMap)
}

func (_ Product) FormatLabels(prodID, prodKey, phase string, ownerID int64) labels.Selector {
	labelMap := make(map[string]string)
	if prodID != "" {
		labelMap[apis.LabelProductID] = prodID
	}
	if prodKey != "" {
		labelMap[apis.LabelProductKey] = prodKey
	}
	if phase != "" {
		labelMap[apis.LabelProductPhase] = phase
	}
	if ownerID != 0 {
		labelMap[apis.LabelProductOwnerID] = strconv.FormatInt(ownerID, 10)
	}
	return labels.SelectorFromSet(labelMap)
}
