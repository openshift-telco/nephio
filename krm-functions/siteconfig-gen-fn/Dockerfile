#  Copyright 2023 The Nephio Authors.
#
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

FROM golang:1.20-alpine
ENV CGO_ENABLED=0
WORKDIR /go/src/
COPY . krm-functions/
WORKDIR krm-functions
RUN go install
RUN go build -o /usr/local/bin/function ./

FROM quay.io/openshift-kni/ztp-site-generator:4.13

# Given the binary requires libc, we need to use the "base" distroless and not the "static"

#$ ldd /kustomize/plugin/ran.openshift.io/v1/siteconfig/SiteConfig
#	linux-vdso.so.1 (0x00007ffe4eaba000)
#	libpthread.so.0 => /lib64/libpthread.so.0 (0x00007f5545b8d000)
#	libc.so.6 => /lib64/libc.so.6 (0x00007f55457c8000)
#	/lib64/ld-linux-x86-64.so.2 (0x00007f5545dad000)

FROM gcr.io/distroless/base:latest
COPY --from=0 /usr/local/bin/function /usr/local/bin/function
COPY --from=1 /home/ztp/extra-manifest /usr/local/bin/extra-manifest
COPY --from=1 /kustomize/plugin/ran.openshift.io/v1/siteconfig/SiteConfig /usr/local/bin/SiteConfig

ENTRYPOINT ["function"]
