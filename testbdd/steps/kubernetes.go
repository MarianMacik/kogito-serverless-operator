package steps

import (
	"github.com/cucumber/godog"
	"github.com/kiegroup/kogito-operator/test/pkg/framework"
)

func registerKubernetesSteps(ctx *godog.ScenarioContext, data *Data) {
	// Added step to the one from the Kogito framework to be able to use quotes inside the text
	ctx.Step(`^Deployment "([^"]*)" pods log contains text '([^']*)' within (\d+) minutes$`, data.deploymentPodsLogContainsTextWithinMinutes)
}

func (data *Data) deploymentPodsLogContainsTextWithinMinutes(dName, logText string, timeoutInMin int) error {
	return framework.WaitForAllPodsByDeploymentToContainTextInLog(data.Namespace, dName, logText, timeoutInMin)
}
