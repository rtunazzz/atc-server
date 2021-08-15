# ATC Server

Lightweight and versatile ATC server made with Go

## Running

### With Go

Make sure you have [go](https://golang.org/doc/install) installed.

1. Download modules

```sh
go mod download
```

2. Install modules

```sh
go install .
```

3. Run the server

```sh
go run main.go
```

### With docker

Make sure you have [docker](https://docs.docker.com/get-docker/) installed.

1. Build the image

```sh
docker build -t atc-server .
```

2. Run the image

```sh
docker run -p 9000:9000 atc-server
```

3. Your server will now be available at `http://localhost:9000/`

## Usage

The server has 4 main routes:

- `/atc` - [docs](./docs/atc.md)
- `/wishlist` - [docs](./docs/wishlist.md)
- `/set_locale` - [docs](./docs/set_locale.md)
- `/redirect` - [docs](./docs/redirect.md)
