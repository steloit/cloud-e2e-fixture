FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /fixture .

FROM scratch
COPY --from=build /fixture /fixture
USER 1000:1000
ENTRYPOINT ["/fixture"]
