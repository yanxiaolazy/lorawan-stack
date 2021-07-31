# Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: default
default: init

SHELL = bash

include tools/mage/mage.make

build: build-init build-latest build-version

build-init:
	@echo "Version is " ${VERSION} && \
	make init && \
	tools/bin/mage js:build && \
	goreleaser release -f .goreleaser.release.yml --snapshot --rm-dist

build-latest:
	docker build -t isrookie/lorawan-stack:latest . && \
	docker push isrookie/lorawan-stack:latest

build-version:
	docker build -t isrookie/lorawan-stack:${VERSION} . && \
	docker push isrookie/lorawan-stack:${VERSION}