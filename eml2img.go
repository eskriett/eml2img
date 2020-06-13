package eml2img

import (
	"context"
	"io"
	"os/exec"
	"strings"

	"github.com/jhillyerd/enmime"
)

const exeName = `wkhtmltoimage`

// EML2Img is a type which provides methods to allow an email to be converted to
// an image. Use New() to construct an instance of this type.
type EML2Img struct {
	exePath string
}

// New creates a new instance of EML2Img.
func New() (*EML2Img, error) {
	e2i := &EML2Img{}

	exePath, err := exec.LookPath(exeName)
	if err != nil {
		return nil, err
	}

	e2i.exePath = exePath

	return e2i, nil
}

// Convert takes as input an eml file and outputs a png image.
func (e2i *EML2Img) Convert(ctx context.Context, r io.Reader, w io.Writer) error {
	env, err := enmime.ReadEnvelope(r)
	if err != nil {
		return err
	}

	exePath := e2i.exePath

	emailText := env.HTML

	// If we've been unable to extract any HTML, fallback to using text
	if emailText == "" {
		emailText = env.Text
		emailText = strings.Replace(emailText, "\n", "<br>", -1)
	}

	cmd := exec.CommandContext(ctx, exePath, "--format", "png", "-", "-")
	cmd.Stdin = strings.NewReader(emailText)
	cmd.Stdout = w

	return cmd.Run()
}
