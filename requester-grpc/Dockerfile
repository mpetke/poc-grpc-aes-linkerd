FROM golang:1.12-alpine

WORKDIR /

RUN apk add --no-cache git bash && \
    go get github.com/fullstorydev/grpcurl/... && \
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl

COPY make-request.sh make-request.sh
COPY echo.proto echo-grpc/api/

RUN chmod +x make-request.sh

ENTRYPOINT ["/make-request.sh"]
