package models

type NodeReqBody struct {
	GameID    string   `json:"gameId"`
	NextNode  string   `json:"nextNode"`
	Status    string   `json:"status"`
	Inventory []string `json:"inventory"`
}
