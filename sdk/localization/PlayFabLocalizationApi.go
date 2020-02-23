package localization

// This code was generated by a tool. Any changes may be overwritten

import (
	"encoding/json"

	playfab "github.com/computerrivet/playfabsdk-go/sdk"

	"github.com/mitchellh/mapstructure"
)

// GetLanguageList retrieves the list of allowed languages, only accessible by title entities
// https://api.playfab.com/Documentation/Localization/method/GetLanguageList
func GetLanguageList(settings *playfab.Settings, postData *GetLanguageListRequestModel, entityToken string) (*GetLanguageListResponseModel, error) {
	if entityToken == "" {
		return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
	}
	b, errMarshal := json.Marshal(postData)
	if errMarshal != nil {
		return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
	}

	sourceMap, err := playfab.Request(settings, b, "/Locale/GetLanguageList", "X-EntityToken", entityToken)
	if err != nil {
		return nil, err
	}

	result := &GetLanguageListResponseModel{}

	config := mapstructure.DecoderConfig{
		DecodeHook: playfab.StringToDateTimeHook,
		Result:     result,
	}

	decoder, errDecoding := mapstructure.NewDecoder(&config)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	errDecoding = decoder.Decode(sourceMap)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	return result, nil
}
