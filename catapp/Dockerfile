FROM golang:alpine AS builder 
RUN apk update && apk add --no-cache git
WORKDIR /go/src/catapp
COPY . .
ENV GO111MODULE=off
RUN go get -d -v
RUN go build -o /go/bin/server

FROM alpine 
COPY --from=builder /go/bin/server /server
COPY --from=builder /go/src/catapp/index.gohtml /index.gohtml
ENTRYPOINT ["/server"]
