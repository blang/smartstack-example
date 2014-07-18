#!/bin/sh

sed -i "s/%PUBLIC_IP%/$PUBLIC_IP/" $1
sed -i "s/%PUBLIC_PORT_9000%/$PUBLIC_PORT_9000/" $1