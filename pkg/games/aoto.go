package games

import "github.com/GoloisaNinja/go-otterly/pkg/models"

var AgeOfTheOtter = models.Game{
	ID:          "1B",
	Title:       "Age of the Otter",
	TitleColor:  "#4fd3d3",
	Image:       "/static/images/ageoftheotter.webp",
	Description: "**Coming Soon** The land of OtterMoragh has been plunged into darkness. An ancient evil has awoken from it's deep and terrible tomb. Buried there by your ancestors, it has woken with only vengeance in mind and returned to OtterMoragh. An old mystic foretold of this day, and foretold that only you could save OtterMoragh from being lost to otter chaos. Get your bag of infinite holding, you're going on an adventure!",
	IsAvailable: false,
	Nodes:       []models.GameNode{},
}
