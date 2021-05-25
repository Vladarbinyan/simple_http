FROM golang:1.15 AS build
WORKDIR /src/simple_http/
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser
WORKDIR /app/
COPY --from=build /src/simple_http/app .
CMD ["./app"]

