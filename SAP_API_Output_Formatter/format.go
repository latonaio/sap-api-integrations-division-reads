package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-division-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToDivision(raw []byte, l *logger.Logger) ([]Division, error) {
	pm := &responses.Division{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Division. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	division := make([]Division, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		division = append(division, Division{
			Division:    data.Division,
			ToText:      data.ToText.Deferred.URI,
		})
	}

	return division, nil
}

func ConvertToText(raw []byte, l *logger.Logger) ([]Text, error) {
	pm := &responses.Text{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Text. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	text := make([]Text, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		text = append(text, Text{
			Division:     data.Division,
			Language:     data.Language,
			DivisionName: data.DivisionName,
		})
	}

	return text, nil
}
