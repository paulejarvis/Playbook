#!/usr/bin/env bash
# TODO(hdh) use MAKE instead of this...
# latest is not a version!
docker pull dgraph/dgraph@sha256:115d77bf2098b183b3b6ac4d55c56c6778d54c9ec72ff17cf44127b9414ae794

docker-compose up --detach