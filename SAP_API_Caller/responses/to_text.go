package responses

type ToText struct {
	D struct {
		Results []struct {
			Division     string `json:"Division"`
			Language     string `json:"Language"`
			DivisionName string `json:"DivisionName"`
		} `json:"results"`
	} `json:"d"`
}
