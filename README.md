# eml2img

[![GoDoc](https://godoc.org/github.com/eskriett/eml2img?status.svg)](https://pkg.go.dev/github.com/eskriett/eml2img)
[![Go Report
Card](https://goreportcard.com/badge/github.com/eskriett/eml2img)](https://goreportcard.com/report/github.com/eskriett/eml2img)

Convert MIME emails to an image.

## Usage

Ensure `wkhtmltoimage` is installed. Then:

```
$ go run cmd/eml2img/main.go mail.eml out.png
```

## Credits

eml2img makes use of several great projects:

* [`jhillyerd/enmime`](https://github.com/jhillyerd/enmime) - for MIME decoding
* [`wkhtmltoimage`](https://wkhtmltopdf.org/) - for converting html to an image
  using the webkit rendering engine
