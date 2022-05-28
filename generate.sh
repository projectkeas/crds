#!/bin/bash

# this is a workaround for WSL installation
export PATH=$PATH:/usr/local/go/bin
execDir=./modules/code-generator

echo "=== Auto-Generating Client ==="
rm -rf ./pkg/client
rm -f ./pkg/apis/keas.io/v1alpha1/zz_generated.deepcopy.go

$execDir/generate-groups.sh all github.com/projectkeas/crds/pkg/client github.com/projectkeas/crds/pkg/apis keas.io:v1alpha1 --go-header-file ./hack/boilerplate.go.txt --output-base ./out

cp -r ./out/github.com/projectkeas/crds/pkg/* ./pkg/
rm -rf ./out

echo "=== Updating CRD's ==="
rm -f ./manifests/keas.io_*.yaml
go run sigs.k8s.io/controller-tools/cmd/controller-gen crd paths=./pkg/apis/keas.io/v1alpha1 output:dir=manifests

echo "=== Updating Kustomization File ==="
echo "resources:" > manifests/kustomization.yml
for path in ./manifests/*.yaml; do
    [[ $path = "./manifests/kustomization.yml" ]] && continue;
    [[ $path = "./manifests/Test.yaml" ]] && continue;
    echo "- \"${path:12}\"" >> manifests/kustomization.yml
done

echo "- \"Test.yaml\"" >> manifests/kustomization.yml
