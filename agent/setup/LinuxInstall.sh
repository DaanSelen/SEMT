#!/bin/bash

if [[ ! $(sudo echo 0) ]]; then echo "Please run as sudo"; fi
chmod 777 ../SEMTA
cp ../SEMTA.service /lib/systemd/system
systemctl enable SEMTA && systemctl restart SEMTA
echo "Script Finished, Showing Status:"
systemctl status SEMTA