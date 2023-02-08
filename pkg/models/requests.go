package models

type NodeReqBody struct {
	NextNode  string   `json:"nextNode"`
	Status    string   `json:"status"`
	Inventory []string `json:"inventory"`
}
