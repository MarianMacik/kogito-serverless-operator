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
	"net/url"
	"path/filepath"
	"strings"
)

func registerSonataFlowSteps(ctx *godog.ScenarioContext, data *Data) {
	ctx.Step(`^SonataFlow orderprocessing example is deployed$`, data.sonataFlowOrderProcessingExampleIsDeployed)
	ctx.Step(`^SonataFlow "([^"]*)" has the condition "(Running|Succeed|Built)" set to "(True|False|Unknown)" within (\d+) minutes?$`, data.sonataFlowHasTheConditionSetToWithinMinutes)
	ctx.Step(`^SonataFlow "([^"]*)" is addressable within (\d+) minutes?$`, data.sonataFlowIsAddressableWithinMinutes)
}

func (data *Data) sonataFlowOrderProcessingExampleIsDeployed() error {
	projectDir, _ := utils.GetProjectDir()
	projectDir = strings.Replace(projectDir, "/testbdd", "", -1)

	// TODO or kubectl
	out, err := kogitoFramework.CreateCommand("oc", "apply", "-f", filepath.Join(projectDir,
		test.GetSonataFlowE2eOrderProcessingFolder()), "-n", data.Namespace).Execute()

	if err != nil {
		kogitoFramework.GetLogger(data.Namespace).Error(err, fmt.Sprintf("Applying SonataFlow failed, output: %s", out))
	}
	return err
}

func (data *Data) sonataFlowHasTheConditionSetToWithinMinutes(name, conditionType, conditionStatus string, timeoutInMin int) error {
	return kogitoFramework.WaitForOnOpenshift(data.Namespace, fmt.Sprintf("SonataFlow %s has the condition %s with status %s", name, conditionType, conditionStatus), timeoutInMin,
		func() (bool, error) {
			if sonataFlow, err := getSonataFlow(data.Namespace, name); err != nil {
				return false, err
			} else if sonataFlow == nil {
				return false, nil
			} else {
				condition := sonataFlow.Status.GetCondition(api.ConditionType(conditionType))
				return condition != nil && condition.Status == v1.ConditionStatus(conditionStatus), nil
			}
		})
}

//func (data *Data) sonataFlowIsAddressable2(name string) error {
//	return kogitoFramework.WaitForOnOpenshift(data.Namespace, fmt.Sprintf("SonataFlow %s is addressable", name), 1,
//		func() (bool, error) {
//			sonataFlow, err := getSonataFlow(data.Namespace, name)
//			if err != nil {
//				return false, err
//			} else if sonataFlow == nil {
//				return false, fmt.Errorf("No SonataFlow found with name %s in namespace %s", name, data.Namespace)
//			}
//
//			if sonataFlow.Status.Address.URL == nil {
//				return false, fmt.Errorf("SonataFlow %s does NOT have an address", name)
//			}
//
//			if _, err := url.ParseRequestURI(sonataFlow.Status.Address.URL.String()); err != nil {
//				return false, fmt.Errorf("SonataFlow %s address '%s' is not valid: %w", name, sonataFlow.Status.Address.URL, err)
//			}
//
//			return true, nil
//		})
//
//}

func (data *Data) sonataFlowIsAddressableWithinMinutes(name string, timeoutInMin int) error {
	return kogitoFramework.WaitForOnOpenshift(data.Namespace, fmt.Sprintf("SonataFlow %s is addressable", name), timeoutInMin,
		func() (bool, error) {
			sonataFlow, err := getSonataFlow(data.Namespace, name)
			if err != nil {
				return false, err
			} else if sonataFlow == nil {
				return false, fmt.Errorf("No SonataFlow found with name %s in namespace %s", name, data.Namespace)
			}

			if sonataFlow.Status.Address.URL == nil {
				return false, fmt.Errorf("SonataFlow %s does NOT have an address", name)
			}

			if _, err := url.ParseRequestURI(sonataFlow.Status.Address.URL.String()); err != nil {
				return false, fmt.Errorf("SonataFlow %s address '%s' is not valid: %w", name, sonataFlow.Status.Address.URL, err)
			}

			return true, nil
		})
}

func getSonataFlow(namespace, name string) (*v1alpha08.SonataFlow, error) {
	sonataFlow := &v1alpha08.SonataFlow{}
	if exists, err := kogitoFramework.GetObjectWithKey(types.NamespacedName{Namespace: namespace, Name: name}, sonataFlow); err != nil {
		return nil, fmt.Errorf("Error while trying to look for SonataFlow %s: %w ", name, err)
	} else if !exists {
		return nil, nil
	}
	return sonataFlow, nil
}
