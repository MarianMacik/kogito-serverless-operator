package steps

import (
	"github.com/cucumber/godog"
	"github.com/kiegroup/kogito-serverless-operator/testbdd/installers"

	"github.com/kiegroup/kogito-operator/test/pkg/config"
	kogitoFramework "github.com/kiegroup/kogito-operator/test/pkg/framework"
	kogitoInstallers "github.com/kiegroup/kogito-operator/test/pkg/installers"
)

func registerOperatorSteps(ctx *godog.ScenarioContext, data *Data) {
	ctx.Step(`^SonataFlow Operator is deployed$`, data.sonataFlowOperatorIsDeployed)
	//ctx.Step(`^SonataFlow Operator has (\d+) (?:pod|pods) running"$`, data.sonataFlowOperatorHasPodsRunning)
	// Not migrated yet
	ctx.Step(`^Kogito operator should be installed$`, data.kogitoOperatorShouldBeInstalled)
	ctx.Step(`^CLI install Kogito operator$`, data.cliInstallKogitoOperator)
}

func (data *Data) kogitoOperatorShouldBeInstalled() error {
	return kogitoFramework.WaitForKogitoOperatorRunning(data.Namespace)
}

func (data *Data) sonataFlowOperatorIsDeployed() (err error) {
	var installer kogitoInstallers.ServiceInstaller
	if config.UseProductOperator() {
		installer, err = kogitoInstallers.GetRhpamKogitoInstaller()
	} else {
		installer, err = installers.GetSonataFlowInstaller()
	}
	if err != nil {
		return err
	}
	return installer.Install(data.Namespace)
}

func (data *Data) cliInstallKogitoOperator() error {
	_, err := kogitoFramework.ExecuteCliCommandInNamespace(data.Namespace, "install", "operator")
	return err
}

//func (data *Data) sonataFlowOperatorHasPodsRunning(numberOfPods int, name, phase string) error {
//	return kogitoFramework.WaitForPodsWithLabel(data.Namespace, "control-plane", "controller-manager", numberOfPods, 1)
//}
