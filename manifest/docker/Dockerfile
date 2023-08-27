FROM alpine:3.18

###############################################################################
#                                INSTALLATION
###############################################################################

ENV TZ Asia/Shanghai
RUN apk add alpine-conf tzdata && \
    /sbin/setup-timezone -z Asia/Shanghai && \
    apk del alpine-conf

ENV WORKDIR                              /app
ADD resource                             $WORKDIR/
ADD ./temp/linux_amd64/hitokoto_reviewer $WORKDIR/hitokoto_reviewer
RUN chmod +x $WORKDIR/hitokoto_reviewer

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD     ["./hitokoto_reviewer"]
