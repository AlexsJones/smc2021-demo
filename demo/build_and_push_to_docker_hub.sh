#!/bin/bash

work(){
    echo "building $VERSION"
    pushd go-openapi-$VERSION
    docker build . -t tibbar/go-openapi:$VERSION
    docker push tibbar/go-openapi:$VERSION
    popd
}
declare -a arr=("v1" "v2")
for i in "${arr[@]}"
do
    VERSION=$i
    work
done

## Build client
pushd go-openapi-client
docker build . -t tibbar/go-openapi-client:v1
docker push tibbar/go-openapi-client:v1
popd
