FROM redis:6-alpine

EXPOSE 26379

ADD sentinel.conf /etc/redis/sentinel.conf

RUN chown redis:redis /etc/redis/sentinel.conf

ENV SENTINEL_QUORUM 2

ENV SENTINEL_DOWN_AFTER 30000

ENV SENTINEL_FAILOVER 180000

ENV REDIS_MASTER redis_master

COPY sential-entrypoint.sh /usr/local/bin/

RUN chmod +x /usr/local/bin/sential-entrypoint.sh

ENTRYPOINT ["sential-entrypoint.sh"]