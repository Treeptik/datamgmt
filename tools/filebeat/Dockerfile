FROM alpine:3.5
MAINTAINER s.musso@treeptik.fr

ENV FILEBEAT_VERSION 5.2

## Get Filebeat binaries
RUN mkdir -p /opt/datamgmt \
    && apk update && apk add glib wget ca-certificates \
    && cd /opt/datamgmt \
    && wget https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-$FILEBEAT_VERSION-linux-x86_64.tar.gz \
    && tar zxf filebeat-$FILEBEAT_VERSION-linux-x86_64.tar.gz \
    && mv filebeat-$FILEBEAT_VERSION-linux-x86_64 filebeat \
    && rm filebeat-$FILEBEAT_VERSION-linux-x86_64.tar.gz \
    && rm -rf /var/cache/apk/*

#RUN adduser filebeat -D \
#    && chown -R filebeat:filebeat /opt/datamgmt

COPY filebeat-conf /opt/datamgmt/filebeat/conf.d

VOLUME /opt/datamgmt/filebeat

#USER filebeat

COPY docker-entrypoint.sh /

ENTRYPOINT ["/docker-entrypoint.sh"]

LABEL origin datamgmt
LABEL application-type filebeat
LABEL application-version $FILEBEAT_VERSION
