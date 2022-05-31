#!/bin/bash

# This script only works with go v1.17 or above

set -o errexit
set -o nounset
set -o pipefail

# this is a workaround for WSL installation
export PATH=$PATH:/usr/local/go/bin
export GOPATH=~/go

# use this if a reset/version upgrade has been done
#rm -rf ~/go/src

LH_MANAGER_DIR="github.com/projectkeas/crds"
OUTPUT_DIR="${LH_MANAGER_DIR}/pkg/client"
APIS_DIR="${LH_MANAGER_DIR}/pkg/apis"
GROUP_VERSION="keas.io:v1alpha1"
CODE_GENERATOR_VERSION="v0.24.1"
CRDS_DIR="./manifests"
CONTROLLER_TOOLS_VERSION="v0.9.0"
KUSTOMIZE_VERSION="kustomize/v4.5.5"

echo "=== Performing Initial Cleanup ==="
#rm -rf ${GOPATH}/src
rm -rf ./pkg/client
rm -f ./pkg/apis/keas.io/v1alpha1/zz_generated.deepcopy.go
rm -f ./manifests/keas.io_*.yaml

mkdir -p ${GOPATH}/src/k8s.io

echo "=== Checking Dependencies ==="
# https://github.com/kubernetes/code-generator/blob/v0.18.0/generate-groups.sh
if [[ ! -d "${GOPATH}/src/k8s.io/code-generator" ]]; then
    echo "  Installing code-generator"
	pushd ${GOPATH}/src/k8s.io
	git clone -b ${CODE_GENERATOR_VERSION} https://github.com/kubernetes/code-generator 2>/dev/null || true
	popd
fi

# https://github.com/kubernetes-sigs/controller-tools/tree/v0.7.0/cmd/controller-gen
if [[ ! -d "${GOPATH}/src/k8s.io/controller-tools" ]]; then
  echo "  Installing controller-gen"
  pushd ${GOPATH}/src/k8s.io
  git clone -b ${CONTROLLER_TOOLS_VERSION} https://github.com/kubernetes-sigs/controller-tools.git 2>/dev/null || true
  cd controller-tools/cmd/controller-gen
  go install .
  popd
fi

# https://github.com/kubernetes-sigs/kustomize/tree/kustomize/v3.10.0/kustomize
if [[ ! -d "${GOPATH}/src/k8s.io/kustomize" ]]; then
    echo "  Installing kustomize"
    pushd ${GOPATH}/src/k8s.io
	git clone -b ${KUSTOMIZE_VERSION} https://github.com/kubernetes-sigs/kustomize.git 2>/dev/null || true
	cd kustomize/kustomize
	go install .
	popd
fi

echo "=== Generating new API types ==="
${GOPATH}/src/k8s.io/code-generator/generate-groups.sh \
  all \
  ${OUTPUT_DIR} \
  ${APIS_DIR} \
  ${GROUP_VERSION} \
  --output-base ./out \
  --go-header-file ./hack/boilerplate.go.txt \
  $@

cp -r ./out/github.com/projectkeas/crds/pkg/* ./pkg/
rm -rf ./out

echo "=== Updating CRD's ==="
~/go/bin/controller-gen crd paths=${APIS_DIR}/... output:crd:dir=${CRDS_DIR}
pushd ${CRDS_DIR}
kustomize create --autodetect 2>/dev/null || true
kustomize edit add label longhorn-manager: 2>/dev/null || true
if [ -e "${CRDS_DIR}/patches" ]; then
  find "${CRDS_DIR}/patches" -type f | xargs -i sh -c 'kustomize edit add patch --path {}'
fi
popd
