FROM registry.access.redhat.com/ubi7/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/hcp-apiserver \
    USER_UID=1001 \
    USER_NAME=hcp-apiserver

# install operator binary
COPY _output/bin/hcp-apiserver ${OPERATOR}

COPY bin /usr/local/bin
#COPY cert /tmp/cert
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
