# Gudtani - Gudang Hasil Tani

Aplikasi ini dibuat untuk memudahkan petani dalam menyewa gudang serta memudahkan pemilik gudang dalam memanajemen barang yang masuk/keluar.

## API Design

Kunjungi link berikut untuk melihat desain API

[Gudtani API Design](https://app.swaggerhub.com/apis/jiranmuhammad7/gudtani/1.0.0)

## Configuration

Untuk melakukan konfigurasi database milik Anda maka perlu membuat _development environment file_ dengan nama `config.env`. Berikut merupakan contoh konfigurasinya:

- `DBNAME = <YOUR_DATABASE_NAME>`
- `DBUSER = <YOUR_DATABASE_USER>`
- `DBPASS = <YOUR_DATABASE_PASSWORD>`
- `DBHOST = <YOUR_DATABASE_HOST>`
- `DBPORT = <YOUR_DATABASE_PORT>`

Jika Anda ingin mengganti nama file `config.env` ke nama _environment file_ yang anda inginkan maka perlu mengubahnya juga pada inisialisasi nama _environment file_. Inisialisasi tersebut dapat dilakukan pada sintaks `godotenv.Load("YOUR_ENV_FILE_NAME")` di dalam file `main.go`.

## Testing

Gunakan _command_ berikut untuk melakukan testing

- _Command_ untuk melakukan _test_ pada _user service_ `go test ./user/service/usecase_test.go -v -coverpkg=./... -coverprofile=./user/cover.out && go tool cover -html=./user/cover.out`
- _Command_ untuk melakukan _test_ pada _product service_`go test ./product/service/usecase_test.go -v -coverpkg=./... -coverprofile=./product/cover.out && go tool cover -html=./product/cover.out`
- _Command_ untuk melakukan _test_ pada _rent service_ `go test ./rent/service/usecase_test.go -v -coverpkg=./... -coverprofile=./rent/cover.out && go tool cover -html=./rent/cover.out`
- _Command_ untuk melakukan _test_ pada _warehouse service_ `go test ./warehouse/service/usecase_test.go -v -coverpkg=./... -coverprofile=./warehouse/cover.out && go tool cover -html=./warehouse/cover.out`

## Build

- Command untuk melakukan _build_ dan _run_ aplikasi menggunakan _docker compose_ `docker-compose up --build`
