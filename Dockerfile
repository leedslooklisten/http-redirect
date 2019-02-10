FROM alpine
COPY http-redirect /redirectserver
COPY cert.pem /
COPY privkey.pem /
EXPOSE 80 443
ENTRYPOINT ["/redirectserver"]