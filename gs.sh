#!/bin/bash
# ------------------------------------------------------------------
# [Author] Bryan McGrane (bmcgrane@sensoryinc.com)
# [Title] Go SDK Helper Script
# ------------------------------------------------------------------
VERSION=0.10.0
SUBJECT=gs
USAGE="Usage: gs [COMMAND]"

COMMANDS="
    Commands:\n
    genproto | gp\t\t Generate Proto Files\n
    releaseversion | rv\t [vX.Y.Z] Sets and releases the internal project version\n
    test | t\t\t Run Tests\n
"

# --- Initialization -----------------------------------------------
set -euo pipefail
export PATH=$PATH:/usr/local/go/bin

# --- Functions ----------------------------------------------------
print_helper() {
  echo
  echo ${USAGE}
  echo
  echo -e ${COMMANDS}
}

throw_not_implemented_error() {
  echo "This feature is not yet implemented!"
  exit 125;
}

generate_proto_files() {
  echo "Generating proto files in"

  mkdir -p ${GENERATED_CODE_DIRECTORY}/generated
  for x in $(find ./proto -iname "*.proto");
  do
    if [[ "$x" == *'validate.proto' ]]; then
      continue
    fi

    protoc \
      --proto_path=./proto \
      --go_out=paths=source_relative:${GENERATED_CODE_DIRECTORY}/generated \
      --go-grpc_out=${GENERATED_CODE_DIRECTORY}/generated \
      --go-grpc_opt=paths=source_relative \
      --validate_out="lang=go,paths=source_relative:${GENERATED_CODE_DIRECTORY}/generated" \
      $x;

    echo "Generated grpc code for $x";
  done

  for x in $(find ./pkg/generated -iname "*.go");
  do
    if [ ${OPERATING_SYSTEM} == "Darwin" ]; then
      sed -i '' "s gitlab.com/sensory-cloud/server/titan.git/pkg/api github.com/Sensory-Cloud/go-sdk/pkg/generated g" $x
    else
      sed -i "s/BuildVersion string = \"[^\"]*\"/BuildVersion string = \"${version}\"/" ${VERSION_FILE}
    fi
  done
}

run_integration_tests() {
  echo "Running Integration Tests"
  # Run all tests
  go test -bench=. ./test/... -benchtime=5x

  # Run face auth test x500
  # go test -benchmem -run=^$ -bench ^BenchmarkAuthenticateWithLiveness$ ./test/integration/v1/video -benchtime=500x

  # Run voice TSSV auth test x500
  # go test -benchmem -run=^$ -bench ^BenchmarkTssvAuthenticate$ ./test/integration/v1/audio -benchtime=500x

  # Run voice TSSV wakeword test x500
  # go test -benchmem -run=^$ -bench ^BenchmarkTssvEvent$ ./test/integration/v1/audio -benchtime=500x

   # Run voice TNL wakeword test x500
  # go test -benchmem -run=^$ -bench ^BenchmarkTnlEvent$ ./test/integration/v1/audio -benchtime=10x

  # Run STT Transcribe test x500
  # go test -benchmem -run=^$ -bench ^BenchmarkTranscribe$ ./test/integration/v1/audio -benchtime=500x
}

set_version() {
  version=$1
  regex_version='^v[0-9]+\.[0-9]+\.[0-9]+$'

  if [[ ! ${version} =~ ${regex_version} ]]; then
    echo "Version string should be of the format v{Major}.{Minor}.{Trivial} ex: v1.2.3"
    exit 1
  fi

  # Check if version exists
  git fetch --tags
  if [ $(git tag -l "${version}") ]; then
    echo "Version ${version} already exists. Exiting."
    exit 1
  fi

  echo "Updating version file to ${1}"
  if [ ${OPERATING_SYSTEM} == "Darwin" ]; then
    sed -i '' "s/BuildVersion string = \"[^\"]*\"/BuildVersion string = \"${version}\"/" ${VERSION_FILE}
    sed -i '' "s/BuildTime    string = \"[^\"]*\"/BuildTime    string = \"$(date)\"/" ${VERSION_FILE}
  else
    sed -i "s/BuildVersion string = \"[^\"]*\"/BuildVersion string = \"${version}\"/" ${VERSION_FILE}
    sed -i "s/BuildTime    string = \"[^\"]*\"/BuildTime    string = \"$(date)\"/" ${VERSION_FILE}
  fi

  git commit -am "Release [${version}]"
  git tag ${version}
  git push --atomic origin HEAD ${version}
}

# --- Options processing -------------------------------------------
if [ $# == 0 ] ; then
    print_helper
    exit 1;
fi

while getopts ":vh" optname
  do
    case "$optname" in
      "v")
        echo "Version $VERSION"
        exit 0;
        ;;
      "h")
        print_helper
        exit 0;
        ;;
      "?")
        echo "Unknown option $OPTARG"
        exit 0;
        ;;
      ":")
        echo "No argument value for option $OPTARG"
        exit 0;
        ;;
      *)
        echo "Unknown error while processing options"
        exit 0;
        ;;
    esac
  done

shift $(($OPTIND - 1))

# --- Constants -----------------------------------------
VERSION_FILE="${PWD}/pkg/version/version.go"
CURRENT_VERSION=`sed -n 's/^	BuildVersion string = "\(.*\)"/\1/p' < ${VERSION_FILE}`

GENERATED_CODE_DIRECTORY=./pkg

export OPERATING_SYSTEM=$(uname -s)
export ENV_FILE_PATH="${PWD}/.env"

# --- Process Environment -----------------------------------------
CURRENT_BRANCH=${CI_COMMIT_BRANCH:-$(git rev-parse --symbolic-full-name --abbrev-ref HEAD)}
GO_ENV=local
LATEST_TAG="latest"
CURRENT_VERSION_SUFFIX=""
if [ ${CURRENT_BRANCH} = "master" ]; then
  echo "Production enabled"
  GO_ENV=production
elif [ ${CURRENT_BRANCH} = "stage" ]; then
  echo "Staging enabled"
  GO_ENV=staging
elif [ ${CURRENT_BRANCH} = "develop" ]; then
  echo "Development enabled"
  GO_ENV=development
else
  LATEST_TAG="next"
  CURRENT_VERSION_SUFFIX="-pre"
fi

export GO_ENV
export GO_PATH=$(go env GOPATH)
export GOPATH=$GO_PATH
export PATH="$PATH:$GO_PATH/bin"
export DOCKER_BUILDKIT=1
# --- Body --------------------------------------------------------
case "$1" in

  "genproto"|"gp")
    generate_proto_files
    exit 0;
  ;;

  "test"|"t")
    go test -cover $(go list ./... | grep -v -f ${PWD}/test/test_exclusion_list.txt)
    exit 0;
  ;;

  "releaseversion"|"rv")
    if [[ $# -eq 1 ]]; then
        echo "This commands a version string to be specified ex: \"./ti.sh rv v1.2.3\""
        exit 1;
    fi

    set_version $2
    exit 0;
  ;;

  *)
    echo "Unknown command: $1"
    print_helper
    exit 0;
    ;;

esac
# -----------------------------------------------------------------
