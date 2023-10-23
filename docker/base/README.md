# CI base image

Run `task` to build all local images. It will build:

- `internal/ci-base:latest`
- `internal/ci-base:5.2`
- `internal/ci-base:5.0`

Run `task diff` to inspect changes between versions.
Run `task info` to print python version for each target.

Run `task tag=latest enter` to spawn a container for a target. The tag
parameter is configurable and may be any valid folder / docker tag.

Run `task -l` to list any other taskfile targets.
