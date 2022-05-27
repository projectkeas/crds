#!/bin/bash

# this is a workaround for WSL installation
export PATH=$PATH:/usr/local/go/bin
execDir=./modules/code-generator

echo "=== Auto-Generating Client ==="
$execDir/generate-groups.sh all github.com/projectkeas/crds/pkg/client github.com/projectkeas/crds/pkg/apis keas.io:v1alpha1 --go-header-file ./hack/boilerplate.go.txt --output-base ./out

cp -r ./out/github.com/projectkeas/crds/pkg/* ./pkg/
rm -rf ./out

echo "=== Updating CRD's ==="
go run sigs.k8s.io/controller-tools/cmd/controller-gen crd paths=./pkg/apis/keas.io/v1alpha1 output:dir=manifests
