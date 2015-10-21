#GoOnlineJudge

GoOnlineJudge is an ACM/ICPC online judge platform.

[**Demo**](http://acm.zjgsu.edu.cn)

##Contents
+ [Installation](https://github.com/ZJGSU-Open-Source/GoOnlineJudge#installation)
	+ [Docker](https://github.com/ZJGSU-Open-Source/GoOnlineJudge#docker)
	+ [Prerequisites](https://github.com/ZJGSU-Open-Source/GoOnlineJudge#prerequisites)
	+ [Quick Start](https://github.com/ZJGSU-Open-Source/GoOnlineJudge#quick-start)
	+ [Manual Installation](https://github.com/ZJGSU-Open-Source/GoOnlineJudge#manual-installation)
	+ [Tips](https://github.com/ZJGSU-Open-Source/GoOnlineJudge#tips)
+ [Maintainers](https://github.com/ZJGSU-Open-Source/GoOnlineJudge#maintainers)
+ [Contributions](https://github.com/ZJGSU-Open-Source/GoOnlineJudge#contributions)
+ [License](https://github.com/ZJGSU-Open-Source/GoOnlineJudge#license)

##Installation
###Prerequisites
**Disclaimer**:
GoOnlineJudge works best on GNU/Linux and has been tested on Ubuntu 14.04 and Arch Linux. Windows and Mac OS X are **not** recommended because [**RunServer**](https://github.com/ZJGSU-Open-Source/RunServer) cannot be built on both of them. 

If you are Windows or Mac OS X user, you can try out [docker-oj](https://github.com/ZJGSU-Open-Source/docker-oj), based on docker image and works out of the box.

### Docker

```bash
docker build -t oj .
docker run --link your_mongo:mongodb --name oj  -e DATA_PATH=your_data_path -e JUDGE_HOST=your_judge_host -v your_data_path:your_data_path -d -p 80:8080 oj
```

If you installed mongodb on your host os, use script like this:
```bash
cat > env_file << EOF
MONGODB_USERNAME=test
MONGODB_PASSWORD=password
MONGODB_PORT_27017_TCP_ADDR=192.168.1.1
MONGODB_PORT_27017_TCP_PORT=27017
MONGODB_INSTANCE_NAME=test
DATA_PATH=your_data_path
JUDGE_HOST=your_judge_host
EOF

docker run --env-file env_file --name oj -v your_data_path:your_data_path -d -p 80:8080 oj
```

### Quick Start
GoOnlineJudge is installed by running one of the following commands in your terminal. You can install it via the command-line with either `curl` or `wget`.

####via curl
```bash
curl -L https://raw.githubusercontent.com/ZJGSU-Open-Source/GoOnlineJudge/master/install.sh | sh
```

####via wget
```bash
wget https://raw.githubusercontent.com/ZJGSU-Open-Source/GoOnlineJudge/master/install.sh | sh
```

###Manual Installation
####Dependences
+ Go
  + GoOnlineJudge is mainly written in Go. 
  + Get Go from [golang.org](http://golang.org)

+ MongoDB
  + MongoDB is a cross-platform document-oriented databases.
  + Get MongoDB from [MongoDB.org](https://www.mongodb.org/)

+ mgo.v2
  + mgo.v2 offers a rich MongoDB driver for Go.
  + Get mgo.v2 via
  ```
  go get gopkg.in/mgo.v2
  ```
  + API documentation is available on [godoc](http://godoc.org/gopkg.in/mgo.v2)

+ [flex](http://flex.sourceforge.net/)
  + flex is the lexical analyzer used in [**RunServer**](https://github.com/ZJGSU-Open-Source/RunServer).
  + Get flex using following command if you are running Ubuntu.
  ```bash
  sudo apt-get install flex
  ```

+ [SIM](http://www.dickgrune.com/Programs/similarity_tester/)
  + SIM is a software and text similarity tester. It's used in RunServer.
  + SIM is shipped along with RunServer.

+ [GCC](https://gcc.gnu.org/)
  + The GNU compiler Collection.
  + Get GCC from [GNU](https://gcc.gnu.org) or using following command if you are running Ubuntu
  ```bash
  sudo apt-get install build-essential
  ``` 

+ iconv-go
  + iconv-go provides iconv support for Go.
  + Get iconv-go via
  ```
  go get github.com/djimenez/iconv-go
  ```

####Install

Obtain latest version via `git`, source codes will be in your $GOPATH/src. 
```bash
git clone https://github.com/ZJGSU-Open-Source/GoOnlineJudge.git $GOPATH/src/GoOnlineJudge
git clone https://github.com/ZJGSU-Open-Source/RunServer.git $GOPATH/src/RunServer
git clone https://github.com/ZJGSU-Open-Source/vjudger.git $GOPATH/src/vjudger
git clone https://github.com/sakeven/restweb.git $GOPATH/src/restweb
```

```bash
# Set $OJ_HOME variable
export OJ_HOME="yourself oj home"

export PATH=$PATH:$GOPATH/bin

#directory for MongoDB Data
mkdir $OJ_HOME/Data

#directory for problem set
mkdir $OJ_HOME/ProblemData

#directory for running user's code
mkdir $OJ_HOME/run

#directory for log
mkdir $OJ_HOME/log
```

Make sure you have these directories in your $GOPATH/src:

```
github.com/  
GoOnlineJudge/  
RunServer/  
gopkg.in/  
restweb/  
```

And these directories in your $OJ_HOME:

```
ProblemData/  
run/  
log/  
```

Now, it's time for compilation.
```bash
cd $GOPATH/src/restweb
go install ./...
cd $GOPATH/src
restweb build GoOnlineJudge
cd $GOPATH/src/RunServer
./make.sh
```

Start OJ
```bash
RunServer&
cd $GOPATH/src
restweb run GoOnlineJudge &
```
Now,you can visit OJ on [http://127.0.0.1:8080](http://127.0.0.1:8080).

####Tips

+ You should always run MongoDB first then followed by OJ.

+ Running web server at 80 port requires administrator privileges. For security reasons, do **not** run our OJ at 80 port.

+ If you want to visit OJ at 80 port, [nginx](http://nginx.org), the HTTP and reverse proxy server is recommended.

##Maintainers
+ memelee

+ sakeven

+ clarkzjw

+ rex-zed

##Contributions
+ We are open for all kinds of pull requests!

+ Just please follow the [Golang style guide](./docs/Golang_Style_Guide.md).

##License
See [LICENSE](LICENSE)
