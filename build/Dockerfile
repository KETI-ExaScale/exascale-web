FROM registry.access.redhat.com/ubi7/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/exascale.exascale-web

# install operator binary
COPY _output/bin/exascale.exascale-web ${OPERATOR}

ENTRYPOINT ["/usr/local/bin/exascale.exascale-web"]