#!/bin/sh
# run docker image command
docker run -p 8001:80 -v /data/logs/kcores.com/:/data/repo/kcores.com/logs/ -d kcores.com:0.0.3