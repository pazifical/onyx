source "onyx.env"

docker build -t "pazifical/onyx:$ONYX_VERSION" .
docker push "pazifical/onyx:$ONYX_VERSION"
