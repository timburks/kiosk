#!/bin/sh
# Copyright 2022 Google LLC. All Rights Reserved.
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
# Use this script to regenerate the Protocol Buffer and gRPC files
# needed to build the example.

if [ ! -d "third_party/api-common-protos" ]; then
  curl -L -O https://github.com/googleapis/api-common-protos/archive/main.zip
  unzip main.zip
  rm -f main.zip
  mkdir -p third_party
  mv api-common-protos-main third_party/api-common-protos
fi

export PROJECT=".."

# This points to the .proto files distributed with protoc.
export PROTOC="/usr/local/include"

# This is a third_party directory containing .proto files used by many APIs.
export COMMON="third_party/api-common-protos"

mkdir -p lib/src/generated

echo "Generating Dart support code."
protoc \
  ${PROJECT}/examples/kiosk/v1/kiosk.proto \
  ${PROTOC}/google/protobuf/empty.proto \
  ${PROTOC}/google/protobuf/timestamp.proto \
  ${COMMON}/google/type/latlng.proto  \
  -I ${PROJECT} \
  -I ${PROTOC} \
  -I ${COMMON} \
  --dart_out=grpc:lib/src/generated

