# Gudtani - Gudang Hasil Tani

Aplikasi ini dibuat untuk memudahkan petani dalam menyewa gudang serta memudahkan pemilik gudang dalam memanajemen barang yang masuk/keluar.

## API Design

Kunjungi link berikut untuk melihat desain API

[Gudtani API Design](https://app.swaggerhub.com/apis/jiranmuhammad7/gudtani/1.0.0)

## Testing

Gunakan _command_ berikut untuk melakukan testing

- _Command_ untuk melakukan _test_ pada _user service_ `go test ./user/service/usecase_test.go -v -coverpkg=./... -coverprofile=cover-user.out && go tool cover -html=cover-user.out`
- _Command_ untuk melakukan _test_ pada _product service_`go test ./product/service/usecase_test.go -v -coverpkg=./... -coverprofile=cover-product.out && go tool cover -html=cover-product.out`
- _Command_ untuk melakukan _test_ pada _rent service_ `go test ./rent/service/usecase_test.go -v -coverpkg=./... -coverprofile=cover-rent.out && go tool cover -html=cover-rent.out`
- _Command_ untuk melakukan _test_ pada _warehouse service_ `go test ./warehouse/service/usecase_test.go -v -coverpkg=./... -coverprofile=cover-warehouse.out && go tool cover -html=cover-warehouse.out`

## Build

- Command untuk melakukan _build_ dan _run_ aplikasi menggunakan _docker compose_ `docker-compose up --build`
