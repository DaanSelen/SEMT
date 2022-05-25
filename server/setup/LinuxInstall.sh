#!/bin/bash

chmod 777 ../SEMTS
cp ../SEMTS.service /lib/systemd/system
systemctl enable SEMTS && systemctl restart SEMTS
echo "Script Finished, Showing Status:"
systemctl status SEMTS