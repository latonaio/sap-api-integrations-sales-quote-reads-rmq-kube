package responses

type SalesQuotePartyCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                                 string `json:"ObjectID"`
			ParentObjectID                           string `json:"ParentObjectID"`
			QuoteID                                  string `json:"QuoteID"`
			PartyID                                  string `json:"PartyID"`
			RoleCode                                 string `json:"RoleCode"`
			RoleCodeText                             string `json:"RoleCodeText"`
			MainIndicator                            bool   `json:"MainIndicator"`
			PartyContactPartyID                      string `json:"PartyContactPartyID"`
			PartyName                                string `json:"PartyName"`
			PartyContactPartyName                    string `json:"PartyContactPartyName"`
			CountryCode                              string `json:"CountryCode"`
			CountryCodeText                          string `json:"CountryCodeText"`
			RegionCode                               string `json:"RegionCode"`
			RegionCodeText                           string `json:"RegionCodeText"`
			CareOfName                               string `json:"CareOfName"`
			StreetPrefixName                         string `json:"StreetPrefixName"`
			AdditionalStreetPrefixName               string `json:"AdditionalStreetPrefixName"`
			HouseID                                  string `json:"HouseID"`
			StreetName                               string `json:"StreetName"`
			StreetSuffixName                         string `json:"StreetSuffixName"`
			AdditionalStreetSuffixName               string `json:"AdditionalStreetSuffixName"`
			DistrictName                             string `json:"DistrictName"`
			CityName                                 string `json:"CityName"`
			AdditionalCityName                       string `json:"AdditionalCityName"`
			StreetPostalCode                         string `json:"StreetPostalCode"`
			StreetPostalCodeText                     string `json:"StreetPostalCodeText"`
			CountyName                               string `json:"CountyName"`
			CompanyPostalCode                        string `json:"CompanyPostalCode"`
			CompanyPostalCodeText                    string `json:"CompanyPostalCodeText"`
			POBoxIndicator                           bool   `json:"POBoxIndicator"`
			POBoxID                                  string `json:"POBoxID"`
			POBoxPostalCode                          string `json:"POBoxPostalCode"`
			POBoxPostalCodeText                      string `json:"POBoxPostalCodeText"`
			POBoxDeviatingCountryCode                string `json:"POBoxDeviatingCountryCode"`
			POBoxDeviatingCountryCodeText            string `json:"POBoxDeviatingCountryCodeText"`
			POBoxDeviatingRegionCode                 string `json:"POBoxDeviatingRegionCode"`
			POBoxDeviatingRegionCodeText             string `json:"POBoxDeviatingRegionCodeText"`
			POBoxDeviatingCityName                   string `json:"POBoxDeviatingCityName"`
			BuildingID                               string `json:"BuildingID"`
			FloorID                                  string `json:"FloorID"`
			RoomID                                   string `json:"RoomID"`
			Phone                                    string `json:"Phone"`
			Mobile                                   string `json:"Mobile"`
			Fax                                      string `json:"Fax"`
			Email                                    string `json:"Email"`
			Web                                      string `json:"Web"`
			CorrespondenceLanguageCode               string `json:"CorrespondenceLanguageCode"`
			CorrespondenceLanguageCodeText           string `json:"CorrespondenceLanguageCodeText"`
			PreferredCommunicationMediumTypeCode     string `json:"PreferredCommunicationMediumTypeCode"`
			PreferredCommunicationMediumTypeCodeText string `json:"PreferredCommunicationMediumTypeCodeText"`
			FirstLineName                            string `json:"FirstLineName"`
			SecondLineName                           string `json:"SecondLineName"`
			ThirdLineName                            string `json:"ThirdLineName"`
			FourthLineName                           string `json:"FourthLineName"`
			ETag                                     string `json:"ETag"`
			ObjectIdentifierMapping                  struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ObjectIdentifierMapping"`
			SalesQuote struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuote"`
		} `json:"results"`
	} `json:"d"`
}
