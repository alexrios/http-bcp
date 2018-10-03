FROM jheise/ubuntu-golang as builder

ENV BIN_NAME http-bcp
ENV REPO_PATH /go/src/github.com/alexrios/http-bcp

RUN mkdir -p $REPO_PATH
COPY ./  $REPO_PATH
WORKDIR $REPO_PATH

## Building
#[Disabling CGO] CGO_ENABLED=0
#[Force Rebuild] -a
#[tagged netgo to make sure we use built-in net package and not the system’s one.] -tags netgo
#[Disabling debug for smaller binary] ldflags -w
#[This will make the linked C part also static into the binary, reducing compatibility risk.] -extldflags "-static"
RUN cp -Rp vendor/* /go/src/ && \
go get -d -v ./... && \
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o $BIN_NAME *.go


FROM mcr.microsoft.com/mssql-tools

ENV BIN_NAME http-bcp
ENV REPO_PATH /go/src/github.com/alexrios/http-bcp

RUN mkdir -p /app
WORKDIR /app

COPY --from=builder /go/src/github.com/alexrios/http-bcp .

EXPOSE 8080

CMD ["./http-bcp"]

