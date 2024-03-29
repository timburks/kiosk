#
# Copyright 2018 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
lite:
	go get ./...
	go install ./...
	
all:
	./tools/GENERATE-RPC.sh
	./tools/GENERATE-GRPC.sh
	./tools/GENERATE-GAPIC.sh
	./tools/GENERATE-CLI.sh
	./tools/GENERATE-DESCRIPTORS.sh
	./tools/GENERATE-DOCS.sh
	go get ./...
	go install ./...

clean:
	rm -rf protos/api-common-protos
	rm -rf generated
	rm -rf gapic/*.go
	rm -rf cmd/kctl

