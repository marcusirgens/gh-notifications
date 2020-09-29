#!/bin/sh

apk add -u curl;
[ -d "./bin" ] || mkdir "./bin"
curl -sL https://taskfile.dev/install.sh | sh

mv bin/task /usr/local/bin/task

# Delete ./bin if it has no files
[ "$(find ./bin -type f | wc -l)" -gt 0 ] || rm -rf bin
