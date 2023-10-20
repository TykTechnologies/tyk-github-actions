FROM golang:1.20 AS tools-build

ENV CGO_ENABLED=0
ENV GOBIN=/usr/local/bin
WORKDIR /usr/local/bin

RUN --mount=type=cache,mode=0755,target=/go/pkg/mod go install github.com/TykTechnologies/exp/cmd/schema-gen@main
RUN --mount=type=cache,mode=0755,target=/go/pkg/mod go install github.com/TykTechnologies/exp/cmd/go-fsck@main
RUN --mount=type=cache,mode=0755,target=/go/pkg/mod go install github.com/golang/mock/mockgen@v1.6.0
RUN --mount=type=cache,mode=0755,target=/go/pkg/mod go install golang.org/x/tools/cmd/goimports@latest
RUN --mount=type=cache,mode=0755,target=/go/pkg/mod go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.0
RUN --mount=type=cache,mode=0755,target=/go/pkg/mod go install github.com/fatih/faillint@latest

ARG TASK_VERSION=v3.27.1
RUN curl -sSL https://github.com/go-task/task/releases/download/${TASK_VERSION}/task_linux_amd64.tar.gz | tar -zxv

# Final tools image
#
# This uses the `scratch` as the image base. The image
# doesn't need a base OS, as it's intended to be used
# in CI pipelines that do. Example action:
#
# ```
# - uses: shrink/actions-docker-extract@v3
#   with:
#     image: <tools>
#     path: /usr/local/bin/.
#     destination: /usr/local/bin
# ```
#
# https://github.com/marketplace/actions/docker-extract

FROM scratch
COPY --from=tools-build /usr/local/bin/ /usr/local/bin/
