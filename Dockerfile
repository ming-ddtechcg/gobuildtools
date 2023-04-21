FROM golang:1.20.2-alpine

RUN apk add make

RUN mkdir resource_converter; mkdir /usr/local/shared-builds
COPY resource_converter/* resource_converter/
RUN cd resource_converter; make; cp resource_converter /usr/local/shared-builds

