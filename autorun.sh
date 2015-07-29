#!/bin/bash

set -ex


# Start MongoDB

ps -A | grep mongod > /dev/null
if [ $? = 1 ]; then
	mongod
fi

# RunServer is in $GOPATH/bin which is exported to $PATH
RunServer&

# Start GoOnineJudge
cd $GOPATH/src
restweb run GoOnlineJudge &