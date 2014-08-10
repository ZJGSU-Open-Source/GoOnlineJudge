##Installation

We integrate [docker](http://www.docker.com) in our system, so a 64bit version of Ubuntu is needed.[Here](./Docker.md) is a quick start of Docker on Ubuntu.

###Independences

```bash
sudo apt-get update
sudo apt-get install build-essential zsh vim git
sudo wget --no-check-certificate http://install.ohmyz.sh -O - | sh
chsh -s /bin/zsh
```
+ Note that zsh is optional, but I prefer zsh instead of bash...

###Install Go
You can download Go for Linux on [Golang.org](http://golang.org/dl/). For Chinese users which you can not download the official verison, you can also download the version we currently use at [here](http://pan.baidu.com/s/1jGfyO2y)

```bash
sudo mkdir /usr/local/go
mkdir $HOME/go
sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
#replace $VERSION, $OS, $ARCH depending on the version you downloaded.
```

**Optional**: If you are using zsh(I assume you are...), cd to the home directory, and run `vim .zshrc` to edit .zshrc file and add the following lines:

```bash
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
```
Start a new terminal, test if golang is correctly installed

```Go
//hello.go

package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

And run
```bash
go build hello.go
./hello
```
If everything is Ok, you will see "hello, world" on the terminal.

###Install Mongodb

Just run
```bash
sudo apt-get install mongodb
cd ~
mkdir Data #Directory for Mongodb database
vim Mongod.conf
```
Add "mongod -port 8090 --dbpath ~/Data/" to Mongod.conf
And then, `chmod 777 Mongod.conf`. 

After that, you can run Mongodb just using

```bash
./Mongod.conf
```
------

After everything is Ok, clone our repos.

```bash
cd ~/Go
mkdir src bin pkg
cd src
git clone https://github.com/ZJGSU-Open-Source/GoOnlineJudge.git
git clone https://github.com/ZJGSU-Open-Source/RunServer.git
git clone https://github.com/ZJGSU-Open-Source/GoServer.git
```

> Note: You need to download [`labix.org`](http://pan.baidu.com/s/1dDf9dID) to support communication between Go and Mongodb.
Just download it and extract to the same directory of GoOnlineJudge and GoServer.

```bash
mkdir ProblemData #Directory for Problem sets 
mkdir run         #Directory for judge
```

Now, make sure you have these folders in the same directory
> + GoOnlineJudge
> + GoServer
> + RunServer
> + labix.org
> + run
> + ProblemData

###Start

```bash
cd $GOPATH/src/GoServer
go build
./GoServer

cd $GOPATH/src/GoOnlineJudge
go build 
./GoOnlineJudge

cd $GOPATH/src/RunServer
go build
mv RunServer ../GoOnlineJudge
```

start a web broswer, and visit http://127.0.0.1:8080