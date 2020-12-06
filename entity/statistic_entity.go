package entity

type StatisticPerYear struct {
	YearAndMon string `json:"year_and_month"`
	Profit     int64  `json:"profit"`
}

type StatisticPerMon struct {
	Mon    string `json:"mon"`
	Profit int64  `json:"profit"`
}

type StatisticPerMonRes struct {
	Year   string            `json:"year"`
	Detail []StatisticPerMon `json:"detail"`
}

type StatisticPerYearReq struct {
	Year string `json:"year"`
}
