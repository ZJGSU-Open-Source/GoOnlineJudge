FROM golang:1.4.2
MAINTAINER Sakeven "sakeven.jiang@daocloud.io"

ENV OJ_HOME $GOPATH/src

WORKDIR $GOPATH/src/

RUN mkdir -p $OJ_HOME/log

ADD . $GOPATH/src/GoOnlineJudge

# Get dependence
RUN git clone https://github.com/sakeven/restweb.git $GOPATH/src/restweb
RUN go get -t GoOnlineJudge

# Build OJ
RUN cd $GOPATH/src/restweb/restweb && go install
RUN cd $GOPATH/src && restweb build GoOnlineJudge

EXPOSE 8080

CMD restweb run GoOnlineJudge
