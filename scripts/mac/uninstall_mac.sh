#!/bin/bash -ex

# Here is the workaround since we have no `deb` or `dmg` like packages for macOS

BIN_NAME=drm

sed -i '' '/alias rm=/d' /etc/profile

# remove cronjob
TEMP_FILE=crontab.tmp
crontab -l -u root > $TEMP_FILE
sed -i '' "/$BIN_NAME gc/d" $TEMP_FILE
crontab -u root $TEMP_FILE && rm -f $TEMP_FILE
