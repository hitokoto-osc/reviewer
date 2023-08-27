FROM loads/alpine:3.8

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
CMD ./hitokoto_reviewer
