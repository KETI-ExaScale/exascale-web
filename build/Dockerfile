FROM registry.access.redhat.com/ubi7/ubi-minimal:latest

ENV OPERATOR=/exascale-web/exascale.exascale-web

# install operator binary
COPY _output/bin/metric-collector ${OPERATOR}

ENTRYPOINT ["/exascale-web/exascale.exascale-web"]