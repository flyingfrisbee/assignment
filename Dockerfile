FROM golang:1.21-alpine AS binary-builder

WORKDIR /app
COPY . .
RUN go build -o /app/start-go .
RUN chmod +x /app/start-go


FROM alpine:3.19

WORKDIR /app
COPY --from=binary-builder /app/start-go .
RUN apk --no-cache add tzdata
ENV TZ=Asia/Jakarta

CMD ["/app/start-go", "prod"]