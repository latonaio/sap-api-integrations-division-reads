package responses

type Text struct {
	D struct {
		Results []struct {
			Division     string `json:"Division"`
			Language     string `json:"Language"`
			DivisionName string `json:"DivisionName"`
		} `json:"results"`
	} `json:"d"`
}
