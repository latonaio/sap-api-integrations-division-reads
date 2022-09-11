package responses

type Division struct {
	D struct {
		Results []struct {
			Division     string `json:"Division"`
			ToText struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Text"`
		} `json:"results"`
	} `json:"d"`
}