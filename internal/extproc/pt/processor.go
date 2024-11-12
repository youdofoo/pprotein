package pt

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"

	"github.com/kaz/pprotein/internal/collect"
)

type (
	processor struct {
		opts []string
	}
)

func (p *processor) Cacheable() bool {
	return true
}

func (p *processor) Process(snapshot *collect.Snapshot) (io.ReadCloser, error) {
	bodyPath, err := snapshot.BodyPath()
	if err != nil {
		return nil, fmt.Errorf("failed to find snapshot body: %w", err)
	}

	args := append([]string{bodyPath}, p.opts...)
	cmd := exec.Command("pt-query-digest", args...)

	res, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("external process aborted: %w", err)
	}

	return io.NopCloser(bytes.NewBuffer(res)), nil
}
