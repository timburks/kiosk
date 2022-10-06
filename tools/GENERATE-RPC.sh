#!/bin/bash
#
# Copyright 2021 Google LLC. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

set -e

. tools/PROTOS.sh
clone_common_protos

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

for proto in ${ALL_PROTOS[@]}; do
	echo "Generating Go types for $proto"
	protoc $proto \
	--proto_path='.' \
	--proto_path=$COMMON_PROTOS_PATH \
	--go_opt='module=github.com/googleapis/kiosk' \
	--go_out='.'
done
