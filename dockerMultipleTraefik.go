package dockerMultipleTraefik

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/traefik/genconf/dynamic"
	"github.com/traefik/genconf/dynamic/tls"
	"log"
	"strings"
)

type Config struct {
	LabelPrefix string `json:"labelPrefix,omitempty"`
}

func CreateConfig() *Config {
	return &Config{
		LabelPrefix: "", // 5 * time.Second
	}
}

type Provider struct {
	name        string
	labelPrefix string

	cancel func()
}

func New(ctx context.Context, config *Config, name string) (*Provider, error) {

	return &Provider{
		name:        name,
		labelPrefix: config.LabelPrefix,
	}, nil
}

func (p *Provider) Init() error {
	if strings.TrimSpace(p.labelPrefix) == "" {
		return fmt.Errorf("Label Prefix cannot be null or empty")
	}

	return nil
}

func (p *Provider) Provide(cfgChan chan<- json.Marshaler) error {
	ctx, cancel := context.WithCancel(context.Background())
	p.cancel = cancel

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Print(err)
			}
		}()

		p.loadConfiguration(ctx, cfgChan)
	}()

	return nil
}

func (p *Provider) loadConfiguration(ctx context.Context, cfgChan chan<- json.Marshaler) {

	dynamicConfiguration := generateConfiguration()

	cfgChan <- &dynamic.JSONPayload{Configuration: dynamicConfiguration}

}

func (p *Provider) Stop() error {
	p.cancel()
	return nil
}

func generateConfiguration() *dynamic.Configuration {
	configuration := &dynamic.Configuration{
		HTTP: &dynamic.HTTPConfiguration{
			Routers:           make(map[string]*dynamic.Router),
			Middlewares:       make(map[string]*dynamic.Middleware),
			Services:          make(map[string]*dynamic.Service),
			ServersTransports: make(map[string]*dynamic.ServersTransport),
		},
		TCP: &dynamic.TCPConfiguration{
			Routers:  make(map[string]*dynamic.TCPRouter),
			Services: make(map[string]*dynamic.TCPService),
		},
		TLS: &dynamic.TLSConfiguration{
			Stores:  make(map[string]tls.Store),
			Options: make(map[string]tls.Options),
		},
		UDP: &dynamic.UDPConfiguration{
			Routers:  make(map[string]*dynamic.UDPRouter),
			Services: make(map[string]*dynamic.UDPService),
		},
	}

	return configuration
}

func boolPtr(v bool) *bool {
	return &v
}
