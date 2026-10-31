package pgbadger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

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

	tmpDir, err := os.MkdirTemp("", "pgbadger-")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	outPath := filepath.Join(tmpDir, "out.html")
	args := append([]string{"--quiet", "-x", "html", "-o", outPath}, p.opts...)
	args = append(args, bodyPath)

	if out, err := exec.Command("pgbadger", args...).CombinedOutput(); err != nil {
		return nil, fmt.Errorf("external process aborted: %w: %s", err, out)
	}

	res, err := os.ReadFile(outPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read report: %w", err)
	}
	return io.NopCloser(bytes.NewReader(res)), nil
}
