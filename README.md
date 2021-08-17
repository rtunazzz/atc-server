<p align="center"><img width=60% src="https://i.imgur.com/ZWAtLwR.png"></p>

# ATC Server made with Go

> Live version available at https://atc.bonzay.io

# Running

### With docker

Make sure you have [docker](https://docs.docker.com/get-docker/) installed.

#### You can either pull the image from [dockerhub](https://hub.docker.com/repository/docker/arturhnat/atc-server/):

```sh
docker run -p 9000:9000 --name atc-server arturhnat/atc-server:latest
```

#### or after cloning this repo, you can:

1. Build the image

```sh
docker build -t atc-server .
```

2. Run the image

```sh
docker run -p 9000:9000 atc-server
```

3. Your server will now be available at `http://localhost:9000/`

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

# Usage

The server has 4 main routes:

- `/atc` - [docs](./docs/atc.md)
- `/wishlist` - [docs](./docs/wishlist.md)
- `/set_locale` - [docs](./docs/set_locale.md)
- `/redirect` - [docs](./docs/redirect.md)

---

# Adding more sites

All supported sites are configured in the `config.json` file and are loaded dynamically, so to add a new site for either ATC/ Wishlist, all you need to do is add the proper URLs to the config file.

If you'd want to add another site to the locale picker, all you need to do is add an entry [into the router file](/internal/router.go#L28). You may also want to configure the [`setLocaleHandler`](./internal/handlers.go#L141) function.

Each site in the config file is expected to have the following structure:

```js
"name_of_the_site": { // this will then be the required query parameter for the /atc and /wishlist routes
    "run_script": false, // whether or not the default form handler should be overwritten with a different script
    "script_name": "site.js", // name of the script to execute
    "atc_supported": true, // whether or not Adding to Cart is supported
    "wishlist_supported": true, // whether or not Adding to Wishlist is supported
    "wait_for_submit": false, // Whether or not should the script wait for the form to finish submitting before redirecting
    "locales": { // each entry in the locales object is understood as a new locale. There always has to be at least one, default one named 'default'.
      "default": {
        "atc_url": "https://site.com/add-to-cart", // URL to which the ATC form gets POSTed
        "atc_redirect_url": "https://site.com/checkout", // URL to which the user will be redirected after ATC, can be anything
        "wishlist_url": "https://site.com/add-to-wishlist", // URL to which the Wishlist form gets POSTed
        "wishlist_redirect_url": "https://site.com/wishlist", // URL to which the user will be redirected after ATC, can be anything
      }
    }
  }
```
