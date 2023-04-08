# apimachinery

## Project Generator Commands

```bash
> kubebuilder init --domain x-helm.dev --skip-go-version-check
> kubebuilder edit --multigroup=true

> kubebuilder create api --group charts --version v1alpha1 --kind ChartPreset
> kubebuilder create api --group charts --version v1alpha1 --kind ClusterChartPreset --namespaced=false

> kubebuilder create api --group drivers --version v1alpha1 --kind AppRelease

> kubebuilder create api --group releases --version v1alpha1 --kind Bundle --namespaced=false
> kubebuilder create api --group releases --version v1alpha1 --kind Order --namespaced=false

> kubebuilder create api --group products --version v1alpha1 --kind Product --namespaced=false
> kubebuilder create api --group products --version v1alpha1 --kind Plan --namespaced=false
```
