#!/bin/bash

export TYK_GW_STORAGE_HOST=${TYK_GW_STORAGE_HOST:-redis}
export TYK_GW_STORAGE_ADDRS=${TYK_GW_STORAGE_HOST}:6379
export PYTHON_VERSION=${PYTHON_VERSION:-3.9}
export CGO_ENABLED=1

TEST_TIMEOUT=15m

package=$(go list .)

# This is a version to print just the root packages, so we
# could test it with `./...`. Alas, test setups have poor
# cleanups and end up conflicting.

#packages=$(go list ./... | tail -n +2 | sed -e "s|$package/||g" | grep -v '/')

# This is a version that prints all the packages as reported
# by go list, sanitized into local filenames.
packages=$(go list ./... | tail -n +2 | sed -e "s|$package/||g")

# Support passing custom flags (-json, etc.)
OPTS="$@"
if [[ -z "$OPTS" ]]; then
	OPTS="-race -count=1 -v -tags=goplugin"
fi

export PKG_PATH=${GOPATH}/src/github.com/TykTechnologies/tyk

# exit on non-zero exit from go test/vet
set -e

# install gotestsum
go install gotest.tools/gotestsum@latest

# build Go-plugin used in tests
echo "Building go plugin"
go build -race -o ./test/goplugins/goplugins.so -buildmode=plugin ./test/goplugins

for pkg in ${packages}; do
    # local package reference
    pkg=${pkg/github.com\/TykTechnologies\/tyk/.}
    echo "# ${pkg}"

    # sanitize coverage output file
    coveragefile=${pkg/.\//}
    coveragefile=${coveragefile//\//-} # `echo $coveragefile | sed -e 's/\//-/g'`

    echo go test ${OPTS} -timeout ${TEST_TIMEOUT} -coverprofile=${coveragefile}.cov ./${pkg}
    gotestsum --rerun-fails=3 --raw-command go test ${OPTS} -json -timeout ${TEST_TIMEOUT} -coverprofile=${coveragefile}.cov -cover ./${pkg}
done
