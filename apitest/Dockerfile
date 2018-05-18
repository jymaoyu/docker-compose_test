FROM golang
ADD . /go/src/apitest

WORKDIR /go/src/apitest

RUN go get -v github.com/go-xorm/xorm \
&& go get -v github.com/go-sql-driver/mysql \
&& go get -v github.com/astaxie/beego \
&& go get -v github.com/beego/bee 

EXPOSE 8080

CMD ["bee","run"]
