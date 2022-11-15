FROM golang:1.18

WORKDIR /build
COPY . ./
RUN CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w' -o nextcloud-webcron ./

FROM alpine:3.12
COPY --from=0 /build/nextcloud-webcron /bin/

ENTRYPOINT [ "/bin/nextcloud-webcron" ]