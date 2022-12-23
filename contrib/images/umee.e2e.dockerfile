# Docker for e2e testing
# Creates dynamic binaries, by building from the latest version of:
# umeed, price-feeder,
# And release version of peggo

FROM golang:1.19-bullseye AS builder

## Build umeed
WORKDIR /src/umee
COPY . .
RUN go mod download
RUN make install && \
    cd price-feeder && make install

## Download Peggo
WORKDIR /src
RUN wget https://github.com/umee-network/peggo/releases/download/v1.4.0/peggo-v1.4.0-linux-amd64.tar.gz && \
    tar -xvf peggo-v*


## Prepare the final clear binary
FROM ubuntu:rolling
COPY --from=builder /go/pkg/mod/github.com/\!cosm\!wasm/wasmvm\@v*/internal/api/libwasmvm.*.so /usr/lib/
COPY --from=builder /go/bin/* /usr/local/bin/
COPY --from=builder /src/peggo-v*/peggo /usr/local/bin/

#EXPOSE 26656 26657 1317 9090 7171
#ENTRYPOINT ["umeed", "start"]
#
