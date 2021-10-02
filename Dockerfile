FROM ghcr.io/kwitsch/docker-buildimage:main AS build-entrypoint

ADD src/entrypoint .
RUN gobuild.sh -o entrypoint

FROM ghcr.io/kwitsch/docker-buildimage:main AS build-healthcheck

ADD src/healthcheck .
RUN gobuild.sh -o healthcheck

FROM scratch
COPY --from=build-entrypoint /builddir/entrypoint /entrypoint
COPY --from=build-healthcheck /builddir/healthcheck /healthcheck

ENTRYPOINT ["/entrypoint"]

HEALTHCHECK --interval=30s --timeout=30s --start-period=30s --retries=3 CMD [ "/healthcheck" ]