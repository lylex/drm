#!/bin/bash -ex

echo "preinstall drm" 

[ -d /etc/drm ] || mkdir /etc/drm

if [ ! -d /var/lib/drm ]; then
  mkdir /var/lib/drm
  chmod -R 666 /var/lib/drm
fi
