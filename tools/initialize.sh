#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
echo $DIR
if [ ! -d "/tmp/hanuctl/" ]
then
	mkdir /tmp/hanuctl/
fi
EXAMPLE="$(dirname "$DIR")"/examples/*
cp $EXAMPLE /tmp/hanuctl/
