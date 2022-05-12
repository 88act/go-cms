FROM golang:1.17 as builder  
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w CGO_ENABLED=0
RUN go env 
RUN go get github.com/cortesi/modd/cmd/modd
#CMD "modd"
CMD ["echo", "hello"]
CMD ["tail", "-f", "/dev/null"]

 
