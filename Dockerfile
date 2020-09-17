FROM golang:1.14.9-alpine3.12
RUN apk update
RUN apk add git

ENV DEP_VERSION="0.5.4"
ENV PROJECT_NAME="DeviceModelService"

RUN apk update; \
    apk add --no-cache \
        ca-certificates \
        curl \
        git \
        make \
        openssl; \
    curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o /bin/dep; \
    chmod +x /bin/dep; \
    rm -rf /var/cache/apk/*; \
    rm -rf /tmp/*;

# PORT
ENV PORT 8080
EXPOSE 8080
# Copy all from current directory to the destination in container
RUN mkdir /app
RUN mkdir /app/src
RUN mkdir /app/src/${PROJECT_NAME}
# setup gopath
ENV GOPATH /app

COPY . /app/src/${PROJECT_NAME}
RUN chmod +x /app/src/${PROJECT_NAME}
#Change working dir to copied source
WORKDIR /app/src/${PROJECT_NAME}
RUN ls .
# run go dependency manager dep
RUN dep ensure -v

RUN pwd
RUN ls .
# build the grpc server main and output to
RUN env GOOS=linux go build -o /bin/main DeviceModelService/devicemodelservice/grpc.server

CMD ["/bin/main"]

MAINTAINER Manoj Ramakrishnan