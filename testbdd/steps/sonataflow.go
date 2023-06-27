package steps

import (
	"fmt"
	"github.com/cucumber/godog"
	kogitoFramework "github.com/kiegroup/kogito-operator/test/pkg/framework"
	"github.com/kiegroup/kogito-serverless-operator/api"
	"github.com/kiegroup/kogito-serverless-operator/api/v1alpha08"
	"github.com/kiegroup/kogito-serverless-operator/test"
	"github.com/kiegroup/kogito-serverless-operator/test/utils"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"path/filepath"
	"strings"
	"time"
)

func registerSonataFlowSteps(ctx *godog.ScenarioContext, data *Data) {
	ctx.Step(`^SonataFlow orderprocessing example is deployed$`, data.sonataFlowOrderProcessingExampleIsDeployed)
	ctx.Step(`^SonataFlow "([^"]*)" has the condition "(Running|Succeed|Built)" set to "(True|False|Unknown)" within (\d+) minutes$'`, data.sonataFlowHasTheConditionSetToWithinMinutes)
}

func (data *Data) sonataFlowOrderProcessingExampleIsDeployed() error {
	projectDir, _ := utils.GetProjectDir()
	projectDir = strings.Replace(projectDir, "/testbdd", "", -1)

	// TODO or kubectl
	out, err := kogitoFramework.CreateCommand("oc", "apply", "-f", filepath.Join(projectDir,
		test.GetServerlessWorkflowE2eOrderProcessingFolder()), "-n", data.Namespace).Execute()

	if err != nil {
		kogitoFramework.GetLogger(data.Namespace).Error(err, fmt.Sprintf("Applying SonataFlow failed, output: %s", out))
	}

	_ = <-time.After(120 * time.Second)
	return err
}

func (data *Data) sonataFlowHasTheConditionSetToWithinMinutes(name, status string, timeoutInMin int) error {
	return kogitoFramework.WaitForOnOpenshift(data.Namespace, fmt.Sprintf("SonataFlow %s has the status %s", name, status), timeoutInMin,
		func() (bool, error) {
			if sonataFlow, err := getSonataFlow(data.Namespace, name); err != nil {
				return false, err
			} else if sonataFlow != nil {
				return false, nil
			} else {
				condition := sonataFlow.Status.GetCondition(api.RunningConditionType)
				return condition != nil && condition.Status == v1.ConditionTrue, nil
			}
		})
}

func getSonataFlow(namespace, name string) (*v1alpha08.KogitoServerlessWorkflow, error) {
	sonataFlow := &v1alpha08.KogitoServerlessWorkflow{}
	if exists, err := kogitoFramework.GetObjectWithKey(types.NamespacedName{Namespace: namespace, Name: name}, sonataFlow); err != nil {
		return nil, fmt.Errorf("Error while trying to look for SonataFlow %s: %v ", name, err)
	} else if !exists {
		return nil, nil
	}
	return sonataFlow, nil
}
