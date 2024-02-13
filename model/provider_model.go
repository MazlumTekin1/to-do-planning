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

type Provider3Model struct {
	Zorluk3 int    `json:"zorluk"`
	Sure3   int    `json:"sure"`
	Id3     string `id:"id"`
}
type Developer struct {
	Id                          int
	Name                        string
	DeveloperWorkHourDifficulty int
}
