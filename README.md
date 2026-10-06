# cloud-e2e-fixture

The repository the end-to-end suite of `steloit/cloud` builds as a Web Service's repository Source (steloit/cloud#77). Its Dockerfile produces a server that listens on `PORT` (8080 unset), answers `GET /healthz` with `ok` and `GET /` with `steloit e2e fixture`.

The suite reads it through the platform's dev GitHub App, which is installed on this repository alone. Keep it small: every e2e run builds it.
