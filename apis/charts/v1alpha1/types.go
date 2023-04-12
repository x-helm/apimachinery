package v1alpha1

type SourceKind string

const (
	SourceGroupHelmRepository = "source.toolkit.fluxcd.io"
	SourceKindHelmRepository  = "HelmRepository"

	SourceGroupLegacy = GroupName
	SourceKindLegacy  = "Legacy"

	SourceGroupLocal = GroupName
	SourceKindLocal  = "Local"

	SourceGroupEmbed = GroupName
	SourceKindEmbed  = "Embed"
)
