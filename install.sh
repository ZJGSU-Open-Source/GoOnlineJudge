# gopath. dir for problemdata, code and log
export GOPATH=/home/acm/go
sudo mkdir -p $GOPATH/src
sudo mkdir -p $GOPATH/ProblemData
sudo mkdir -p $GOPATH/run
sudo mkdir -p $GOPATH/log

# mongodb is is an open-source document database, and the leading NoSQL database.
sudo apt-get install mongodb
go get gopkg.in/mgo.v2

# Also you should have C/C++ compiler g++ installed to support cpp source code compilation.
sudo apt-get install build-essential

# flex is required for the code similarity test
sudo apt-get install flex

# Get our code from Github
git clone https://github.com/ZJGSU-Open-Source/GoOnlineJudge.git $GOPATH/src/GoOnlineJudge
git clone https://github.com/ZJGSU-Open-Source/RunServer.git $GOPATH/src/RunServer
git clone https://github.com/sakeven/golog.git $GOPATH/src/golog


# Now, it's time for compilation.
cd $GOPATH/src/GoOnlineJudge/
go build            
cd ../RunServer/
./make.sh

echo
echo ----------
echo installed.
echo ----------
echo

#######
# Run #
#######
cd $GOPATH/src/GoOnlineJudge/
./GoOnlineJudge
./RunServer
