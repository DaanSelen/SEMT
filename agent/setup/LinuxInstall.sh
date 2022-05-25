#!/bin/bash

chmod 777 ../SEMTA
cp ../SEMTA.service /lib/systemd/system
systemctl enable SEMTA && systemctl restart SEMTA
echo "Script Finished, Showing Status:"
systemctl status SEMTA