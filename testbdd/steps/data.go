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

package steps

import (
	"fmt"
	"github.com/cucumber/godog"
	"github.com/kiegroup/kogito-operator/test/pkg/framework"
	"github.com/kiegroup/kogito-operator/test/pkg/steps"
	"strconv"
	"time"
)

// Data contains all data needed by Gherkin steps to run
type Data struct {
	*steps.Data
}

// RegisterAllSteps register all steps available to the test suite
func (data *Data) RegisterAllSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^Asked to print out the name "([^"]*)"$`, data.askedToPrintOutTheName)
	ctx.Step(`^Print the name "([^"]*)"$`, data.printTheName)
	ctx.Step(`^Wait (\d+) seconds?$`, data.waitSeconds)
	registerOperatorSteps(ctx, data)
	registerPlatformSteps(ctx, data)
	registerSonataFlowSteps(ctx, data)

	//registerKogitoBuildSteps(ctx, data)
	//registerKogitoDeployFilesSteps(ctx, data)
	//registerKafkaSteps(ctx, data)
	//registerKogitoInfraSteps(ctx, data)
	//registerKogitoRuntimeSteps(ctx, data)
	//registerOpenShiftSteps(ctx, data)
	//registerOperatorSteps(ctx, data)
}

func (data *Data) askedToPrintOutTheName(arg1 string) error {
	fmt.Printf("Hello Print Out The Name %s", arg1)
	return nil
}

func (data *Data) printTheName(arg1 string) error {
	fmt.Printf("Hello The Name %s", arg1)
	return nil
}

func (data *Data) waitSeconds(seconds int) error {
	framework.GetMainLogger().Info("Waiting for " + strconv.Itoa(seconds) + " s")
	_ = <-time.After(time.Duration(seconds) * time.Second)
	return nil
}
