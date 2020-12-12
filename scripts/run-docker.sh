#!/bin/sh
# run docker image command

# check param
if [ ! $1 ]; then
    echo "Please input docker tag, e.g. kcores.com:0.0.3-b1" 
fi

if [ ! $2 ]; then
    echo "Please input instance name, e.g. kcores.com" 
fi

# run
docker run --name $2 -p 8001:80 -v /data/logs/kcores.com/:/data/repo/kcores.com/logs/ -d $1 