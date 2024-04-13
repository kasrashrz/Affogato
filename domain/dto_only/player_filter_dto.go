package dto_only

type PlayerFilters struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	Operation string `json:"operation"`
}

type AutoGenerated struct {
	PlayerFilters []PlayerFilters `json:"player_filters"`
}
