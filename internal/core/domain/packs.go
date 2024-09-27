package domain

type Packs struct {
	Quantity int `json:"quantity"`
	Value    int `json:"value"`
}

type PackList struct {
	Packs []Packs `json:"packs"`
}
