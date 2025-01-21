package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Config représente la configuration globale de l'application
type Config struct {
	// Paramètres généraux
	LogLevel     string        `json:"log_level"`
	RefreshRate  time.Duration `json:"refresh_rate"`
	DataDir      string        `json:"data_dir"`
	MaxProcesses int           `json:"max_processes"`

	// Paramètres UI
	Theme       string   `json:"theme"`
	Columns     []string `json:"columns"`
	DefaultView string   `json:"default_view"`

	// Paramètres monitoring
	EnableMetrics bool             `json:"enable_metrics"`
	MetricsPath   string           `json:"metrics_path"`
	Alerts        map[string]Alert `json:"alerts"`
}

// Alert définit les seuils d'alerte pour différentes métriques
type Alert struct {
	Warning  float64 `json:"warning"`
	Critical float64 `json:"critical"`
	Enabled  bool    `json:"enabled"`
}

// DefaultConfig retourne une configuration par défaut
func DefaultConfig() *Config {
	return &Config{
		LogLevel:      "info",
		RefreshRate:   2 * time.Second,
		DataDir:       "data",
		MaxProcesses:  50,
		Theme:         "dark",
		DefaultView:   "dashboard",
		EnableMetrics: true,
		Columns:       []string{"pid", "name", "cpu", "memory", "status"},
		Alerts: map[string]Alert{
			"cpu": {
				Warning:  80.0,
				Critical: 95.0,
				Enabled:  true,
			},
			"memory": {
				Warning:  85.0,
				Critical: 95.0,
				Enabled:  true,
			},
		},
	}
}

// Load charge la configuration depuis un fichier
func Load() (*Config, error) {
	cfg := DefaultConfig()

	configPath := os.Getenv("ANTON_CONFIG")
	if configPath == "" {
		configPath = filepath.Join("config", "config.json")
	}

	if _, err := os.Stat(configPath); err == nil {
		file, err := os.Open(configPath)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		if err := decoder.Decode(cfg); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

// Save sauvegarde la configuration dans un fichier
func (c *Config) Save() error {
	configPath := os.Getenv("ANTON_CONFIG")
	if configPath == "" {
		configPath = filepath.Join("config", "config.json")
	}

	// Créer le répertoire si nécessaire
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return err
	}

	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(c)
}

// Validate vérifie la validité de la configuration
func (c *Config) Validate() error {
	// TODO: Ajouter des validations
	return nil
}
