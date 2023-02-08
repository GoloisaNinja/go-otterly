package config

import (
	"github.com/GoloisaNinja/go-otterly/pkg/models"
	"html/template"
	"log"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	IsProduction  bool
}
type GameConfig struct {
	Games     []models.Game
	Game      models.Game
	GameState models.GameState
}
