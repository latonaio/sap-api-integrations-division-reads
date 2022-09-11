package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-division-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetDivision(division, language, divisionName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Division":
			func() {
				c.Division(division)
				wg.Done()
			}()
		case "Text":
			func() {
				c.Text(language, divisionName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Division(division string) {
	divisionData, err := c.callDivisionSrvAPIRequirementDivision("A_Division", division)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(divisionData)

	textData, err := c.callToText(divisionData[0].ToText)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(textData)
}

func (c *SAPAPICaller) callDivisionSrvAPIRequirementDivision(api, division string) ([]sap_api_output_formatter.Division, error) {
	url := strings.Join([]string{c.baseURL, "API_DIVISION_SRV", api}, "/")
	param := c.getQueryWithDivision(map[string]string{}, division)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToDivision(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToText(url string) ([]sap_api_output_formatter.Text, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Text(language, divisionName string) {
	data, err := c.callDivisionSrvAPIRequirementText("A_DivisionText", language, divisionName)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callDivisionSrvAPIRequirementText(api, language, divisionName string) ([]sap_api_output_formatter.Text, error) {
	url := strings.Join([]string{c.baseURL, "API_DIVISION_SRV", api}, "/")

	param := c.getQueryWithText(map[string]string{}, language, divisionName)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}


func (c *SAPAPICaller) getQueryWithDivision(params map[string]string, division string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Division eq '%s'", division)
	return params
}

func (c *SAPAPICaller) getQueryWithText(params map[string]string, language, divisionName string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Language eq '%s' and substringof('%s', DivisionName)", language, divisionName)
	return params
}
