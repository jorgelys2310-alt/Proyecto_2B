FROM golang:1.24 AS build
WORKDIR /app
ENV GOTOOLCHAIN=auto

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /app/bin/api .

FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=build /app/bin/api /app/api

EXPOSE 8080
ENTRYPOINT ["/app/api"]
