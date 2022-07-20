#!/bin/bash

if [[ ! $(sudo echo 0) ]]; then echo "Please run as sudo"; fi
chmod 777 ../SEMTS
cp ../SEMTS.service /lib/systemd/system
systemctl enable SEMTS && systemctl restart SEMTS
echo "Script Finished, Showing Status:"
systemctl status SEMTS