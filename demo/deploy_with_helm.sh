#!/bin/bash

work(){
    echo "deploying $VERSION"
    pushd go-openapi-$VERSION/helm
    helm install go-openapi-$VERSION . --namespace=apps --set=image.tag=$VERSION
    popd
}
declare -a arr=("v1" "v2")
for i in "${arr[@]}"
do
    VERSION=$i
    work
done

## Deploy client
pushd go-openapi-client/helm
helm install go-openapi-client . --namespace=apps
popd
