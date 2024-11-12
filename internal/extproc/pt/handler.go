package pt

import (
	_ "embed"
	"fmt"

	"github.com/kaz/pprotein/internal/collect"
	"github.com/kaz/pprotein/internal/extproc"
	"github.com/kaz/pprotein/internal/persistent"
	"github.com/kaz/pprotein/internal/storage"
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
)

type (
	handler struct {
		opts   *collect.Options
		config *persistent.Handler
	}
)

//go:embed pt.yml
var defaultConfig []byte

func NewHandler(opts *collect.Options, store storage.Storage) (*handler, error) {
	h := &handler{
		opts: opts,
	}

	config, err := persistent.New(store, "pt.yml", defaultConfig, h.sanitize)
	if err != nil {
		return nil, fmt.Errorf("failed to create targets: %w", err)
	}
	h.config = config

	return h, nil
}

func (h *handler) Register(g *echo.Group) error {
	h.config.RegisterHandlers(g.Group("/config"))

	content, err := h.config.GetContent()
	if err != nil {
		return fmt.Errorf("failed to get config content: %w", err)
	}
	var conf struct {
		Options []string `yaml:"options"`
	}
	if err := yaml.Unmarshal(content, &conf); err != nil {
		return fmt.Errorf("failed to unmarshal: %w", err)
	}

	if err := extproc.NewHandler(&processor{opts: conf.Options}, h.opts).Register(g); err != nil {
		return fmt.Errorf("failed to register extproc handlers: %w", err)
	}
	return nil
}

func (h *handler) sanitize(raw []byte) ([]byte, error) {
	var config interface{}
	if err := yaml.Unmarshal(raw, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	res, err := yaml.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}
	return res, nil
}
