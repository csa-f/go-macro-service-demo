FROM go-bin
ADD . /app/src
WORKDIR /app/src
RUN mkdir -p /app/bin /app/log /app/base/static /app/data

RUN export GOROOT=/usr/local/go1.23.1 && \
	export GOPROXY=https://goproxy.cn && \
    export PATH=$PATH:$GOROOT/bin && \
    go mod tidy && \
    go build cmd/test-api.go && \
    mv test-api /app/bin && \
    cp -r conf /app && \
    cp start.sh /app && \
    chmod +x /app/start.sh

EXPOSE 8080

CMD ["/app/start.sh"]