package sap_api_output_formatter

type DivisionReads struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	Division      string `json:"division"`
	Deleted       string `json:"deleted"`
}

type Division struct {
	Division     string `json:"Division"`
	ToText		 string `json:"to_Text"`
}

type Text struct {
	Division     string `json:"Division"`
	Language     string `json:"Language"`
	DivisionName string `json:"DivisionName"`
}
