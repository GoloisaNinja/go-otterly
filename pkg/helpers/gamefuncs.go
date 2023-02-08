package helpers

import (
	"github.com/GoloisaNinja/go-otterly/pkg/config"
	"github.com/GoloisaNinja/go-otterly/pkg/games"
	"github.com/GoloisaNinja/go-otterly/pkg/models"
)

var (
	gameslice = []models.Game{
		games.OneWayOrTheOtter,
		games.AgeOfTheOtter,
		games.AnotterTime,
	}
)

func LoadGames(gc *config.GameConfig) {
	gc.Games = append(gc.Games, gameslice...)
}

func LoadGame(gc *config.GameConfig, id string) {
	for _, game := range gc.Games {
		if game.ID == id {
			gc.Game = game
		}
	}
}
func findIndex(a []string, v string) int {
	index := -1
	for i, str := range a {
		if str == v {
			index = i
		}
	}
	return index
}
func ValidateNodeOptions(n models.GameNode, rb models.NodeReqBody) models.GameNode {
	cleanNode := n
	cleanNode.Options = []models.GameOptions{}
	for _, op := range n.Options {
		if !op.HasReqs {
			cleanNode.Options = append(cleanNode.Options, op)
			continue
		}
		if op.Requires.Mood != "" && op.Requires.Mood != rb.Status {
			continue
		}
		if len(op.Requires.Inventory) > 0 {
			im := true
			for _, inv := range op.Requires.Inventory {
				index := findIndex(rb.Inventory, inv)
				if index == -1 {
					im = false
				}
			}
			if !im {
				continue
			}
		}
		cleanNode.Options = append(cleanNode.Options, op)
	}
	return cleanNode
}
