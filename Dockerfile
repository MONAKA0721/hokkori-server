FROM golang:1.18.0-buster AS go_build

WORKDIR /hokkori

COPY go.mod go.sum /hokkori/
RUN go mod download

COPY . /hokkori
RUN make

FROM debian:buster-slim

WORKDIR /hokkori
RUN apt update && apt install -y ca-certificates
COPY ./entry-point.sh ./entry-point.sh
COPY --from=go_build /hokkori/bin/ ./bin/

RUN chmod 755 ./entry-point.sh
ENTRYPOINT [ "./entry-point.sh" ]
