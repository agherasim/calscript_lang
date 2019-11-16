FROM anapsix/alpine-java:8
LABEL maintainer="Andrei Gherasim <andrei@gherasim.xyz>"

ENV ANTLR_VERSION=4.7.2
ENV ANTLR=antlr-${ANTLR_VERSION}-complete.jar
ENV CLASSPATH .:/antlr-${ANTLR_VERSION}-complete.jar:$CLASSPATH

RUN apk update && \
    apk add make

ADD https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar /
RUN chmod +r /${ANTLR}

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh", "/${ANTLR}" ]