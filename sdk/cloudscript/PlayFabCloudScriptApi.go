package cloudscript

// This code was generated by a tool. Any changes may be overwritten

import (
	"encoding/json"

	playfab "github.com/dgkanatsios/playfabsdk-go/sdk"

	"github.com/mitchellh/mapstructure"
)

// ExecuteEntityCloudScript cloud Script is one of PlayFab's most versatile features. It allows client code to request execution of any kind of
// custom server-side functionality you can implement, and it can be used in conjunction with virtually anything.
// https://api.playfab.com/Documentation/CloudScript/method/ExecuteEntityCloudScript
func ExecuteEntityCloudScript(settings *playfab.Settings, postData *ExecuteEntityCloudScriptRequestModel, entityToken string) (*ExecuteCloudScriptResultModel, error) {
	if entityToken == "" {
		return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
	}
	b, errMarshal := json.Marshal(postData)
	if errMarshal != nil {
		return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
	}

	sourceMap, err := playfab.Request(settings, b, "/CloudScript/ExecuteEntityCloudScript", "X-EntityToken", entityToken)
	if err != nil {
		return nil, err
	}

	result := &ExecuteCloudScriptResultModel{}

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
