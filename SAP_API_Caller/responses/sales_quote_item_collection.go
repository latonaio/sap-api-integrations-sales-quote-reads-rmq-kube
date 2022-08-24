package responses

type SalesQuoteItemCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                          string `json:"ObjectID"`
			ParentObjectID                    string `json:"ParentObjectID"`
			QuoteID                           string `json:"QuoteID"`
			ItemID                            string `json:"ItemID"`
			ParentItemID                      string `json:"ParentItemID"`
			AlternativeToItemID               string `json:"AlternativeToItemID"`
			ProcessingTypeCode                string `json:"ProcessingTypeCode"`
			ProcessingTypeCodeText            string `json:"ProcessingTypeCodeText"`
			Description                       string `json:"Description"`
			ProductInternalID                 string `json:"ProductInternalID"`
			ProductBuyerID                    string `json:"ProductBuyerID"`
			ProductStandardID                 string `json:"ProductStandardID"`
			ProductCategoryID                 string `json:"ProductCategoryID"`
			Quantity                          string `json:"Quantity"`
			QuantityMeasureUnitCode           string `json:"QuantityMeasureUnitCode"`
			QuantityMeasureUnitCodeText       string `json:"QuantityMeasureUnitCodeText"`
			StartDateTime                     string `json:"StartDateTime"`
			TimeZoneCode                      string `json:"TimeZoneCode"`
			TimeZoneCodeText                  string `json:"TimeZoneCodeText"`
			ProductRecipientPartyID           string `json:"ProductRecipientPartyID"`
			ProductRecipientPartyName         string `json:"ProductRecipientPartyName"`
			CancellationReasonCode            string `json:"CancellationReasonCode"`
			CancellationReasonCodeText        string `json:"CancellationReasonCodeText"`
			MainPrice                         string `json:"MainPrice"`
			MainPriceCurrencyCode             string `json:"MainPriceCurrencyCode"`
			MainPriceCurrencyCodeText         string `json:"MainPriceCurrencyCodeText"`
			MainPriceBaseQuantity             string `json:"MainPriceBaseQuantity"`
			MainPriceBaseQuantityUnitCode     string `json:"MainPriceBaseQuantityUnitCode"`
			MainPriceBaseQuantityUnitCodeText string `json:"MainPriceBaseQuantityUnitCodeText"`
			MainDiscount                      string `json:"MainDiscount"`
			NetPriceAmount                    string `json:"NetPriceAmount"`
			NetPriceAmountCurrencyCode        string `json:"NetPriceAmountCurrencyCode"`
			NetPriceAmountCurrencyCodeText    string `json:"NetPriceAmountCurrencyCodeText"`
			NetPriceBaseQuantity              string `json:"NetPriceBaseQuantity"`
			NetPriceBaseQuantityUnitCode      string `json:"NetPriceBaseQuantityUnitCode"`
			NetPriceBaseQuantityUnitCodeText  string `json:"NetPriceBaseQuantityUnitCodeText"`
			NetAmount                         string `json:"NetAmount"`
			NetAmountCurrencyCode             string `json:"NetAmountCurrencyCode"`
			NetAmountCurrencyCodeText         string `json:"NetAmountCurrencyCodeText"`
			GrossAmount                       string `json:"GrossAmount"`
			GrossAmountCurrencyCode           string `json:"GrossAmountCurrencyCode"`
			GrossAmountCurrencyCodeText       string `json:"GrossAmountCurrencyCodeText"`
			TaxAmount                         string `json:"TaxAmount"`
			TaxAmountCurrencyCode             string `json:"TaxAmountCurrencyCode"`
			TaxAmountCurrencyCodeText         string `json:"TaxAmountCurrencyCodeText"`
			CashDiscountTermsCode             string `json:"CashDiscountTermsCode"`
			CashDiscountTermsCodeText         string `json:"CashDiscountTermsCodeText"`
			CancellationStatusCode            string `json:"CancellationStatusCode"`
			CancellationStatusCodeText        string `json:"CancellationStatusCodeText"`
			ConsistencyStatusCode             string `json:"ConsistencyStatusCode"`
			ConsistencyStatusCodeText         string `json:"ConsistencyStatusCodeText"`
			LifeCycleStatusCode               string `json:"LifeCycleStatusCode"`
			LifeCycleStatusCodeText           string `json:"LifeCycleStatusCodeText"`
			OrderingStatusCode                string `json:"OrderingStatusCode"`
			OrderingStatusCodeText            string `json:"OrderingStatusCodeText"`
			ConfirmationExistenceIndicator    bool   `json:"ConfirmationExistenceIndicator"`
			CreationDateTime                  string `json:"CreationDateTime"`
			LastChangeDateTime                string `json:"LastChangeDateTime"`
			DeliveryPriorityCode              string `json:"DeliveryPriorityCode"`
			DeliveryPriorityCodeText          string `json:"DeliveryPriorityCodeText"`
			IncotermsClassificationCode       string `json:"IncotermsClassificationCode"`
			IncotermsClassificationCodeText   string `json:"IncotermsClassificationCodeText"`
			IncotermsTransferLocationName     string `json:"IncotermsTransferLocationName"`
			GrossWeightMeasure                string `json:"GrossWeightMeasure"`
			GrossWeightUnitCode               string `json:"GrossWeightUnitCode"`
			GrossWeightUnitCodeText           string `json:"GrossWeightUnitCodeText"`
			NetWeightMeasure                  string `json:"NetWeightMeasure"`
			NetWeightUnitCode                 string `json:"NetWeightUnitCode"`
			NetWeightUnitCodeText             string `json:"NetWeightUnitCodeText"`
			VolumeMeasure                     string `json:"VolumeMeasure"`
			VolumeUnitCode                    string `json:"VolumeUnitCode"`
			VolumeUnitCodeText                string `json:"VolumeUnitCodeText"`
			OrderedQuantity                   string `json:"OrderedQuantity"`
			OrderedQuantityUnitCode           string `json:"OrderedQuantityUnitCode"`
			OrderedQuantityUnitCodeText       string `json:"OrderedQuantityUnitCodeText"`
			PricingProductID                  string `json:"PricingProductID"`
			PlantPartyID                      string `json:"PlantPartyID"`
			PlantPartyName                    string `json:"PlantPartyName"`
			EntityLastChangedOn               string `json:"EntityLastChangedOn"`
			ETag                              string `json:"ETag"`
			ObjectIdentifierMapping           struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ObjectIdentifierMapping"`
			SalesQuote struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuote"`
			SalesQuoteExternalPriceItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuoteExternalPriceItem"`
			SalesQuoteItemAttachmentFolder struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuoteItemAttachmentFolder"`
			SalesQuoteItemParty struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuoteItemParty"`
			SalesQuoteItemReference struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuoteItemReference"`
			SalesQuoteItemScheduleLine struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuoteItemScheduleLine"`
			SalesQuoteItemTextCollection struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuoteItemTextCollection"`
			SalesQuotePriceItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"SalesQuotePriceItem"`
		} `json:"results"`
	} `json:"d"`
}
