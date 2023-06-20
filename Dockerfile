FROM golang:bullseye as builder

ENV GOPROXY=https://goproxy.cn

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 \
	GOOS=linux \
    GOARCH=amd64 \
	go build -ldflags "-s -w  -X 'gitee.com/lyhuilin/version.Version=`git describe --tags --abbrev=0`' -X 'gitee.com/lyhuilin/version.Commit=`git rev-parse --short HEAD`' -X 'gitee.com/lyhuilin/version.BuildDate=`date +%FT%T%z`'   -X 'gitee.com/lyhuilin/version.Author=`git log |head -n 3| grep Author| awk '{print $2;}'`'  -X 'gitee.com/lyhuilin/version.BranchName=`git branch | awk '/\*/ { print $2; }'`'  -X 'gitee.com/lyhuilin/version.CommitId=`git log |head -n 1| awk '{print $2;}'`' " -o url_location ./cmd/url_location/

FROM debian:bullseye as runner
ENV TZ=Asia/Shanghai

WORKDIR /app

COPY --from=builder /app/url_location .
COPY --from=builder /app/cacert.pem /etc/ssl/certs/cacert.pem
VOLUME /app/conf
VOLUME /app/autocert
VOLUME /app/log
VOLUME /app/upload
EXPOSE 8080

ENTRYPOINT ["./url_location" ,"-c","/app/conf/config.yaml"]
