#!/bin/bash -ex

DRM_USER=drm

crontab -l -u $DRM_USER
userdel -r $DRM_USER

sed -i '/alias rm=/d' /etc/profile
