FROM golang:1.24.2 AS baseimage

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o main .

#distroless image

FROM gcr.io/distroless/base

COPY --from=baseimage /app/main .

COPY --from=baseimage /app/templates ./templates
COPY --from=baseimage /app/static ./static

EXPOSE 8080

CMD [ "./main" ]

