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

package operator

const (
	// Name is the name of the Kogito Operator deployed in a namespace
	Name = "kogito-operator"
	// KogitoHomeDir path for Kogito home mounted within the pod of a Kogito Service
	KogitoHomeDir = "/home/kogito"
	// KogitoRuntimeKey ...
	KogitoRuntimeKey = "kogito.kie.org/runtime"
	// KogitoSupportingServiceKey ...
	KogitoSupportingServiceKey = "kogito.kie.org/supporting.service"
)
