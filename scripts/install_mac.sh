#!/bin/bash -ex

# Here is the workaround since we have no `deb` or `dmg` like packages for macOS

WORK_DIR="$(dirname "${BASH_SOURCE[0]}")"
CFG_DIR=/etc/drm
DATA_DIR=/usr/local/lib/drm
BIN_DIR=/usr/local/bin/

[ -d $CFG_DIR ] || mkdir $CFG_DIR

if [ ! -d $DATA_DIR ]; then
  mkdir $DATA_DIR
  chmod -R 777 $DATA_DIR
fi

cp $WORK_DIR/drm.conf $CFG_DIR/drm.conf
cp $WORK_DIR/../drm $BIN_DIR
