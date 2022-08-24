package responses

type SalesQuoteCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                               string `json:"ObjectID"`
			ID                                     string `json:"ID"`
			BuyerID                                string `json:"BuyerID"`
			Name                                   string `json:"Name"`
			ProcessingTypeCode                     string `json:"ProcessingTypeCode"`
			ProcessingTypeCodeText                 string `json:"ProcessingTypeCodeText"`
			BuyerPartyID                           string `json:"BuyerPartyID"`
			BuyerPartyName                         string `json:"BuyerPartyName"`
			BuyerContactPartyID                    string `json:"BuyerContactPartyID"`
			BuyerContactPartyName                  string `json:"BuyerContactPartyName"`
			ProductRecipientPartyID                string `json:"ProductRecipientPartyID"`
			ProductRecipientPartyName              string `json:"ProductRecipientPartyName"`
			EmployeeResponsiblePartyID             string `json:"EmployeeResponsiblePartyID"`
			EmployeeResponsiblePartyName           string `json:"EmployeeResponsiblePartyName"`
			SalesUnitPartyID                       string `json:"SalesUnitPartyID"`
			SalesUnitPartyName                     string `json:"SalesUnitPartyName"`
			SalesOrganisationID                    string `json:"SalesOrganisationID"`
			SalesOrganisationName                  string `json:"SalesOrganisationName"`
			SalesOfficeID                          string `json:"SalesOfficeID"`
			SalesOfficeName                        string `json:"SalesOfficeName"`
			SalesGroupID                           string `json:"SalesGroupID"`
			SalesGroupName                         string `json:"SalesGroupName"`
			DistributionChannelCode                string `json:"DistributionChannelCode"`
			DistributionChannelCodeText            string `json:"DistributionChannelCodeText"`
			DivisionCode                           string `json:"DivisionCode"`
			DivisionCodeText                       string `json:"DivisionCodeText"`
			SalesTerritoryID                       string `json:"SalesTerritoryID"`
			SalesTerritoryName                     string `json:"SalesTerritoryName"`
			DateTime                               string `json:"DateTime"`
			RequestedFulfillmentStartDateTime      string `json:"RequestedFulfillmentStartDateTime"`
			PriceDateTime                          string `json:"PriceDateTime"`
			TimeZoneCode                           string `json:"TimeZoneCode"`
			TimeZoneCodeText                       string `json:"TimeZoneCodeText"`
			ValidFromDate                          string `json:"ValidFromDate"`
			ValidToDate                            string `json:"ValidToDate"`
			CurrencyCode                           string `json:"CurrencyCode"`
			CurrencyCodeText                       string `json:"CurrencyCodeText"`
			DocumentLanguageCode                   string `json:"DocumentLanguageCode"`
			DocumentLanguageCodeText               string `json:"DocumentLanguageCodeText"`
			DeliveryPriorityCode                   string `json:"DeliveryPriorityCode"`
			DeliveryPriorityCodeText               string `json:"DeliveryPriorityCodeText"`
			IncotermsClassificationCode            string `json:"IncotermsClassificationCode"`
			IncotermsClassificationCodeText        string `json:"IncotermsClassificationCodeText"`
			IncotermsTransferLocationName          string `json:"IncotermsTransferLocationName"`
			ProbabilityPercent                     string `json:"ProbabilityPercent"`
			CancellationReasonCode                 string `json:"CancellationReasonCode"`
			CancellationReasonCodeText             string `json:"CancellationReasonCodeText"`
			OrderReasonCode                        string `json:"OrderReasonCode"`
			OrderReasonCodeText                    string `json:"OrderReasonCodeText"`
			MainDiscount                           string `json:"MainDiscount"`
			NetAmount                              string `json:"NetAmount"`
			NetAmountCurrencyCode                  string `json:"NetAmountCurrencyCode"`
			NetAmountCurrencyCodeText              string `json:"NetAmountCurrencyCodeText"`
			GrossAmount                            string `json:"GrossAmount"`
			GrossAmountCurrencyCode                string `json:"GrossAmountCurrencyCode"`
			GrossAmountCurrencyCodeText            string `json:"GrossAmountCurrencyCodeText"`
			TaxAmount                              string `json:"TaxAmount"`
			TaxAmountCurrencyCode                  string `json:"TaxAmountCurrencyCode"`
			TaxAmountCurrencyCodeText              string `json:"TaxAmountCurrencyCodeText"`
			CashDiscountTermsCode                  string `json:"CashDiscountTermsCode"`
			CashDiscountTermsCodeText              string `json:"CashDiscountTermsCodeText"`
			ConfirmationExistenceIndicator         bool   `json:"ConfirmationExistenceIndicator"`
			ConsistencyStatusCode                  string `json:"ConsistencyStatusCode"`
			ConsistencyStatusCodeText              string `json:"ConsistencyStatusCodeText"`
			LifeCycleStatusCode                    string `json:"LifeCycleStatusCode"`
			LifeCycleStatusCodeText                string `json:"LifeCycleStatusCodeText"`
			CancellationStatusCode                 string `json:"CancellationStatusCode"`
			CancellationStatusCodeText             string `json:"CancellationStatusCodeText"`
			ResultStatusCode                       string `json:"ResultStatusCode"`
			ResultStatusCodeText                   string `json:"ResultStatusCodeText"`
			ApprovalStatusCode                     string `json:"ApprovalStatusCode"`
			ApprovalStatusCodeText                 string `json:"ApprovalStatusCodeText"`
			OrderingStatusCode                     string `json:"OrderingStatusCode"`
			OrderingStatusCodeText                 string `json:"OrderingStatusCodeText"`
			CreditWorthinessStatusCode             string `json:"CreditWorthinessStatusCode"`
			CreditWorthinessStatusCodeText         string `json:"CreditWorthinessStatusCodeText"`
			ReplicationProcessingStatusCode        string `json:"ReplicationProcessingStatusCode"`
			ReplicationProcessingStatusCodeText    string `json:"ReplicationProcessingStatusCodeText"`
			ProductAvailabilityStatusCode          string `json:"ProductAvailabilityStatusCode"`
			ProductAvailabilityStatusCodeText      string `json:"ProductAvailabilityStatusCodeText"`
			PriceCalculationStatusCode             string `json:"PriceCalculationStatusCode"`
			PriceCalculationStatusCodeText         string `json:"PriceCalculationStatusCodeText"`
			PricingProcedureCode                   string `json:"PricingProcedureCode"`
			PricingProcedureCodeText               string `json:"PricingProcedureCodeText"`
			ExternalPriceCalculationStatusCode     string `json:"ExternalPriceCalculationStatusCode"`
			ExternalPriceCalculationStatusCodeText string `json:"ExternalPriceCalculationStatusCodeText"`
			ExternalPricingProcedureCode           string `json:"ExternalPricingProcedureCode"`
			ExternalPricingProcedureCodeText       string `json:"ExternalPricingProcedureCodeText"`
			CreationDateTime                       string `json:"CreationDateTime"`
			LastChangeDateTime                     string `json:"LastChangeDateTime"`
			CreatedBy                              string `json:"CreatedBy"`
			LastChangedBy                          string `json:"LastChangedBy"`
			CreationIdentityUUID                   string `json:"CreationIdentityUUID"`
			LastChangeIdentityUUID                 string `json:"LastChangeIdentityUUID"`
			VersionGroupID                         string `json:"VersionGroupID"`
			VersionID                              string `json:"VersionID"`
			External                               bool   `json:"External"`
			Submit                                 string `json:"Submit"`
			SetAsWon                               string `json:"SetAsWon"`
			SetAsLost                              bool   `json:"SetAsLost"`
			RequestExtData                         string `json:"RequestExtData"`
			PrimaryQuote                           bool   `json:"PrimaryQuote"`
			UpdateOpportunity                      string `json:"UpdateOpportunity"`
			RequestExtFollowup                     bool   `json:"RequestExtFollowup"`
			GrossWeightMeasure                     string `json:"GrossWeightMeasure"`
			GrossWeightUnitCode                    string `json:"GrossWeightUnitCode"`
			GrossWeightUnitCodeText                string `json:"GrossWeightUnitCodeText"`
			NetWeightMeasure                       string `json:"NetWeightMeasure"`
			NetWeightUnitCode                      string `json:"NetWeightUnitCode"`
			NetWeightUnitCodeText                  string `json:"NetWeightUnitCodeText"`
			VolumeMeasure                          string `json:"VolumeMeasure"`
			VolumeUnitCode                         string `json:"VolumeUnitCode"`
			VolumeUnitCodeText                     string `json:"VolumeUnitCodeText"`
			ExternalApprovalStatusCode             string `json:"ExternalApprovalStatusCode"`
			ExternalApprovalStatusCodeText         string `json:"ExternalApprovalStatusCodeText"`
			PlantPartyID                           string `json:"PlantPartyID"`
			PlantPartyName                         string `json:"PlantPartyName"`
			EntityLastChangedOn                    string `json:"EntityLastChangedOn"`
			ETag                                   string `json:"ETag"`
			DataloaderKUT                          string `json:"dataloader_KUT"`
			ObjectIdentifierMapping                struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ObjectIdentifierMapping"`
			SalesQuoteAttachmentFolder struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuoteAttachmentFolder"`
			SalesQuoteExternalPrice struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuoteExternalPrice"`
			SalesQuoteItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuoteItem"`
			SalesQuoteOutput struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuoteOutput"`
			SalesQuoteParty struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuoteParty"`
			SalesQuotePrice struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuotePrice"`
			SalesQuoteReference struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuoteReference"`
			SalesQuoteReferenceObject struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuoteReferenceObject"`
			SalesQuoteTextCollection struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuoteTextCollection"`
			SalesQuoteWorklistItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuoteWorklistItem"`
		} `json:"results"`
	} `json:"d"`
}
