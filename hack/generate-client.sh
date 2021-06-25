#!/bin/bash

set -e

GV="$1"

# require github.com/kubernetes/code-generator v0.19.0 or above
${GOPATH}/src/k8s.io/code-generator/generate-groups.sh "deepcopy,client,informer,lister" github.com/kubesphere/cloud-client github.com/kubesphere/cloud-client/apis "${GV}" -h "$PWD/hack/boilerplate.go.txt"