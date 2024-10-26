FROM golang:1.22.2 as builder
WORKDIR /src
COPY ./go.mod ./
COPY ./go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /main ./cmd/gateway/main.go

FROM scratch as runner
COPY --from=builder /main /main
COPY ./config /config
CMD ["/main"]