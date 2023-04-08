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
	"x-helm.dev/apimachinery/apis"
	"x-helm.dev/apimachinery/crds"

	"k8s.io/apimachinery/pkg/fields"
	"kmodules.xyz/client-go/apiextensions"
)

func (_ Plan) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(GroupVersion.WithResource(ResourcePlans))
}

func (plan *Plan) SetLabels(planName, planID, prodID, phase string) {
	labelMap := map[string]string{
		apis.LabelPlanName:  planName,
		apis.LabelPlanID:    planID,
		apis.LabelProductID: prodID,
		apis.LabelPlanPhase: phase,
	}
	plan.ObjectMeta.SetLabels(labelMap)
}

func (_ Plan) FormatLabels(planName, planID, prodID, phase string) string {
	labelMap := make(map[string]string)
	if planName != "" {
		labelMap[apis.LabelPlanName] = planName
	}
	if planID != "" {
		labelMap[apis.LabelPlanID] = planID
	}
	if prodID != "" {
		labelMap[apis.LabelProductID] = prodID
	}
	if phase != "" {
		labelMap[apis.LabelPlanPhase] = phase
	}
	return fields.SelectorFromSet(labelMap).String()
}
