package responses

type SalesQuoteItemScheduleLineCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID         string `json:"ObjectID"`
			ParentObjectID   string `json:"ParentObjectID"`
			QuoteID          string `json:"QuoteID"`
			QuoteItemID      string `json:"QuoteItemID"`
			ID               string `json:"ID"`
			TypeCode         string `json:"TypeCode"`
			TypeCodeText     string `json:"TypeCodeText"`
			Quantity         string `json:"Quantity"`
			UnitCode         string `json:"unitCode"`
			UnitCodeText     string `json:"unitCodeText"`
			StartDateTime    string `json:"StartDateTime"`
			EndDateTime      string `json:"EndDateTime"`
			TimeZoneCode     string `json:"timeZoneCode"`
			TimeZoneCodeText string `json:"timeZoneCodeText"`
			ETag             string `json:"ETag"`
			SalesQuote       struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuote"`
			SalesQuoteItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuoteItem"`
		} `json:"results"`
	} `json:"d"`
}