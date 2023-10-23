# Local tests

After building the base image locally from `docker/base`, you can run the
local tests for the individual release targets locally as well.

To run the test suite for 5-LTS release:

```
test -t Taskfile-LTS.yml
```

To run the test suite for 5.2/master release:

```
test -t Taskfile-5.2.yml
```

Both test suites must pass locally. In case of issues that require
fiddling with the base CI image, care should be taken that any changes
made don't impact the release. Obvious issues may arise from bumping the
base OS, Go version or Python versions installed.
