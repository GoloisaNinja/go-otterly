package handlers

import (
	"encoding/json"
	"github.com/GoloisaNinja/go-otterly/pkg/config"
	"github.com/GoloisaNinja/go-otterly/pkg/helpers"
	"github.com/GoloisaNinja/go-otterly/pkg/models"
	"github.com/GoloisaNinja/go-otterly/pkg/render"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	GC  *config.GameConfig
}

func NewRepository(a *config.AppConfig, gc *config.GameConfig) *Repository {
	return &Repository{
		App: a,
		GC:  gc,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Template Routes

func (e *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (e *Repository) Game(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gameId := params["id"]
	// load this game into gameconfig
	goodGameId := helpers.LoadGame(e.GC, gameId)
	if !goodGameId {
		render.RenderTemplate(w, "notfound.page.tmpl", &models.TemplateData{})
		return
	}
	type InitialGameData struct {
		Title       string
		Description string
		Node        models.GameNode
		TotalNodes  int
	}
	gnd := InitialGameData{
		Title:       e.GC.Game.Title,
		Description: e.GC.Game.Description,
		Node:        e.GC.Game.Nodes[0],
		TotalNodes:  len(e.GC.Game.Nodes),
	}
	render.RenderTemplate(w, "game.page.tmpl", &models.TemplateData{
		Data: map[string]interface{}{"game": gnd},
	})
}

func (e *Repository) Games(w http.ResponseWriter, r *http.Request) {
	type GameCard struct {
		ID          string
		Title       string
		TitleColor  string
		Image       string
		Description string
		IsAvailable bool
	}
	type GameCardData struct {
		Games []GameCard
	}
	var gc []GameCard
	for _, game := range e.GC.Games {
		card := GameCard{
			ID:          game.ID,
			Title:       game.Title,
			TitleColor:  game.TitleColor,
			Image:       game.Image,
			Description: game.Description,
			IsAvailable: game.IsAvailable,
		}
		gc = append(gc, card)
	}
	render.RenderTemplate(w, "games.page.tmpl", &models.TemplateData{
		Data: map[string]interface{}{"games": gc},
	})
}

func (e *Repository) NotFound(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "notfound.page.tmpl", &models.TemplateData{})
}

// API Routes

// GetNode handler
func (e *Repository) GetNode(w http.ResponseWriter, r *http.Request) {
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("read request failed")
	}
	var rb models.NodeReqBody
	json.Unmarshal(rBody, &rb)
	n, err := helpers.FindRequestedNode(e.GC, rb.NextNode)
	if err != nil {
		en := models.EmptyGameNode
		json.NewEncoder(w).Encode(en)
		return
	}
	cn := helpers.ValidateNodeOptions(n, rb)
	json.NewEncoder(w).Encode(cn)
}
