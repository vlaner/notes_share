FROM golang:1.22.0-alpine3.18 AS build

WORKDIR /app

RUN apk add build-base
RUN apk --no-cache add tzdata
COPY . ./
RUN go mod download

RUN go build -ldflags "-s -w"  -o ./notes ./cmd/notes/main.go
RUN go build -ldflags "-s -w"  -o ./migrate ./cmd/migrate/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Europe/London

COPY --from=build /app/notes ./
COPY --from=build /app/migrate ./
COPY --from=build /app/internal/adapter/storage/sqlite/migrations ./migrations
# COPY .env ./
# RUN chmod +x exec
# ENTRYPOINT ["./exec"]