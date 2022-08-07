package model

//WebInfo .
type WebInfo struct {
	ID          string `json:"id"`
	SortNumber  int    `json:"sort"`
	Name        string `json:"name"`
	IntnetURL   string `json:"intnetURL"`
	IntranetURL string `json:"intranetURL"`
	URL         string `json:"URL"`
	ICON        string `json:"icon"`
}
