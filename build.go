package builder

import (
	"fmt"

	"github.com/taubyte/go-interfaces/builders"
	ci "github.com/taubyte/go-simple-container"
)

// Build will build the given directory as configured and return a builder output
func (b *builder) Build(ops ...ci.ContainerOption) (builders.Output, error) {
	var (
		out *output
		err error
	)

	out, err = new(b.wd)
	if err != nil {
		return out, fmt.Errorf("creating new output failed with: %w", err)
	}
	defer out.deferHandler()

	environment := b.config.HandleDepreciatedEnvironment()
	clientImage, err := b.buildImage()
	if err != nil {
		return out, fmt.Errorf("initializing image failed with: %w", err)
	}

	if err = b.run(out, clientImage, environment, ops...); err != nil {
		return out, fmt.Errorf("running container failed with: %w", err)
	}

	return out, nil
}
