#!/bin/bash -ex

# Here is the workaround since we have no `deb` or `dmg` like packages for macOS

WORK_DIR="$(dirname "${BASH_SOURCE[0]}")"
CFG_DIR=/etc/drm
DATA_DIR=/usr/local/lib/drm
BLOB_DIR=$DATA_DIR/blob
BIN_DIR=/usr/local/bin/
BIN_NAME=drm

CRON="0 0,6,12,18 * * * $BIN_NAME gc"

[ -d $CFG_DIR ] || mkdir $CFG_DIR

if [ ! -d $DATA_DIR ]; then
  mkdir $DATA_DIR $BLOB_DIR
  chmod -R 777 $DATA_DIR
fi

cp $WORK_DIR/../drm.conf $CFG_DIR/drm.conf
cp $WORK_DIR/../../drm $BIN_DIR

echo "alias rm=\"$BIN_NAME\"" >> /etc/profile

# setup cronjob
crontab -l -u root > crontab.tmp
printf '%s\n' "$CRON" >> crontab.tmp
crontab -u root crontab.tmp && rm -f crontab.tmp
