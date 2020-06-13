# eml2img

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