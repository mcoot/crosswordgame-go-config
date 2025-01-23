package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/goccy/go-yaml"
)

type BackendConfig struct {
	Image           string `json:"image"`
	Tag             string `json:"tag"`
	ContainerName   string `json:"container_name"`
	Port            int    `json:"port"`
	HealthcheckPath string `json:"healthcheck_path"`
}

type EnvoyTLSConfig struct {
	Enabled    bool   `json:"enabled"`
	CertDomain string `json:"cert_domain,omitempty"`
	CertEmail  string `json:"cert_email,omitempty"`
}

type EnvoyConfig struct {
	Image     string         `json:"image"`
	Tag       string         `json:"tag"`
	AdminPort int            `json:"admin_port"`
	Port      int            `json:"port"`
	TLS       EnvoyTLSConfig `json:"tls"`
}

type Config struct {
	InputPath  string        `json:"input_path"`
	OutputPath string        `json:"output_path"`
	Backend    BackendConfig `json:"backend"`
	Envoy      EnvoyConfig   `json:"envoy"`
}

func main() {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = "config.yaml"
	}

	config, err := readConfig(configFile)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	err = renderTemplates(config.InputPath, config.OutputPath, config)
	if err != nil {
		log.Fatalf("failed to render templates: %v", err)
	}
}

func readConfig(filepath string) (*Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	err = yaml.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func renderTemplates(inDir, outDir string, config *Config) error {
	t := template.Must(template.New("template").ParseGlob(filepath.Join(inDir, "*.template")))

	return filepath.Walk(inDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".template" {
			outFileName := info.Name()[:len(info.Name())-len(".template")]
			outFile, err := os.Create(filepath.Join(outDir, outFileName))
			if err != nil {
				return err
			}
			defer outFile.Close()

			err = t.ExecuteTemplate(outFile, info.Name(), config)
			log.Printf("Rendered %s", outFile.Name())
			if err != nil {
				return err
			}
		}
		return nil
	})
}
