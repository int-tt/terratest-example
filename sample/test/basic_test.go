package test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/require"
	"go.mozilla.org/sops/decrypt"
	appsv1 "k8s.io/api/apps/v1"
)

// func TestHelmBasicExampleTemplateRenderedDeplyment(t *testing.T) {
// 	t.Parallel()

// 	helmChartPath, err := filepath.Abs("../")
// 	require.NoError(t, err)

// 	options := &helm.Options{
// 		SetValues: map[string]string{
// 			"containerImageRepo": "nginx",
// 			"containerImageTag":  "1.15.8",
// 		},
// 	}

// 	output := helm.RenderTemplate(t, options, helmChartPath, []string{"templates/deployment.yaml"})

// 	var deployment appsv1.Deployment
// 	helm.UnmarshalK8SYaml(t, output, &deployment)
// 	expectedContainerImage := "nginx:1.15.8"
// 	deploymentContainers := deployment.Spec.Template.Spec.Containers
// 	require.Equal(t, len(deploymentContainers), 1)
// 	require.Equal(t, deploymentContainers[0].Image, expectedContainerImage)
// }

func TestHelmImportValuesTemplateRnderedDev(t *testing.T) {
	t.Parallel()
	helmChartPath, err := filepath.Abs("../")
	require.NoError(t, err)
	options := &helm.Options{
		ValuesFiles: []string{"../env/dev.yaml"},
	}
	output := helm.RenderTemplate(t, options, helmChartPath, []string{"templates/deployment.yaml"})

	var deployment appsv1.Deployment
	helm.UnmarshalK8SYaml(t, output, &deployment)
	expectedContainerImage := "nginx:dev"
	deploymentContainers := deployment.Spec.Template.Spec.Containers
	require.Equal(t, len(deploymentContainers), 1)
	require.Equal(t, deploymentContainers[0].Image, expectedContainerImage)
}

func TestHelmImportValuesTemplateRnderedProd(t *testing.T) {
	t.Parallel()
	helmChartPath, err := filepath.Abs("../")
	require.NoError(t, err)

	content, _ := decrypt.File("../env/prod.enc.yaml", "yaml")
	fp, _ := ioutil.TempFile("", "prod.yaml")

	ioutil.WriteFile(fp.Name(), content, 0644)

	options := &helm.Options{
		ValuesFiles: []string{fp.Name()},
	}
	output := helm.RenderTemplate(t, options, helmChartPath, []string{"templates/deployment.yaml"})

	var deployment appsv1.Deployment
	helm.UnmarshalK8SYaml(t, output, &deployment)
	expectedContainerImage := "nginx:prod"
	deploymentContainers := deployment.Spec.Template.Spec.Containers
	require.Equal(t, len(deploymentContainers), 1)
	require.Equal(t, deploymentContainers[0].Image, expectedContainerImage)
}
