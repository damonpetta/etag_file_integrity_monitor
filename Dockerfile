FROM alpine:edge

COPY ./ /etag_monitor

RUN apk add -U bash go git make gcc libstdc++ build-base \
    && cd /etag_monitor && make \
    && cd /etag_monitor && GOBIN=/usr/bin make install \
    && apk del bash go git make gcc libstdc++ build-base \
    && apk add -U ca-certificates \
    && rm -rf /var/cache/apk/* /monitor

CMD /usr/local/bin/etag_monitor
