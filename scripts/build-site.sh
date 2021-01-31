#!/bin/sh
# build docker image command

# build content
cd ./content-builder/src/
go build 
./content-builder.exe
cd ../../


# build markdown
cd ./markdown-builder
php ./generate-html.php 
cd ../

# copy generated content
cp ./generated/* ./