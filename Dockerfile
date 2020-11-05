FROM golang:latest AS build
WORKDIR /go/src/work
ENV GO111MODULE=on

# make cache
COPY go.mod .
COPY go.sum .
RUN go mod download

# COPY the source code as the last step
COPY . .
# build app for next stage
RUN go build -o /go/bin/app

FROM alpine:latest
WORKDIR /app
COPY --from=build /go/bin/app .
# add user
RUN addgroup go && \
    adduser -D -G go go && \
    chown -R go:go /app/app
    CMD [ "./cmd" ]