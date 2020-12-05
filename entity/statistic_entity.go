package entity

type StatisticPerYear struct {
	YearAndMon string `json:"year_and_month"`
	Profit     int64  `json:"profit"`
}

type StatisticPerYearReq struct {
	Year string `json:"year"`
}
