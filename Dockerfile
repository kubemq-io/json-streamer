FROM kubemq/gobuilder as builder
ARG VERSION
ARG GIT_COMMIT
ARG BUILD_TIME
ENV GOPATH=/go
ENV PATH=$GOPATH:$PATH
ENV ADDR=0.0.0.0
ADD . $GOPATH/github.com/kubemq-io/player
WORKDIR $GOPATH/github.com/kubemq-io/player
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -mod=vendor -installsuffix cgo -ldflags="-w -s -X main.version=$VERSION" -o player-run .
FROM registry.access.redhat.com/ubi8/ubi-minimal
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH
RUN mkdir /player
COPY --from=builder $GOPATH/github.com/kubemq-io/player/player-run ./player
WORKDIR player
ENTRYPOINT ["./player-run"]

