package helpers

import (
	"github.com/GoloisaNinja/go-otterly/pkg/config"
	"github.com/GoloisaNinja/go-otterly/pkg/models"
)

func ResetState(app *config.GameConfig) {
	dgs := models.GameState{
		NodesCompleted:   0,
		Survived:         true,
		DominantPlayType: "...undetermined",
	}
	app.GameState = dgs
}
