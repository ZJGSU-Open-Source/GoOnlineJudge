#GoOnlineJudge
##Desciption
GoOnlineJudge is a web server for ZJGSU.  
It requires Ubuntu 14.04, go1.3 and mongodb v2.4.9.
##Install
In order to install our online judge, you should install go1.3, mongodb, and mgo.
###Dependences
+ go1.3 is a complier for go.You can get it from [golang.org](http://golang.org)

+ [mongodb](http://www.mongodb.org/) is is an open-source document database, and the leading NoSQL database.  
```bash
sudo apt-get install mongodb
```

+ [mgo](http://gopkg.in/mgo.v2) is a go driver for mongodb to support communication between go and mongodb. 
```bash
go get gopkg.in/mgo.v2
```
You can get an api documentation from [godoc](http://godoc.org/gopkg.in/mgo.v2)

+ Also you should have C/C++ compiler g++ installed to support cpp source code compilation.
```bash
sudo apt-get install build-essential
```

+ flex is required for the code similarity test
```bash
sudo apt-get install flex
```

###Compile
Then you can download our source code.
```bash
#Get our code from Github
git clone https://github.com/ZJGSU-Open-Source/GoOnlineJudge.git $GOPATH/src/GoOnlineJudge
git clone https://github.com/ZJGSU-Open-Source/RunServer.git $GOPATH/src/RunServer
git clone https://github.com/sakeven/golog.git $GOPATH/src/golog
```
Those source codes file should be in your $GOPATH/src. 
```bash
#directory for problem set
mkdir ProblemData
#directory for running user's code
mkdir run
#directory for log
mkdir log
#configure
cd $GOPATH/src/RunServer
vim Cjudger/config.h
```

Set variable `oj_home` equals to`$GOPATH/src`, make sure use absolute path to replace `$GOPATH`

Make sure you have these directories in your $GOPATH/src:

	GoOnlineJudge/
	RunServer/
	gopkg.in/
	ProblemData/
	run/
	log/
	golog/

Now, it's time for compilation.
```bash
cd $GOPATH/src/
cd GoOnlineJudge/	
go build			
cd ../RunServer/
make
```

##Run
```bash
cd $GOPATH/src/GoOnlineJudge/
./GoOnlineJudge
./RunServer
```
Now,you can visit [http://127.0.0.1:8080](http://127.0.0.1:8080).

#Notice
If you want to visit it at 80 port, we suggest you install [nginx](http://nginx.org/) as a reverse proxy and run nginx at 80 port. 

Because that running web server at 80 port requires administrator privileges and in order to protect your OS, don't run our oj at 80 port.