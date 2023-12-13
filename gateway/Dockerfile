FROM golang:1.21

WORKDIR /app
COPY common /app/common
COPY gateway /app/gateway

RUN cd gateway && go mod tidy && sh build.sh

WORKDIR /app/account/output
CMD [ "./bootstrap.sh" ]