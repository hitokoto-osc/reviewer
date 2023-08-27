FROM alpine:3.18

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR                              /app
ADD resource                             $WORKDIR/
ADD ./temp/linux_amd64/hitokoto_reviewer $WORKDIR/hitokoto_reviewer
RUN chmod +x $WORKDIR/hitokoto_reviewer

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
ENTRYPOINT ["sh"]
