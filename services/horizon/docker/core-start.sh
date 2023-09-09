#!/usr/bin/env bash

set -e
set -x

source /etc/profile
# work within the current docker working dir
if [ ! -f "./gramr.cfg" ]; then
   cp /gramr.cfg ./
fi   

echo "using config:"
cat gramr.cfg

# initialize new db
gramr new-db

if [ "$1" = "standalone" ]; then
  # initialize for new history archive path, remove any pre-existing on same path from base image
  rm -rf ./history
  gramr new-hist vs

  # serve history archives to horizon on port 1570
  pushd ./history/vs/
  python3 -m http.server 1570 &
  popd
fi

exec gramr run --console
