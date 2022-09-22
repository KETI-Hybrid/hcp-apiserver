# FROM docker:17.12.0-ce as static-docker-source
FROM debian:stretch

ENV OPERATOR=/usr/local/bin/hcp-apiserver \
    USER_UID=root \
    USER_NAME=hcp-apiserver
# ARG CLOUD_SDK_VERSION=232.0.0
# ENV CLOUD_SDK_VERSION=$CLOUD_SDK_VERSION

COPY _output/bin/hcp-apiserver ${OPERATOR}

COPY bin /usr/local/bin

RUN apt update -y && apt upgrade -y

RUN apt install curl -y

RUN apt-get -y install openssl

RUN apt-get -y install unzip

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"

RUN unzip awscliv2.zip

RUN ./aws/install

# RUN apt update -y && apt upgrade -y

# RUN apt install curl -y

# RUN curl -sL https://aka.ms/InstallAzureCLIDeb | bash

# RUN curl https://sdk.cloud.google.com > install.sh

# RUN bash install.sh --disable-prompts --install-dir=/root

# RUN cp -r /root/google-cloud-sdk/lib/* /usr/local/lib

# RUN cp -r /root/google-cloud-sdk/bin/* /usr/local/bin

RUN  /usr/local/bin/user_setup

RUN chmod +x /usr/local/bin/entrypoint

VOLUME ["/root/.config"]

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
