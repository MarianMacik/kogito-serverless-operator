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

package main

import (
	"testing"

	"github.com/cucumber/godog"
	kogitoExecutor "github.com/kiegroup/kogito-operator/test/pkg/executor"
	kogitoFramework "github.com/kiegroup/kogito-operator/test/pkg/framework"
	"github.com/kiegroup/kogito-operator/test/pkg/meta"
	kogitoSteps "github.com/kiegroup/kogito-operator/test/pkg/steps"
	"github.com/kiegroup/kogito-serverless-operator/testbdd/steps"
)

func TestMain(m *testing.M) {
	// Create kube client
	if err := kogitoFramework.InitKubeClient(meta.GetRegisteredSchema()); err != nil {
		panic(err)
	}

	kogitoExecutor.PreRegisterStepsHook = func(ctx *godog.ScenarioContext, d *kogitoSteps.Data) {
		data := &steps.Data{Data: d}
		data.RegisterAllSteps(ctx)
		//data.RegisterLogsKubernetesObjects(&imgv1.ImageStreamList{}, &v1.KogitoRuntimeList{}, &v1.KogitoBuildList{}, &olmapiv1alpha1.ClusterServiceVersionList{})
	}
	//kogitoExecutor.DisableLogsKogitoCommunityObjects()

	kogitoExecutor.ExecuteBDDTests(nil)
}
