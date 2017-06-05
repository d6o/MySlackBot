FROM fabiorphp/golang-glide:1.8-alpine

ENV APP_DIR $GOPATH/src/github.com/disiqueira/MySlackBot
ENV APP_FUN -help

RUN apk add --no-cache make

RUN apk add --no-cache git

COPY ./Makefile ${APP_DIR}/Makefile
COPY ./glide.lock ${APP_DIR}/glide.lock
COPY ./glide.yaml ${APP_DIR}/glide.yaml

WORKDIR ${APP_DIR}

RUN make deps

COPY . ${APP_DIR}

CMD CompileDaemon -build="make install" -command="msb -cmd=${APP_FUN}"