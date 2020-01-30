#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
echo $DIR
if [ ! -d "/tmp/tacoctl/" ]
then
	mkdir /tmp/tacoctl/
fi
EXAMPLE="$(dirname "$DIR")"/examples/*
cp $EXAMPLE /tmp/tacoctl/
