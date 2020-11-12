#!/bin/sh
# build docker image command

# check param
if [ ! $1 ]; then
    echo "Please input docker tag, e.g. kcores.com:0.0.3-b1" 
fi

# build
docker build ./ -t $1