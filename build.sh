#!/bin/bash

if [[ $1 = "run" ]]
then
	echo "running..."

	./dist/gobox

	echo "done..."
elif [[ $1 = "build" ]]
then
	echo "building..."
	
	[ -e "./dist" ] && rm -rf ./dist
	cd src
	go build -o ../dist/gobox

	echo "done..."
else
	./build.sh build
	./build.sh run
fi