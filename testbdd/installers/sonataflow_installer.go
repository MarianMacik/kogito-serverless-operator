// Copyright 2023 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package installers

import (
	"errors"
	"fmt"
	"github.com/kiegroup/kogito-operator/version/app"
	"regexp"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/kiegroup/kogito-operator/test/pkg/config"
	"github.com/kiegroup/kogito-operator/test/pkg/framework"
	"github.com/kiegroup/kogito-operator/test/pkg/installers"
	srvframework "github.com/kiegroup/kogito-serverless-operator/testbdd/framework"
)

var (
	// kogitoYamlClusterInstaller installs Kogito operator cluster wide using YAMLs
	sonataFlowYamlClusterInstaller = installers.YamlClusterWideServiceInstaller{
		InstallClusterYaml:               installSonataFlowUsingYaml,
		InstallationNamespace:            KogitoNamespace,
		WaitForClusterYamlServiceRunning: waitForKogitoOperatorUsingYamlRunning,
		GetAllClusterYamlCrsInNamespace:  getKogitoCrsInNamespace,
		UninstallClusterYaml:             uninstallKogitoUsingYaml,
		ClusterYamlServiceName:           kogitoServiceName,
		CleanupClusterYamlCrsInNamespace: cleanupKogitoCrsInNamespace,
	}

	// kogitoCustomOlmClusterWideInstaller installs Kogito cluster wide using OLM with custom catalog
	sonataFlowCustomOlmClusterWideInstaller = installers.OlmClusterWideServiceInstaller{
		SubscriptionName:                    kogitoOperatorSubscriptionName,
		Channel:                             kogitoOperatorSubscriptionChannel,
		Catalog:                             framework.GetCustomKogitoOperatorCatalog,
		InstallationTimeoutInMinutes:        5,
		GetAllClusterWideOlmCrsInNamespace:  getKogitoCrsInNamespace,
		CleanupClusterWideOlmCrsInNamespace: cleanupKogitoCrsInNamespace,
	}

	// kogitoOlmClusterWideInstaller installs Kogito cluster wide using OLM with community catalog
	sonataFlowOlmClusterWideInstaller = installers.OlmClusterWideServiceInstaller{
		SubscriptionName:                    kogitoOperatorSubscriptionName,
		Channel:                             kogitoOperatorSubscriptionChannel,
		Catalog:                             framework.GetCommunityCatalog,
		InstallationTimeoutInMinutes:        5,
		GetAllClusterWideOlmCrsInNamespace:  getKogitoCrsInNamespace,
		CleanupClusterWideOlmCrsInNamespace: cleanupKogitoCrsInNamespace,
	}

	// KogitoNamespace is the kogito namespace for yaml cluster-wide deployment
	KogitoNamespace   = "kogito-serverless-operator-system"
	kogitoServiceName = "Kogito serverless operator"

	kogitoOperatorSubscriptionName    = "kogito-serverless-operator"
	kogitoOperatorSubscriptionChannel = "alpha"
)

// GetSonataFlowInstaller returns Kogito installer
func GetSonataFlowInstaller() (installers.ServiceInstaller, error) {
	// If user doesn't pass Kogito operator image then use community OLM catalog to install operator
	if len(config.GetOperatorImageTag()) == 0 {
		framework.GetMainLogger().Info("Installing SonataFlow operator using community catalog.")
		return &sonataFlowOlmClusterWideInstaller, nil
	}

	if config.IsOperatorInstalledByYaml() || config.IsOperatorProfiling() {
		return &sonataFlowYamlClusterInstaller, nil
	}

	if config.IsOperatorInstalledByOlm() {
		return &sonataFlowCustomOlmClusterWideInstaller, nil
	}

	return nil, errors.New("no SonataFlow operator installer available for provided configuration")
}

func installSonataFlowUsingYaml() error {
	framework.GetMainLogger().Info("Installing Kogito operator")

	yamlContent, err := framework.ReadFromURI(config.GetOperatorYamlURI())
	if err != nil {
		framework.GetMainLogger().Error(err, "Error while reading the operator YAML file")
		return err
	}

	regexp, err := regexp.Compile("quay.io/kiegroup/kogito-serverless-operator-nightly.*" + ":" + app.Version) // TODO use version/version.go from this repository
	if err != nil {
		return err
	}
	yamlContent = regexp.ReplaceAllString(yamlContent, config.GetOperatorImageTag())

	tempFilePath, err := framework.CreateTemporaryFile("kogito-operator*.yaml", yamlContent)
	if err != nil {
		framework.GetMainLogger().Error(err, "Error while storing adjusted YAML content to temporary file")
		return err
	}

	_, err = framework.CreateCommand("oc", "apply", "-f", tempFilePath).Execute()
	if err != nil {
		framework.GetMainLogger().Error(err, "Error while installing Kogito operator from YAML file")
		return err
	}

	return nil
}

func waitForKogitoOperatorUsingYamlRunning() error {
	return srvframework.WaitForKogitoOperatorRunning(KogitoNamespace)
}

func uninstallKogitoUsingYaml() error {
	framework.GetMainLogger().Info("Uninstalling Kogito operator")

	output, err := framework.CreateCommand("oc", "delete", "-f", config.GetOperatorYamlURI(), "--timeout=30s", "--ignore-not-found=true").Execute()
	if err != nil {
		framework.GetMainLogger().Error(err, fmt.Sprintf("Deleting Kogito operator failed, output: %s", output))
		return err
	}

	return nil
}

func getKogitoCrsInNamespace(namespace string) ([]client.Object, error) {
	var crs []client.Object

	//kogitoRuntimes := &v1beta1.KogitoRuntimeList{}
	//if err := framework.GetObjectsInNamespace(namespace, kogitoRuntimes); err != nil {
	//	return nil, err
	//}
	//for i := range kogitoRuntimes.Items {
	//	crs = append(crs, &kogitoRuntimes.Items[i])
	//}
	//
	//kogitoBuilds := &v1beta1.KogitoBuildList{}
	//if err := framework.GetObjectsInNamespace(namespace, kogitoBuilds); err != nil {
	//	return nil, err
	//}
	//for i := range kogitoBuilds.Items {
	//	crs = append(crs, &kogitoBuilds.Items[i])
	//}
	//
	//kogitoSupportingServices := &v1beta1.KogitoSupportingServiceList{}
	//if err := framework.GetObjectsInNamespace(namespace, kogitoSupportingServices); err != nil {
	//	return nil, err
	//}
	//for i := range kogitoSupportingServices.Items {
	//	crs = append(crs, &kogitoSupportingServices.Items[i])
	//}
	//
	//kogitoInfras := &v1beta1.KogitoInfraList{}
	//if err := framework.GetObjectsInNamespace(namespace, kogitoInfras); err != nil {
	//	return nil, err
	//}
	//for i := range kogitoInfras.Items {
	//	crs = append(crs, &kogitoInfras.Items[i])
	//}

	return crs, nil
}

func cleanupKogitoCrsInNamespace(namespace string) bool {
	crs, err := getKogitoCrsInNamespace(namespace)
	if err != nil {
		framework.GetLogger(namespace).Error(err, "Error getting Kogito CRs.")
		return false
	}

	for _, cr := range crs {
		if err := framework.DeleteObject(cr); err != nil {
			framework.GetLogger(namespace).Error(err, "Error deleting Kogito CR.", "CR name", cr.GetName())
			return false
		}
	}
	return true
}
