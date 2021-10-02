FROM ghcr.io/kwitsch/docker-buildimage:main AS build-env

ADD src .
WORKDIR /builddir/entrypoint
RUN gobuild.sh -o entrypoint
WORKDIR /builddir/healthcheck
RUN gobuild.sh -o healthcheck

FROM scratch
COPY --from=build-env /builddir/entrypoint/entrypoint /entrypoint
COPY --from=build-env /builddir/healthcheck/healthcheck /healthcheck

ENTRYPOINT ["/entrypoint"]

HEALTHCHECK --interval=30s --timeout=30s --start-period=30s --retries=3 CMD [ "/healthcheck" ]