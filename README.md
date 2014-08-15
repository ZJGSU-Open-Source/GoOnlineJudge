#Desciption
GoOnlineJudge is a web server for ZJGSU.  
It built in Ubuntu 14.04, go1.3 and mongodb v2.4.9.
#Install
In order to intall our online judge, you should install go1.3, mongodb, and mgo.

1. go1.3 is a complier for go.You can get it from [golang.org](http://golang.org)
2. [mongodb](http://www.mongodb.org/) is is an open-source document database, and the leading NoSQL database.  

		sudo apt-get install mongodb
3. [mgo](http://gopkg.in/mgo.v2) is a go driver for mongodb to support communication between go and mongodb. 

		go get gopkg.in/mgo.v2
You can get an api documentation from [godoc](http://godoc.org/gopkg.in/mgo.v2)
4. Also you should hava installed g++ compiler to support cpp source code compilation.

		sudo apt-get install g++

Then you can download our source code.

	go get github.com/ZJGSU-Open-Source/GoOnlineJudge
	go get github.com/ZJGSU-Open-Source/GoServer
	go get github.com/ZJGSU-Open-Source/RunServer

Those source code file will be in your $GOPATH/src. Also you should create two directories in your $GOPATH/src.

	mkdir ProblemData	//directory for problem set
	mkdir run	//directory for run user's code

Make sure you hava these directories in your $GOPATH/src:

	GoOnlineJudge/
	GoServer/
	RunServer/
	gopkg.in/
	ProblemData/
	run/

Now, it's time for compilation.

	cd $GOPATH/src/
	cd GoOnlineJudge/	//executabele file GoOnlineJudge
	go build
	cd ../GoServer/
	go build			//executabele file GoServer
	cd ../RunServer/
	go build			//executabele file RunServer
	cd CJudger/
	g++ runner.cc -o runner 	//executabele file runner
	g++ compiler.cc -o compiler //executabele file compiler

Then you should move 3 executable files runner, compiler and RunServer to GoOnlineJudge/.

#Run

	cd $GOPATH/src/GoOnlineJudge/
	./GoOnlineJudge
	cd ../GoServer/
	./GoServer

Now,you can visit [http://127.0.0.1:8080](http://127.0.0.1:8080).

#notice
If you want to visit it at 80 port, we suggest you install nginx as a reverse proxy and run nginx at 80 port. Because that run web server at 80 port requires administrator privileges and in order to protect your OS, don't run our oj at 80 port.

