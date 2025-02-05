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

package workflowdef

import (
	"context"
	"encoding/json"

	"github.com/kiegroup/kogito-serverless-operator/log"

	"k8s.io/klog/v2"

	operatorapi "github.com/kiegroup/kogito-serverless-operator/api/v1alpha08"
)

// GetJSONWorkflow return a Kogito compliant JSON format workflow as bytearray give a specific workflow CR
func GetJSONWorkflow(workflowCR *operatorapi.SonataFlow, ctx context.Context) ([]byte, error) {
	// apply workflow metadata
	workflow, err := operatorapi.ToCNCFWorkflow(workflowCR, ctx)
	if err != nil {
		klog.V(log.E).ErrorS(err, "Failed converting SonataFlow into Workflow")
		return nil, err
	}
	jsonWorkflow, err := json.Marshal(workflow)
	if err != nil {
		klog.V(log.E).ErrorS(err, "Failed converting SonataFlow into JSON")
		return nil, err
	}
	return jsonWorkflow, nil
}
