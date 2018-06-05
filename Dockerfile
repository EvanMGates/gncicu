FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 8080

ADD gnc /bin/gnc
ADD config.yml.dist /etc/gnc/config.yml

CMD ["gnc", "-config", "/etc/gnc/config.yml"]