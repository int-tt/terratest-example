package test

import (
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
)

func TestHelmBasicExampleTemplateRenderedDeplyment(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../")
	require.NoError(t, err)

	options := &helm.Options{
		SetValues: map[string]string{
			"containerImageRepo": "nginx",
			"containerImageTag":  "1.15.8",
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, []string{"templates/deployment.yaml"})

	var deployment appsv1.Deployment
	helm.UnmarshalK8SYaml(t, output, &deployment)
	expectedContainerImage := "nginx:1.15.8"
	deploymentContainers := deployment.Spec.Template.Spec.Containers
	require.Equal(t, len(deploymentContainers), 1)
	require.Equal(t, deploymentContainers[0].Image, expectedContainerImage)
}
