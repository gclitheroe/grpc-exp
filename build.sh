#!/bin/bash -e

# Builds Docker images for the arg list.  These must be project directories
# where this script is executed.
#
# Builds a statically linked executable and adds it to the container.
# Adds the assets dir from each project to the container e.g., origin/assets
# It is not an error for the assets dir to not exist.
# Any assets needed by the application should be read from the assets dir
# relative to the executable. 
#
# usage: ./build.sh project [project]

if [ $# -eq 0 ]; then
    echo Error: please supply a project to build. Usage: ./build.sh project [project]
    exit 1
fi

# code will be compiled in this container
BUILD_CONTAINER=golang:1.7.0-alpine

DOCKER_TMP=docker-build-tmp

mkdir -p $DOCKER_TMP
chmod +s $DOCKER_TMP

rm -rf $DOCKER_TMP/*

VERSION='git-'`git rev-parse --short HEAD`

# Prefix for the logs.  Will be ignored if log/logentries is not used.
BUILD='-X github.com/GeoNet/haz/vendor/github.com/GeoNet/log/logentries.Prefix='$VERSION

# The current working dir to use in GOBIN etc e.g., geonet-web
CWD=${PWD##*/}

# Assemble common resource for ssl and timezones from the build container
docker run --rm -v "$PWD":"$PWD"  ${BUILD_CONTAINER} \
	apk add --update ca-certificates tzdata; \
	mkdir -p "$PWD"/${DOCKER_TMP}/etc/ssl/certs; \
	mkdir -p "$PWD"/${DOCKER_TMP}/usr/share; \
	cp /etc/ssl/certs/ca-certificates.crt "$PWD"/${DOCKER_TMP}/etc/ssl/certs; \
	cp -Ra /usr/share/zoneinfo "$PWD"/${DOCKER_TMP}/usr/share

# Assemble common resource for user.
echo "nobody:x:65534:65534:Nobody:/:" > ${DOCKER_TMP}/etc/passwd

for i in "$@"
do
	docker run -e "GOBIN=/usr/src/go/src/github.com/gclitheroe/${CWD}/${DOCKER_TMP}" -e "GOPATH=/usr/src/go" -e "CGO_ENABLED=0" -e "GOOS=linux" -e "BUILD=$BUILD" --rm \
		-v "$PWD":/usr/src/go/src/github.com/gclitheroe/${CWD} \
		-w /usr/src/go/src/github.com/gclitheroe/${CWD} ${BUILD_CONTAINER} \
		go install -a  -ldflags "${BUILD}" -installsuffix cgo ./${i}

		rm -rf $DOCKER_TMP/assets
		mkdir $DOCKER_TMP/assets
		rsync --archive --quiet --ignore-missing-args  ${i}/assets docker-build-tmp/

        # Add a default Dockerfile

		rm -f $DOCKER_TMP/Dockerfile

		echo "FROM scratch" > docker-build-tmp/Dockerfile
		echo "ADD ./ /" >> docker-build-tmp/Dockerfile
		echo "USER nobody" >> docker-build-tmp/Dockerfile
		echo "EXPOSE 8443" >> docker-build-tmp/Dockerfile
		echo "CMD [\"/${i}\"]" >> docker-build-tmp/Dockerfile

        # If a project specifies a Dockerfile then copy it over the top of the default file.

        rsync --ignore-missing-args ${i}/Dockerfile docker-build-tmp/

		docker build -t quay.io/gclitheroe/${i}:$VERSION -f docker-build-tmp/Dockerfile docker-build-tmp

        docker tag quay.io/gclitheroe/${i}:$VERSION quay.io/gclitheroe/${i}:latest

		rm -f $DOCKER_TMP/$i
done
