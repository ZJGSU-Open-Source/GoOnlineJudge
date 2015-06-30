#!/bin/sh

# Install Ubuntu.
sed -i 's/# \(.*multiverse$\)/\1/g' /etc/apt/sources.list
apt-get update
apt-get -y upgrade
apt-get install -y build-essential git vim wget flex
rm -rf /var/lib/apt/lists/*

# Install Golang.
mkdir -p /home/acm/goroot
wget https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz
tar xzf go1.4.2.linux-amd64.tar.gz
cp -r go/* /home/acm/goroot/
mkdir -p /home/acm/go/src /home/acm/go/pkg /home/acm/go/bin

# Set environment variables for Golang.
export GOROOT=/home/acm/goroot
export GOPATH=/home/acm/go
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

# Install MongoDB.
apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 7F0CEB10
echo 'deb http://downloads-distro.mongodb.org/repo/ubuntu-upstart dist 10gen' > /etc/apt/sources.list.d/mongodb.list
apt-get update
apt-get install -y mongodb
mkdir -p $GOPATH/Data
rm -rf /var/lib/apt/lists/*

# Get OJ Source Code
mkdir -p $GOPATH/src/ProblemData
mkdir -p $GOPATH/src/run
mkdir -p $GOPATH/src/log
go get gopkg.in/mgo.v2
go get github.com/djimenez/iconv-go
git clone https://github.com/ZJGSU-Open-Source/GoOnlineJudge.git $GOPATH/src/GoOnlineJudge
git clone https://github.com/ZJGSU-Open-Source/RunServer.git $GOPATH/src/RunServer
git clone https://github.com/ZJGSU-Open-Source/vjudger.git $GOPATH/src/vjudger
git clone https://github.com/sakeven/restweb.git $GOPATH/src/restweb

# Build OJ
cd $GOPATH/src/restweb
cd restweb
go install
cd $GOPATH/src
restweb build GoOnlineJudge
cd $GOPATH/src/RunServer
./make.sh

echo
echo ----------
echo installed.
echo ----------
echo

# Run MongoDB, GoOnlineJudge, RunServer
mongod --dbpath /home/acm/go/Data --logpath /home/acm/go/Data/mongo.log
cd $GOPATH/src
restweb run GoOnlineJudge &
cd $GOPATH/src/GoOnlineJudge
./RunServer &
