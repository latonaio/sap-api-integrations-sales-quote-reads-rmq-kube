package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-sales-quote-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToSalesQuoteCollection(raw []byte, l *logger.Logger) ([]SalesQuoteCollection, error) {
	pm := &responses.SalesQuoteCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesQuoteCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	salesQuoteCollection := make([]SalesQuoteCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		salesQuoteCollection = append(salesQuoteCollection, SalesQuoteCollection{
			ObjectID:                               data.ObjectID,
			ID:                                     data.ID,
			BuyerID:                                data.BuyerID,
			Name:                                   data.Name,
			ProcessingTypeCode:                     data.ProcessingTypeCode,
			ProcessingTypeCodeText:                 data.ProcessingTypeCodeText,
			BuyerPartyID:                           data.BuyerPartyID,
			BuyerPartyName:                         data.BuyerPartyName,
			BuyerContactPartyID:                    data.BuyerContactPartyID,
			BuyerContactPartyName:                  data.BuyerContactPartyName,
			ProductRecipientPartyID:                data.ProductRecipientPartyID,
			ProductRecipientPartyName:              data.ProductRecipientPartyName,
			EmployeeResponsiblePartyID:             data.EmployeeResponsiblePartyID,
			EmployeeResponsiblePartyName:           data.EmployeeResponsiblePartyName,
			SalesUnitPartyID:                       data.SalesUnitPartyID,
			SalesUnitPartyName:                     data.SalesUnitPartyName,
			SalesOrganisationID:                    data.SalesOrganisationID,
			SalesOrganisationName:                  data.SalesOrganisationName,
			SalesOfficeID:                          data.SalesOfficeID,
			SalesOfficeName:                        data.SalesOfficeName,
			SalesGroupID:                           data.SalesGroupID,
			SalesGroupName:                         data.SalesGroupName,
			DistributionChannelCode:                data.DistributionChannelCode,
			DistributionChannelCodeText:            data.DistributionChannelCodeText,
			DivisionCode:                           data.DivisionCode,
			DivisionCodeText:                       data.DivisionCodeText,
			SalesTerritoryID:                       data.SalesTerritoryID,
			SalesTerritoryName:                     data.SalesTerritoryName,
			DateTime:                               data.DateTime,
			RequestedFulfillmentStartDateTime:      data.RequestedFulfillmentStartDateTime,
			PriceDateTime:                          data.PriceDateTime,
			TimeZoneCode:                           data.TimeZoneCode,
			TimeZoneCodeText:                       data.TimeZoneCodeText,
			ValidFromDate:                          data.ValidFromDate,
			ValidToDate:                            data.ValidToDate,
			CurrencyCode:                           data.CurrencyCode,
			CurrencyCodeText:                       data.CurrencyCodeText,
			DocumentLanguageCode:                   data.DocumentLanguageCode,
			DocumentLanguageCodeText:               data.DocumentLanguageCodeText,
			DeliveryPriorityCode:                   data.DeliveryPriorityCode,
			DeliveryPriorityCodeText:               data.DeliveryPriorityCodeText,
			IncotermsClassificationCode:            data.IncotermsClassificationCode,
			IncotermsClassificationCodeText:        data.IncotermsClassificationCodeText,
			IncotermsTransferLocationName:          data.IncotermsTransferLocationName,
			ProbabilityPercent:                     data.ProbabilityPercent,
			CancellationReasonCode:                 data.CancellationReasonCode,
			CancellationReasonCodeText:             data.CancellationReasonCodeText,
			OrderReasonCode:                        data.OrderReasonCode,
			OrderReasonCodeText:                    data.OrderReasonCodeText,
			MainDiscount:                           data.MainDiscount,
			NetAmount:                              data.NetAmount,
			NetAmountCurrencyCode:                  data.NetAmountCurrencyCode,
			NetAmountCurrencyCodeText:              data.NetAmountCurrencyCodeText,
			GrossAmount:                            data.GrossAmount,
			GrossAmountCurrencyCode:                data.GrossAmountCurrencyCode,
			GrossAmountCurrencyCodeText:            data.GrossAmountCurrencyCodeText,
			TaxAmount:                              data.TaxAmount,
			TaxAmountCurrencyCode:                  data.TaxAmountCurrencyCode,
			TaxAmountCurrencyCodeText:              data.TaxAmountCurrencyCodeText,
			CashDiscountTermsCode:                  data.CashDiscountTermsCode,
			CashDiscountTermsCodeText:              data.CashDiscountTermsCodeText,
			ConfirmationExistenceIndicator:         data.ConfirmationExistenceIndicator,
			ConsistencyStatusCode:                  data.ConsistencyStatusCode,
			ConsistencyStatusCodeText:              data.ConsistencyStatusCodeText,
			LifeCycleStatusCode:                    data.LifeCycleStatusCode,
			LifeCycleStatusCodeText:                data.LifeCycleStatusCodeText,
			CancellationStatusCode:                 data.CancellationStatusCode,
			CancellationStatusCodeText:             data.CancellationStatusCodeText,
			ResultStatusCode:                       data.ResultStatusCode,
			ResultStatusCodeText:                   data.ResultStatusCodeText,
			ApprovalStatusCode:                     data.ApprovalStatusCode,
			ApprovalStatusCodeText:                 data.ApprovalStatusCodeText,
			OrderingStatusCode:                     data.OrderingStatusCode,
			OrderingStatusCodeText:                 data.OrderingStatusCodeText,
			CreditWorthinessStatusCode:             data.CreditWorthinessStatusCode,
			CreditWorthinessStatusCodeText:         data.CreditWorthinessStatusCodeText,
			ReplicationProcessingStatusCode:        data.ReplicationProcessingStatusCode,
			ReplicationProcessingStatusCodeText:    data.ReplicationProcessingStatusCodeText,
			ProductAvailabilityStatusCode:          data.ProductAvailabilityStatusCode,
			ProductAvailabilityStatusCodeText:      data.ProductAvailabilityStatusCodeText,
			PriceCalculationStatusCode:             data.PriceCalculationStatusCode,
			PriceCalculationStatusCodeText:         data.PriceCalculationStatusCodeText,
			PricingProcedureCode:                   data.PricingProcedureCode,
			PricingProcedureCodeText:               data.PricingProcedureCodeText,
			ExternalPriceCalculationStatusCode:     data.ExternalPriceCalculationStatusCode,
			ExternalPriceCalculationStatusCodeText: data.ExternalPriceCalculationStatusCodeText,
			ExternalPricingProcedureCode:           data.ExternalPricingProcedureCode,
			ExternalPricingProcedureCodeText:       data.ExternalPricingProcedureCodeText,
			CreationDateTime:                       data.CreationDateTime,
			LastChangeDateTime:                     data.LastChangeDateTime,
			CreatedBy:                              data.CreatedBy,
			LastChangedBy:                          data.LastChangedBy,
			CreationIdentityUUID:                   data.CreationIdentityUUID,
			LastChangeIdentityUUID:                 data.LastChangeIdentityUUID,
			VersionGroupID:                         data.VersionGroupID,
			VersionID:                              data.VersionID,
			External:                               data.External,
			Submit:                                 data.Submit,
			SetAsWon:                               data.SetAsWon,
			SetAsLost:                              data.SetAsLost,
			RequestExtData:                         data.RequestExtData,
			PrimaryQuote:                           data.PrimaryQuote,
			UpdateOpportunity:                      data.UpdateOpportunity,
			RequestExtFollowup:                     data.RequestExtFollowup,
			GrossWeightMeasure:                     data.GrossWeightMeasure,
			GrossWeightUnitCode:                    data.GrossWeightUnitCode,
			GrossWeightUnitCodeText:                data.GrossWeightUnitCodeText,
			NetWeightMeasure:                       data.NetWeightMeasure,
			NetWeightUnitCode:                      data.NetWeightUnitCode,
			NetWeightUnitCodeText:                  data.NetWeightUnitCodeText,
			VolumeMeasure:                          data.VolumeMeasure,
			VolumeUnitCode:                         data.VolumeUnitCode,
			VolumeUnitCodeText:                     data.VolumeUnitCodeText,
			ExternalApprovalStatusCode:             data.ExternalApprovalStatusCode,
			ExternalApprovalStatusCodeText:         data.ExternalApprovalStatusCodeText,
			PlantPartyID:                           data.PlantPartyID,
			PlantPartyName:                         data.PlantPartyName,
			EntityLastChangedOn:                    data.EntityLastChangedOn,
			ETag:                                   data.ETag,
			DataloaderKUT:                          data.DataloaderKUT,
		})
	}

	return salesQuoteCollection, nil
}

func ConvertToSalesQuoteItemCollection(raw []byte, l *logger.Logger) ([]SalesQuoteItemCollection, error) {
	pm := &responses.SalesQuoteItemCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesQuoteItemCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	salesQuoteItemCollection := make([]SalesQuoteItemCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		salesQuoteItemCollection = append(salesQuoteItemCollection, SalesQuoteItemCollection{
			ObjectID:                          data.ObjectID,
			ParentObjectID:                    data.ParentObjectID,
			QuoteID:                           data.QuoteID,
			ItemID:                            data.ItemID,
			ParentItemID:                      data.ParentItemID,
			AlternativeToItemID:               data.AlternativeToItemID,
			ProcessingTypeCode:                data.ProcessingTypeCode,
			ProcessingTypeCodeText:            data.ProcessingTypeCodeText,
			Description:                       data.Description,
			ProductInternalID:                 data.ProductInternalID,
			ProductBuyerID:                    data.ProductBuyerID,
			ProductStandardID:                 data.ProductStandardID,
			ProductCategoryID:                 data.ProductCategoryID,
			Quantity:                          data.Quantity,
			QuantityMeasureUnitCode:           data.QuantityMeasureUnitCode,
			QuantityMeasureUnitCodeText:       data.QuantityMeasureUnitCodeText,
			StartDateTime:                     data.StartDateTime,
			TimeZoneCode:                      data.TimeZoneCode,
			TimeZoneCodeText:                  data.TimeZoneCodeText,
			ProductRecipientPartyID:           data.ProductRecipientPartyID,
			ProductRecipientPartyName:         data.ProductRecipientPartyName,
			CancellationReasonCode:            data.CancellationReasonCode,
			CancellationReasonCodeText:        data.CancellationReasonCodeText,
			MainPrice:                         data.MainPrice,
			MainPriceCurrencyCode:             data.MainPriceCurrencyCode,
			MainPriceCurrencyCodeText:         data.MainPriceCurrencyCodeText,
			MainPriceBaseQuantity:             data.MainPriceBaseQuantity,
			MainPriceBaseQuantityUnitCode:     data.MainPriceBaseQuantityUnitCode,
			MainPriceBaseQuantityUnitCodeText: data.MainPriceBaseQuantityUnitCodeText,
			MainDiscount:                      data.MainDiscount,
			NetPriceAmount:                    data.NetPriceAmount,
			NetPriceAmountCurrencyCode:        data.NetPriceAmountCurrencyCode,
			NetPriceAmountCurrencyCodeText:    data.NetPriceAmountCurrencyCodeText,
			NetPriceBaseQuantity:              data.NetPriceBaseQuantity,
			NetPriceBaseQuantityUnitCode:      data.NetPriceBaseQuantityUnitCode,
			NetPriceBaseQuantityUnitCodeText:  data.NetPriceBaseQuantityUnitCodeText,
			NetAmount:                         data.NetAmount,
			NetAmountCurrencyCode:             data.NetAmountCurrencyCode,
			NetAmountCurrencyCodeText:         data.NetAmountCurrencyCodeText,
			GrossAmount:                       data.GrossAmount,
			GrossAmountCurrencyCode:           data.GrossAmountCurrencyCode,
			GrossAmountCurrencyCodeText:       data.GrossAmountCurrencyCodeText,
			TaxAmount:                         data.TaxAmount,
			TaxAmountCurrencyCode:             data.TaxAmountCurrencyCode,
			TaxAmountCurrencyCodeText:         data.TaxAmountCurrencyCodeText,
			CashDiscountTermsCode:             data.CashDiscountTermsCode,
			CashDiscountTermsCodeText:         data.CashDiscountTermsCodeText,
			CancellationStatusCode:            data.CancellationStatusCode,
			CancellationStatusCodeText:        data.CancellationStatusCodeText,
			ConsistencyStatusCode:             data.ConsistencyStatusCode,
			ConsistencyStatusCodeText:         data.ConsistencyStatusCodeText,
			LifeCycleStatusCode:               data.LifeCycleStatusCode,
			LifeCycleStatusCodeText:           data.LifeCycleStatusCodeText,
			OrderingStatusCode:                data.OrderingStatusCode,
			OrderingStatusCodeText:            data.OrderingStatusCodeText,
			ConfirmationExistenceIndicator:    data.ConfirmationExistenceIndicator,
			CreationDateTime:                  data.CreationDateTime,
			LastChangeDateTime:                data.LastChangeDateTime,
			DeliveryPriorityCode:              data.DeliveryPriorityCode,
			DeliveryPriorityCodeText:          data.DeliveryPriorityCodeText,
			IncotermsClassificationCode:       data.IncotermsClassificationCode,
			IncotermsClassificationCodeText:   data.IncotermsClassificationCodeText,
			IncotermsTransferLocationName:     data.IncotermsTransferLocationName,
			GrossWeightMeasure:                data.GrossWeightMeasure,
			GrossWeightUnitCode:               data.GrossWeightUnitCode,
			GrossWeightUnitCodeText:           data.GrossWeightUnitCodeText,
			NetWeightMeasure:                  data.NetWeightMeasure,
			NetWeightUnitCode:                 data.NetWeightUnitCode,
			NetWeightUnitCodeText:             data.NetWeightUnitCodeText,
			VolumeMeasure:                     data.VolumeMeasure,
			VolumeUnitCode:                    data.VolumeUnitCode,
			VolumeUnitCodeText:                data.VolumeUnitCodeText,
			OrderedQuantity:                   data.OrderedQuantity,
			OrderedQuantityUnitCode:           data.OrderedQuantityUnitCode,
			OrderedQuantityUnitCodeText:       data.OrderedQuantityUnitCodeText,
			PricingProductID:                  data.PricingProductID,
			PlantPartyID:                      data.PlantPartyID,
			PlantPartyName:                    data.PlantPartyName,
			EntityLastChangedOn:               data.EntityLastChangedOn,
			ETag:                              data.ETag,
		})
	}

	return salesQuoteItemCollection, nil
}

func ConvertToSalesQuotePartyCollection(raw []byte, l *logger.Logger) ([]SalesQuotePartyCollection, error) {
	pm := &responses.SalesQuotePartyCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesQuotePartyCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	salesQuotePartyCollection := make([]SalesQuotePartyCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		salesQuotePartyCollection = append(salesQuotePartyCollection, SalesQuotePartyCollection{
			ObjectID:                                 data.ObjectID,
			ParentObjectID:                           data.ParentObjectID,
			QuoteID:                                  data.QuoteID,
			PartyID:                                  data.PartyID,
			RoleCode:                                 data.RoleCode,
			RoleCodeText:                             data.RoleCodeText,
			MainIndicator:                            data.MainIndicator,
			PartyContactPartyID:                      data.PartyContactPartyID,
			PartyName:                                data.PartyName,
			PartyContactPartyName:                    data.PartyContactPartyName,
			CountryCode:                              data.CountryCode,
			CountryCodeText:                          data.CountryCodeText,
			RegionCode:                               data.RegionCode,
			RegionCodeText:                           data.RegionCodeText,
			CareOfName:                               data.CareOfName,
			StreetPrefixName:                         data.StreetPrefixName,
			AdditionalStreetPrefixName:               data.AdditionalStreetPrefixName,
			HouseID:                                  data.HouseID,
			StreetName:                               data.StreetName,
			StreetSuffixName:                         data.StreetSuffixName,
			AdditionalStreetSuffixName:               data.AdditionalStreetSuffixName,
			DistrictName:                             data.DistrictName,
			CityName:                                 data.CityName,
			AdditionalCityName:                       data.AdditionalCityName,
			StreetPostalCode:                         data.StreetPostalCode,
			StreetPostalCodeText:                     data.StreetPostalCodeText,
			CountyName:                               data.CountyName,
			CompanyPostalCode:                        data.CompanyPostalCode,
			CompanyPostalCodeText:                    data.CompanyPostalCodeText,
			POBoxIndicator:                           data.POBoxIndicator,
			POBoxID:                                  data.POBoxID,
			POBoxPostalCode:                          data.POBoxPostalCode,
			POBoxPostalCodeText:                      data.POBoxPostalCodeText,
			POBoxDeviatingCountryCode:                data.POBoxDeviatingCountryCode,
			POBoxDeviatingCountryCodeText:            data.POBoxDeviatingCountryCodeText,
			POBoxDeviatingRegionCode:                 data.POBoxDeviatingRegionCode,
			POBoxDeviatingRegionCodeText:             data.POBoxDeviatingRegionCodeText,
			POBoxDeviatingCityName:                   data.POBoxDeviatingCityName,
			BuildingID:                               data.BuildingID,
			FloorID:                                  data.FloorID,
			RoomID:                                   data.RoomID,
			Phone:                                    data.Phone,
			Mobile:                                   data.Mobile,
			Fax:                                      data.Fax,
			Email:                                    data.Email,
			Web:                                      data.Web,
			CorrespondenceLanguageCode:               data.CorrespondenceLanguageCode,
			CorrespondenceLanguageCodeText:           data.CorrespondenceLanguageCodeText,
			PreferredCommunicationMediumTypeCode:     data.PreferredCommunicationMediumTypeCode,
			PreferredCommunicationMediumTypeCodeText: data.PreferredCommunicationMediumTypeCodeText,
			FirstLineName:                            data.FirstLineName,
			SecondLineName:                           data.SecondLineName,
			ThirdLineName:                            data.ThirdLineName,
			FourthLineName:                           data.FourthLineName,
			ETag:                                     data.ETag,
		})
	}

	return salesQuotePartyCollection, nil
}

func ConvertToSalesQuoteItemScheduleLineCollection(raw []byte, l *logger.Logger) ([]SalesQuoteItemScheduleLineCollection, error) {
	pm := &responses.SalesQuoteItemScheduleLineCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesQuoteItemScheduleLineCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	salesQuoteItemScheduleLineCollection := make([]SalesQuoteItemScheduleLineCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		salesQuoteItemScheduleLineCollection = append(salesQuoteItemScheduleLineCollection, SalesQuoteItemScheduleLineCollection{
			ObjectID:         data.ObjectID,
			ParentObjectID:   data.ParentObjectID,
			QuoteID:          data.QuoteID,
			QuoteItemID:      data.QuoteItemID,
			ID:               data.ID,
			TypeCode:         data.TypeCode,
			TypeCodeText:     data.TypeCodeText,
			Quantity:         data.Quantity,
			UnitCode:         data.UnitCode,
			UnitCodeText:     data.UnitCodeText,
			StartDateTime:    data.StartDateTime,
			EndDateTime:      data.EndDateTime,
			TimeZoneCode:     data.TimeZoneCode,
			TimeZoneCodeText: data.TimeZoneCodeText,
			ETag:             data.ETag,
		})
	}

	return salesQuoteItemScheduleLineCollection, nil
}