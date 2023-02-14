package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/GoloisaNinja/go-otterly/pkg/config"
	"github.com/GoloisaNinja/go-otterly/pkg/helpers"
	"github.com/GoloisaNinja/go-otterly/pkg/models"
	"github.com/GoloisaNinja/go-otterly/pkg/render"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
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

func (e *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
func (e *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
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
		GameID      string
		Title       string
		Description string
		Node        models.GameNode
		TotalNodes  int
	}
	gnd := InitialGameData{
		GameID:      e.GC.Game.ID,
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

func ThankYou(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "thankyou.page.tmpl", &models.TemplateData{})
}
func SubmissionError(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "suberror.page.tmpl", &models.TemplateData{})
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

// HandleContactSubmission handler which will process contact requests from the client
func HandleContactSubmission(w http.ResponseWriter, r *http.Request) {
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("read request failed")
	}
	type ContactFormData struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}
	var cd ContactFormData
	json.Unmarshal(rBody, &cd)
	if cd.Subject != "" {
		// means a bot submitted our form from the client - do nothing
		// send fake 200 resp
		var fr models.StandardResponse
		fr.Build(200, "ok", "no data")
		json.NewEncoder(w).Encode(fr)
		return
	}
	var sessErr models.StandardResponse
	var sendErr models.StandardResponse
	var emailOk models.StandardResponse
	// create SendEmail arguments
	emailText := cd.Message
	subject := "Contact Request From Otterly"
	sender := cd.Email
	toAddresses := []*string{
		aws.String("jonathan.collins@live.com"),
	}
	AKID := os.Getenv("AWS_ACCESS_KEY_ID")
	SAK := os.Getenv("AWS_SECRET_ACCESS_KEY")
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region:      aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(AKID, SAK, ""),
		},
	})
	if err != nil {
		sessErr.Build(500, "failed to create aws session", "no data")
		json.NewEncoder(w).Encode(sessErr)
		return
	}
	err = helpers.SendEmail(sess, toAddresses, emailText, sender, subject)
	if err != nil {
		fmt.Println(err)
		sendErr.Build(500, "good session - but email send error", "no data")
		json.NewEncoder(w).Encode(sendErr)
		return
	}
	emailOk.Build(200, "email sent successfully", "no data")
	json.NewEncoder(w).Encode(emailOk)
}
