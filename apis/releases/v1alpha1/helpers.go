package v1alpha1

import (
	kmapi "kmodules.xyz/client-go/api/v1"
)

func (ref *ChartSourceFlatRef) SetDefaults() *ChartSourceFlatRef {
	if ref.SourceAPIGroup == "" {
		ref.SourceAPIGroup = "source.toolkit.fluxcd.io"
	}
	if ref.SourceKind == "" {
		ref.SourceKind = "HelmRepository"
	} else if ref.SourceKind == "Legacy" || ref.SourceKind == "Local" || ref.SourceKind == "Embed" {
		ref.SourceAPIGroup = "charts.x-helm.dev"
	}
	return ref
}

func (ref *ChartSourceFlatRef) FromAPIObject(obj ChartSourceRef) *ChartSourceFlatRef {
	ref.Name = obj.Name
	ref.Version = obj.Version
	ref.SourceAPIGroup = obj.SourceRef.APIGroup
	ref.SourceKind = obj.SourceRef.Kind
	ref.SourceNamespace = obj.SourceRef.Namespace
	ref.SourceName = obj.SourceRef.Name
	return ref.SetDefaults()
}

func (ref *ChartSourceFlatRef) ToAPIObject() ChartSourceRef {
	ref.SetDefaults()
	return ChartSourceRef{
		Name:    ref.Name,
		Version: ref.Version,
		SourceRef: kmapi.TypedObjectReference{
			APIGroup:  ref.SourceAPIGroup,
			Kind:      ref.SourceKind,
			Namespace: ref.SourceNamespace,
			Name:      ref.Name,
		},
	}
}
