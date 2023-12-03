FROM golang:1.19
WORKDIR /app
COPY . .
RUN go build -o minio-processing ./cmd
EXPOSE 3000
CMD ["./minio-processing"]