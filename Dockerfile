FROM node:20 AS template_builder

COPY . /src
WORKDIR /src

RUN npm install -g pnpm
RUN make mail

FROM golang:1.22 as builder

COPY --from=template_builder /src/out /src/out

COPY . /src
WORKDIR /src

RUN make install
RUN make compile_arm

FROM gcr.io/distroless/base-debian12 AS build-release-stage

COPY --from=template_builder /src/out /out

WORKDIR /

COPY --from=builder /src/bin/dartsly-mail-service /dartsly-mail-service

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/dartsly-mail-service"]
