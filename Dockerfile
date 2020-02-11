FROM golang:rc-buster as build-env

RUN mkdir /workdir
WORKDIR /workdir

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o zcpm-central

FROM scratch
EXPOSE 8000
COPY --from=build-env /workdir/zcpm-receiver /zcpm-receiver

ENTRYPOINT [ "/zcpm-receiver" ]