package steps

import (
	"fmt"
	"github.com/cucumber/godog"
	kogitoFramework "github.com/kiegroup/kogito-operator/test/pkg/framework"
	"github.com/kiegroup/kogito-serverless-operator/test"
	"github.com/kiegroup/kogito-serverless-operator/test/utils"
	"path/filepath"
	"strings"
	//"github.com/kiegroup/kogito-serverless-operator/test"
	"os"
)

const (
	minikubePlatform  = "minikube"
	openshiftPlatform = "openshift"
)

func registerPlatformSteps(ctx *godog.ScenarioContext, data *Data) {
	ctx.Step(`^SonataFlowPlatform is deployed$`, data.sonataFlowPlatformIsDeployed)
}

func (data *Data) sonataFlowPlatformIsDeployed() error {
	projectDir, _ := utils.GetProjectDir()
	projectDir = strings.Replace(projectDir, "/testbdd", "", -1)

	// TODO or kubectl
	out, err := kogitoFramework.CreateCommand("oc", "apply", "-f", filepath.Join(projectDir, getSonataFlowPlatformFilename()), "-n", data.Namespace).Execute()

	if err != nil {
		kogitoFramework.GetLogger(data.Namespace).Error(err, fmt.Sprintf("Applying SonataFlowPlatform failed, output: %s", out))
	}

	return err
}

func getSonataFlowPlatformFilename() string {
	if getClusterPlatform() == openshiftPlatform {
		return test.GetPlatformOpenshiftE2eTest()
	}
	return test.GetPlatformMinikubeE2eTest()
}

func getClusterPlatform() string {
	if v, ok := os.LookupEnv("CLUSTER_PLATFORM"); ok {
		return v
	}
	return minikubePlatform
}
