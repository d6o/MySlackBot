FROM fabiorphp/golang-glide:1.8-alpine

ENV APP_DIR $GOPATH/src/github.com/disiqueira/MySlackBot
ENV APP_FUN -help

RUN apk add --no-cache make

RUN apk add --no-cache git

COPY . ${APP_DIR}
WORKDIR ${APP_DIR}

RUN make deps

CMD CompileDaemon -build="make install" -command="msb -cmd=${APP_FUN}"