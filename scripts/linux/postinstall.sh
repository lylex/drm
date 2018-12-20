#!/bin/bash -ex

DRM_USER=drm
BIN_NAME=drm

CRON="0 0,6,12,18 * * * $BIN_NAME gc"

useradd -M -s /sbin/nologin $DRM_USER
printf '%s\n' "$CRON" >> crontab.tmp
crontab -u $DRM_USER crontab.tmp && rm -f crontab.tmp

echo "alias rm=\"$BIN_NAME\"" >> /etc/profile
