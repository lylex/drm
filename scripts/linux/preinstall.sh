#!/bin/bash -ex

CFG_DIR=/etc/drm
DATA_DIR=/usr/local/lib/drm
BLOB_DIR=$DATA_DIR/blob

[ -d $CFG_DIR ] || mkdir $CFG_DIR

if [ ! -d $DATA_DIR ]; then
  mkdir $DATA_DIR $BLOB_DIR
  chmod -R 777 $DATA_DIR
fi
