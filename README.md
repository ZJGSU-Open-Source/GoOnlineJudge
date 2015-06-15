#GoOnlineJudge

GoOnlineJudge is an ACM/ICPC online judge platform.

[**Demo**](http://acm.zjgsu.edu.cn)

##Contents
+ [Installation](https://github.com/ZJGSU-Open-Source/GoOnlineJudge#installation)
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
GoOnlineJudge works best on Linux. Windows and Mac OS X are **not** recommended because [**RunServer**](https://github.com/ZJGSU-Open-Source/RunServer) cannot be built on both of them.

+ Linux-based **x86** operating system (Ubuntu 14.04 is **recommended**)
+ `wget` or `curl` should be installed

### Quick Start
GoOnlineJudge is installed by running one of the following commands in your terminal. You can install this via the command-line with either `curl` or `wget`.

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
  + Get flex via
  ```bash
  sudo apt-get install flex
  ```

+ [SIM](http://www.dickgrune.com/Programs/similarity_tester/)
  + SIM is a software and text similarity tester. It's used in RunServer.
  + SIM is shipped along with RunServer.

+ [GCC](https://gcc.gnu.org/)
  + The GNU compiler Collection.
  + Get GCC from [GNU](https://gcc.gnu.org) or via
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
#directory for MongoDB Data
mkdir $GOPATH/Data

#directory for problem set
mkdir $GOPATH/src/ProblemData

#directory for running user's code
mkdir $GOPATH/src/run

#directory for log
mkdir $GOPATH/src/log
mkdir $GOPATH/src/GoOnlineJudge/log

#configure
cd $GOPATH/src/RunServer
vim Cjudger/config.h
```

Set variable `oj_home` equals to`$GOPATH/src`, make sure use absolute path to replace `$GOPATH`

Make sure you have these directories in your $GOPATH/src:

	github.com/
	GoOnlineJudge/
	RunServer/
	gopkg.in/
	ProblemData/
	run/
	log/
	restweb/

Now, it's time for compilation.
```bash
cd $GOPATH/src/restweb
cd restweb
go install
cd $GOPATH/src/
restweb build GoOnlineJudge/	
cd $GOPATH/src/RunServer/
./make.sh
```

####Run
Start MongoDB
```bash
mongod --dbpath $GOPATH/Data --logpath $GOPATH/Data/mongo.log
```
Start OJ
```bash
cd $GOPATH/src/GoOnlineJudge
./RunServer&
cd ../
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
