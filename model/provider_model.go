package model

type Task struct {
	Name       string
	Duration   int
	Difficulty int
}

type Provider1Model struct {
	Zorluk int    `json:"zorluk"`
	Sure   int    `json:"sure"`
	Id     string `id:"id"`
}

type Provider2Model struct {
	Value             int    `json:"value"`
	EstimatedDuration int    `json:"estimated_duration"`
	Id                string `id:"id"`
}
