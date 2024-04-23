#!/bin/bash
while :; do
	rsync -ur . root@138.68.167.74:ep.club --info=progress2
	sleep 0.1
done
