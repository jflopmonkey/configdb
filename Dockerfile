FROM iron/go:dev

ENV SRC_DIR=/go/src/github.com/jflopmonkey/configdb
WORKDIR /app
ADD . $SRC_DIR

RUN cd $SRC_DIR; go get; go build -o configdb; cp configdb /app/

ENTRYPOINT ["./configdb"]