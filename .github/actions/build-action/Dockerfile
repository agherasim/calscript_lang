FROM anapsix/alpine-java:8
LABEL maintainer="Andrei Gherasim <andrei@gherasim.xyz>"

ENV ANTLR_VERSION=4.7.2
ENV CLASSPAHT .:/antlr-${ANTLR_VERSION}-complete.jar

RUN apk update && \
    apk add make

ADD https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar /
RUN chmod +r /antlr-${ANTLR_VERSION}-complete.jar

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]