
FROM golang:alpine AS builder
ARG GIT_PERSONAL_ACCESS_TOKEN
WORKDIR /go/src/
ADD . .
RUN apk add git
RUN go env -w GOPROXY=https://goproxy.cn,direct && go env -w GOPRIVATE=github.com/wxbsocial 
RUN git config --global url."https://wxbsocial:${GIT_PERSONAL_ACCESS_TOKEN}@github.com/wxbsocial".insteadOf "https://github.com/wxbsocial"
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app . 


FROM alpine
WORKDIR /root/
COPY --from=builder /go/src/app .
ENTRYPOINT [ "./app" ]

EXPOSE 8080

ENV GRPC_ADDRESS=:8080
ENV DATA_CENTER_ID=1
ENV NODE_ID=1
ENV BITS_DATA_CENTER_ID=5
ENV BITS_NODE_ID=5
ENV BITS_SEQUENCE=12
ENV BASE_TIMESTAMP=1622476800000
ENV BUFFER_SIZE=100
