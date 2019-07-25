FROM golang:1.12-stretch AS builder
ENV GO111MODULE on
ENV CGO_ENABLED 1
ENV GOFLAGS -mod=vendor
COPY . /go/src
WORKDIR /go/src
RUN go build -v -o pm_linux64 -ldflags "-s -w" .

# ===========================

FROM debian:stretch
LABEL maintainer="ohko <ohko@qq.com>"
COPY --from=builder /go/src/pm_linux64 /
COPY public/ /public/
COPY view/ /view/
WORKDIR /
ENV TZ Asia/Shanghai
ENV LOG_LEVEL 1
EXPOSE 8080
ENTRYPOINT [ "/pm_linux64" ]