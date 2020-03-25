// Copyright 2019 Red Hat, Inc. and/or its affiliates
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

package steps

import (
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/gherkin"
	"github.com/kiegroup/kogito-cloud-operator/pkg/infrastructure"
	"github.com/kiegroup/kogito-cloud-operator/test/framework"
)

func registerGraphQLSteps(s *godog.Suite, data *Data) {
	s.Step(`^GraphQL request on service "([^"]*)" is successful within (\d+) minutes with path "([^"]*)" and query:$`, data.graphqlRequestOnServiceIsSuccessfulWithinMinutesWithPathAndQuery)
	s.Step(`^GraphQL request on Data Index service returns ProcessInstances processName "([^"]*)" within (\d+) minutes$`, data.graphqlRequestOnDataIndexReturnsProcessInstancesProcessNameWithinMinutes)
}

func (data *Data) graphqlRequestOnServiceIsSuccessfulWithinMinutesWithPathAndQuery(serviceName string, timeoutInMin int, path string, query *gherkin.DocString) error {
	framework.GetLogger(data.Namespace).Debugf("graphqlRequestOnServiceWithPathAndBodyIsSuccessfulWithinMinutes with service %s, path %s, query %s and timeout %d", serviceName, path, query, timeoutInMin)
	routeURI, err := framework.WaitAndRetrieveRouteURI(data.Namespace, serviceName)
	if err != nil {
		return err
	}
	var response interface{}
	return framework.WaitForSuccessfulGraphQLRequest(data.Namespace, routeURI, path, query.Content, timeoutInMin, response, nil)
}

func (data *Data) graphqlRequestOnDataIndexReturnsProcessInstancesProcessNameWithinMinutes(processName string, timeoutInMin int) error {
	serviceName := infrastructure.DefaultDataIndexName
	query := getProcessInstancesNameQuery
	path := "graphql"

	framework.GetLogger(data.Namespace).Debugf("graphqlProcessNameRequestOnDataIndexIsSuccessfulWithinMinutes with service %s, path %s, query %s and timeout %d", serviceName, path, query, timeoutInMin)
	routeURI, err := framework.WaitAndRetrieveRouteURI(data.Namespace, serviceName)
	if err != nil {
		return err
	}
	response := GraphqlDataIndexProcessInstanceQueryResponse{}
	return framework.WaitForSuccessfulGraphQLRequest(data.Namespace, routeURI, path, query, timeoutInMin, &response, func(response interface{}) (bool, error) {
		resp := response.(*GraphqlDataIndexProcessInstanceQueryResponse)
		for _, processInstance := range resp.ProcessInstances {
			if processInstance.ProcessName == processName {
				return true, nil
			}
		}
		return false, nil
	})
}
