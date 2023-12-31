FROM golang:1.21

WORKDIR /app
COPY common /app/common
COPY gateway /app/gateway

RUN go env -w GOPROXY='https://goproxy.io/'
RUN cd gateway && go mod tidy && sh build.sh

WORKDIR /app/gateway/output
CMD [ "./bootstrap.sh" ]