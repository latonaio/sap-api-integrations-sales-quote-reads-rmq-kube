package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-sales-quote-reads-rmq-kube/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type SAPAPICaller struct {
	baseURL      string
	apiKey       string
	outputQueues []string
	outputter    RMQOutputter
	log          *logger.Logger
}

func NewSAPAPICaller(baseUrl string, outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:      baseUrl,
		apiKey:       GetApiKey(),
		outputQueues: outputQueueTo,
		outputter:    outputter,
		log:          l,
	}
}

func (c *SAPAPICaller) AsyncGetSalesQuote(iD, itemID, partyID, quoteItemID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "SalesQuoteCollection":
			func() {
				c.SalesQuoteCollection(iD)
				wg.Done()
			}()
		case "SalesQuoteItemCollection":
			func() {
				c.SalesQuoteItemCollection(itemID)
				wg.Done()
			}()
		case "SalesQuotePartyCollection":
			func() {
				c.SalesQuotePartyCollection(partyID)
				wg.Done()
			}()
		case "SalesQuoteItemScheduleLineCollection":
			func() {
				c.SalesQuoteItemScheduleLineCollection(quoteItemID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) SalesQuoteCollection(iD string) {
	data, err := c.callSalesQuoteSrvAPIRequirementSalesQuoteCollection("SalesQuoteCollectionData", iD)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "SalesQuoteCollectionData"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callSalesQuoteSrvAPIRequirementSalesQuoteCollection(api, iD string) ([]sap_api_output_formatter.SalesQuoteCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithSalesQuoteCollection(req, iD)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSalesQuoteCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SalesQuoteItemCollection(itemID string) {
	data, err := c.callSalesQuoteSrvAPIRequirementSalesQuoteItemCollection("SalesQuoteItemCollectionData", itemID)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "SalesQuoteItemCollectionData"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callSalesQuoteSrvAPIRequirementSalesQuoteItemCollection(api, itemID string) ([]sap_api_output_formatter.SalesQuoteItemCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithSalesQuoteItemCollection(req, itemID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSalesQuoteItemCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SalesQuotePartyCollection(partyID string) {
	data, err := c.callSalesQuoteSrvAPIRequirementSalesQuotePartyCollection("SalesQuotePartyCollectionData", partyID)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "SalesQuotePartyCollectionData"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callSalesQuoteSrvAPIRequirementSalesQuotePartyCollection(api, partyID string) ([]sap_api_output_formatter.SalesQuotePartyCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithSalesQuotePartyCollection(req, partyID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSalesQuotePartyCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SalesQuoteItemScheduleLineCollection(quoteItemID string) {
	data, err := c.callSalesQuoteSrvAPIRequirementSalesQuoteItemScheduleLineCollection("SalesQuoteItemScheduleLineCollectionData", quoteItemID)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "SalesQuoteItemScheduleLineCollectionData"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callSalesQuoteSrvAPIRequirementSalesQuoteItemScheduleLineCollection(api, quoteItemID string) ([]sap_api_output_formatter.SalesQuoteItemScheduleLineCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithSalesQuoteItemScheduleLineCollection(req, quoteItemID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSalesQuoteItemScheduleLineCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithSalesQuoteCollection(req *http.Request, iD string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ID eq '%s'", iD))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithSalesQuoteItemCollection(req *http.Request, itemID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ItemID eq '%s'", itemID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithSalesQuotePartyCollection(req *http.Request, partyID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PartyID eq '%s'", partyID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithSalesQuoteItemScheduleLineCollection(req *http.Request, quoteItemID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("QuoteItemID eq '%s'", quoteItemID))
	req.URL.RawQuery = params.Encode()
}
