#!/bin/bash
# Copyright 2019 Red Hat, Inc. and/or its affiliates
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

golint ./... | grep -v zz_generated | tee -a golint_errors
if [ -s golint_errors ]  ; then
    code=1
fi
rm -f golint_errors
# see this link to understand why we are using a different go.mod file: https://github.com/golang/go/issues/34506
which golangci-lint > /dev/null || go get -modfile=go.tools.mod github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0
golangci-lint run ./... --enable golint

exit ${code:0}