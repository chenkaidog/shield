FROM golang:1.21

WORKDIR /app
COPY common /app/common
COPY account /app/account

RUN apt-get update && apt-get install -y gcc
RUN apt-get update && apt-get install -y build-essential

ENV CGO_ENABLED=1
RUN go env -w GOPROXY='https://goproxy.io/'
RUN cd account && go mod tidy && sh build.sh

WORKDIR /app/account/output
CMD [ "./bootstrap.sh" ]
