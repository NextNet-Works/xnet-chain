FROM golang:1.19-alpine AS buildenv

ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev

# Set up dependencies
RUN apk add --update --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /go/src/github.com/NextNet-Works/

# Add source files
RUN git clone --recursive https://https://github.com/NextNet-Works/xnet-chain.git
WORKDIR /go/src/github.com/NextNet-Works/xnet-chain

RUN make install

# ------------------------------------------------------------------ #

FROM alpine:edge

ENV XNETD_HOME=/.xnet

# Install ca-certificates
RUN apk add --no-cache --update ca-certificates py3-setuptools supervisor wget lz4 gzip

# Temp directory for copying binaries
RUN mkdir -p /tmp/bin
WORKDIR /tmp/bin

COPY --from=buildenv /go/bin/xnetd /tmp/bin
RUN install -m 0755 -o root -g root -t /usr/local/bin xnetd

# Remove temp files
RUN rm -r /tmp/bin

# Add supervisor configuration files
RUN mkdir -p /etc/supervisor/conf.d/
COPY /supervisor/supervisord.conf /etc/supervisor/supervisord.conf 
COPY /supervisor/conf.d/* /etc/supervisor/conf.d/


WORKDIR $XNETD_HOME

# Expose ports
EXPOSE 26656 26657 26658
EXPOSE 1317

# Add entrypoint script
COPY ./scripts/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod u+x /usr/local/bin/entrypoint.sh
ENTRYPOINT [ "/usr/local/bin/entrypoint.sh" ]

STOPSIGNAL SIGINT