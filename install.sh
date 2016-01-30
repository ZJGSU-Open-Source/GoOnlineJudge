#!/bin/sh

# Install Ubuntu.
sed -i 's/# \(.*multiverse$\)/\1/g' /etc/apt/sources.list
apt-get update
apt-get -y upgrade
apt-get install -y build-essential mongodb flex openjdk-7-jdk
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
export OJ_HOME=$GOPATH/src
export DATA_PATH=$GOPATH/Data
export LOG_PATH=$OJ_HOME/log
export RUN_PATH=$OJ_HOME/run
export JUDGE_HOST="http://127.0.0.1:8888"
export MONGODB_PORT_27017_TCP_ADDR=127.0.0.1
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

# Install MongoDB.
apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 7F0CEB10
echo 'deb http://downloads-distro.mongodb.org/repo/ubuntu-upstart dist 10gen' > /etc/apt/sources.list.d/mongodb.list
apt-get update
apt-get install -y mongodb
rm -rf /var/lib/apt/lists/*

# Get OJ Source Code
mkdir -p $OJ_HOME/ProblemData
mkdir -p $OJ_HOME/run
mkdir -p $OJ_HOME/log
go get gopkg.in/mgo.v2
go get github.com/djimenez/iconv-go
git clone https://github.com/ZJGSU-Open-Source/GoOnlineJudge.git $GOPATH/src/GoOnlineJudge
git clone https://github.com/ZJGSU-Open-Source/RunServer.git $GOPATH/src/RunServer
git clone https://github.com/ZJGSU-Open-Source/vjudger.git $GOPATH/src/vjudger
git clone https://github.com/sakeven/restweb.git $GOPATH/src/restweb

# Build OJ
cd $OJ_HOME/restweb
cd restweb
go install
cd $OJ_HOME
restweb build GoOnlineJudge
cd $OJ_HOME/RunServer
./make.sh

echo
echo ----------
echo installed.
echo ----------
echo

# Run MongoDB, GoOnlineJudge, RunServer
mongod --dbpath /home/acm/go/Data --logpath /home/acm/go/Data/mongo.log
cd $OJ_HOME/
restweb run GoOnlineJudge &
cd $GOPATH/src/GoOnlineJudge
RunServer &
