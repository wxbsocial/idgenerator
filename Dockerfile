FROM golang:alpine AS builder
WORKDIR /go/src/
ADD . .
RUN apk add git
RUN go env -w GOPROXY=https://goproxy.cn,direct && go env -w GOPRIVATE=github.com/wxbsocial 
RUN git config --global url."https://wxbsocial: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@github.com/wxbsocial".insteadOf "https://github.com/wxbsocial"
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app . 


FROM alpine
WORKDIR /root/
COPY --from=builder /go/src/app .
ENTRYPOINT [ "./app" ]

EXPOSE 8080
