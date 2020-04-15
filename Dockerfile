FROM golang:1.12-alpine as builder

RUN apk add --no-cache --virtual .build-deps \
bash \
git \
musl-dev

RUN mkdir build
COPY . /build
WORKDIR /build

RUN go mod tidy
RUN go build -o vscale-task .
RUN adduser -S -D -H -h /build vscale-task
USER vscale-task

FROM scratch
COPY --from=builder /build/vscale-task /app/
WORKDIR /app
EXPOSE 5000
CMD ["./vscale-task --port 8181 --token 10824-149857-29356208-zpl"]