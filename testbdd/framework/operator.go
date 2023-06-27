package framework

import (
	"fmt"
	"github.com/kiegroup/kogito-operator/core/operator"
	kogitoFramework "github.com/kiegroup/kogito-operator/test/pkg/framework"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	sonataFlowOperatorTimeoutInMin = 5

	sonataFlowOperatorName                  = "kogito-serverless-operator"
	sonataFlowOperatorDeploymentName        = sonataFlowOperatorName + "-controller-manager"
	sonataFlowOperatorPullImageSecretPrefix = sonataFlowOperatorName + "-dockercfg"
)

// WaitForKogitoOperatorRunning waits for Kogito operator running
func WaitForKogitoOperatorRunning(namespace string) error {
	return kogitoFramework.WaitForOnOpenshift(namespace, "Kogito operator running", sonataFlowOperatorTimeoutInMin,
		func() (bool, error) {
			running, err := IsSonataFlowOperatorRunning(namespace)
			if err != nil {
				return false, err
			}

			// If not running, make sure the image pull secret is present in pod
			// If not present, delete the pod to allow its reconstruction with correct pull secret
			// Note that this is specific to Openshift
			if !running && kogitoFramework.IsOpenshift() {
				podList, err := kogitoFramework.GetPodsWithLabels(namespace, map[string]string{"name": sonataFlowOperatorName})
				if err != nil {
					kogitoFramework.GetLogger(namespace).Error(err, "Error while trying to retrieve Kogito Operator pods")
					return false, nil
				}
				for _, pod := range podList.Items {
					if !kogitoFramework.CheckPodHasImagePullSecretWithPrefix(&pod, sonataFlowOperatorPullImageSecretPrefix) {
						// Delete pod as it has been misconfigured (missing pull secret)
						kogitoFramework.GetLogger(namespace).Info("Kogito Operator pod does not have the image pull secret needed. Deleting it to renew it.")
						err := kogitoFramework.DeleteObject(&pod)
						if err != nil {
							kogitoFramework.GetLogger(namespace).Error(err, "Error while trying to delete Kogito Operator pod")
							return false, nil
						}
					}
				}
			}
			return running, nil
		})
}

// IsSonataFlowOperatorRunning returns whether SonataFlow operator is running
func IsSonataFlowOperatorRunning(namespace string) (bool, error) {
	exists, err := SonataFlowOperatorExists(namespace)
	if err != nil {
		if exists {
			return false, nil
		}
		return false, err
	}

	return exists, nil
}

// SonataFlowOperatorExists returns whether SonataFlow operator exists and is running. If it is existing but not running, it returns true and an error
func SonataFlowOperatorExists(namespace string) (bool, error) {
	kogitoFramework.GetLogger(namespace).Debug("Checking Operator", "Deployment", sonataFlowOperatorDeploymentName, "Namespace", namespace)

	operatorDeployment := &v1.Deployment{}
	namespacedName := types.NamespacedName{Namespace: namespace, Name: sonataFlowOperatorDeploymentName} // done to reuse the framework function
	if exists, err := kogitoFramework.GetObjectWithKey(namespacedName, operatorDeployment); err != nil {
		return false, fmt.Errorf("Error while trying to look for Deploment %s: %v ", sonataFlowOperatorDeploymentName, err)
	} else if !exists {
		return false, nil
	}

	if operatorDeployment.Status.AvailableReplicas == 0 {
		return true, fmt.Errorf("%s Operator seems to be created in the namespace '%s', but there's no available pods replicas deployed ", operator.Name, namespace)
	}

	if operatorDeployment.Status.AvailableReplicas != 1 {
		return false, fmt.Errorf("Unexpected number of pods for %s Operator. Expected %d but got %d ", operator.Name, 1, operatorDeployment.Status.AvailableReplicas)
	}

	return true, nil
}
