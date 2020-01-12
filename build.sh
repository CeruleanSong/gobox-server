#!/bin/bash

if [[ $1 = "run" ]]
then
	echo "running..."

	./dist/gopy

	echo "done..."
elif [[ $1 = "build" ]]
then
	echo "building..."
	
	[ -e "./dist" ] && rm -rf ./dist
	cd src
	go build -o ../dist/gopy

	echo "done..."
else
	./build.sh build
	./build.sh run
fi