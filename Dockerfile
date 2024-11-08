FROM golang:1.22.7 as builder
WORKDIR /src
COPY ./go.mod ./
COPY ./go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /main ./cmd/gateway/main.go

FROM alpine:3.15 as runner
COPY --from=builder /main /main
COPY ./config /config
CMD ["/main"]