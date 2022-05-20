#builder
FROM golang:1.17.7 as builder

#tempat nyimpan image, tidak masuk di HDD tapi di dockernya
WORKDIR /app

COPY . .

RUN go mod download

#outputin file project go ke executable file
RUN go build -o api main.go

#runner
FROM debian:11-slim

COPY --from=builder /app/api /app/

WORKDIR /app

#port yang diakses di luar
EXPOSE 9500

#running command, gak bisa pake spasi, harus dipisah kek gini
CMD ["./api"]