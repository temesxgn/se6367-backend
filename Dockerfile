FROM golang:1.13.4
#MAINTAINER Temesxgn Gebrehiwet, temesxgn@gmail.com

WORKDIR /go/src/github.com/temesxgn/se6367-backend
COPY . .

# RUN go test -v ./...
RUN go get github.com/pilu/fresh
RUN go get ./...

CMD [ "fresh" ]
