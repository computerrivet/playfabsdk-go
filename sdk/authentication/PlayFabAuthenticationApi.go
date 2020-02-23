package authentication

// This code was generated by a tool. Any changes may be overwritten

import (
	"encoding/json"

	playfab "github.com/computerrivet/playfabsdk-go/sdk"

	"github.com/mitchellh/mapstructure"
)

// GetEntityToken method to exchange a legacy AuthenticationTicket or title SecretKey for an Entity Token or to refresh a still valid
// Entity Token.
// https://api.playfab.com/Documentation/Authentication/method/GetEntityToken
func GetEntityToken(settings *playfab.Settings, postData *GetEntityTokenRequestModel, entityToken string, clientSessionTicket string, developerSecretKey string) (*GetEntityTokenResponseModel, error) {
	var authKey, authValue string
	if entityToken != "" {
		authKey = "X-EntityToken"
		authValue = entityToken
	} else if clientSessionTicket != "" {
		authKey = "X-Authentication"
		authValue = clientSessionTicket
	} else if developerSecretKey != "" {
		authKey = "X-SecretKey"
		authValue = developerSecretKey
	}

	b, errMarshal := json.Marshal(postData)
	if errMarshal != nil {
		return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
	}

	sourceMap, err := playfab.Request(settings, b, "/Authentication/GetEntityToken", authKey, authValue)
	if err != nil {
		return nil, err
	}

	result := &GetEntityTokenResponseModel{}

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

// GetEntityToken method to exchange a legacy AuthenticationTicket or title SecretKey for an Entity Token or to refresh a still valid
// Entity Token.
// https://api.playfab.com/Documentation/Authentication/method/GetEntityToken
func GetEntityTokenWithEntityToken(settings *playfab.Settings, postData *GetEntityTokenRequestModel, entityToken string) (*GetEntityTokenResponseModel, error) {
	if entityToken == "" {
		return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
	}
	b, err := json.Marshal(postData)
	if err != nil {
		return nil, err
	}
	sourceMap, err := playfab.Request(settings, b, "/Authentication/GetEntityToken", "X-EntityToken", entityToken)
	if err != nil {
		return nil, err
	}

	result := &GetEntityTokenResponseModel{}

	config := mapstructure.DecoderConfig{
		DecodeHook: playfab.StringToDateTimeHook,
		Result:     result,
	}

	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(sourceMap)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetEntityToken method to exchange a legacy AuthenticationTicket or title SecretKey for an Entity Token or to refresh a still valid
// Entity Token.
// https://api.playfab.com/Documentation/Authentication/method/GetEntityToken
func GetEntityTokenWithSessionTicket(settings *playfab.Settings, postData *GetEntityTokenRequestModel, clientSessionTicket string) (*GetEntityTokenResponseModel, error) {
	if clientSessionTicket == "" {
		return nil, playfab.NewCustomError("clientSessionTicket should not be an empty string", playfab.ErrorGeneric)
	}
	b, err := json.Marshal(postData)
	if err != nil {
		return nil, err
	}
	sourceMap, err := playfab.Request(settings, b, "/Authentication/GetEntityToken", "X-Authentication", clientSessionTicket)
	if err != nil {
		return nil, err
	}

	result := &GetEntityTokenResponseModel{}

	config := mapstructure.DecoderConfig{
		DecodeHook: playfab.StringToDateTimeHook,
		Result:     result,
	}

	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(sourceMap)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetEntityToken method to exchange a legacy AuthenticationTicket or title SecretKey for an Entity Token or to refresh a still valid
// Entity Token.
// https://api.playfab.com/Documentation/Authentication/method/GetEntityToken
func GetEntityTokenWithSecretKey(settings *playfab.Settings, postData *GetEntityTokenRequestModel, developerSecretKey string) (*GetEntityTokenResponseModel, error) {
	if developerSecretKey == "" {
		return nil, playfab.NewCustomError("developerSecretKey should not be an empty string", playfab.ErrorGeneric)
	}
	b, err := json.Marshal(postData)
	if err != nil {
		return nil, err
	}
	sourceMap, err := playfab.Request(settings, b, "/Authentication/GetEntityToken", "X-SecretKey", developerSecretKey)
	if err != nil {
		return nil, err
	}

	result := &GetEntityTokenResponseModel{}

	config := mapstructure.DecoderConfig{
		DecodeHook: playfab.StringToDateTimeHook,
		Result:     result,
	}

	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(sourceMap)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// ValidateEntityToken method for a server to validate a client provided EntityToken. Only callable by the title entity.
// https://api.playfab.com/Documentation/Authentication/method/ValidateEntityToken
func ValidateEntityToken(settings *playfab.Settings, postData *ValidateEntityTokenRequestModel, entityToken string) (*ValidateEntityTokenResponseModel, error) {
	if entityToken == "" {
		return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
	}
	b, errMarshal := json.Marshal(postData)
	if errMarshal != nil {
		return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
	}

	sourceMap, err := playfab.Request(settings, b, "/Authentication/ValidateEntityToken", "X-EntityToken", entityToken)
	if err != nil {
		return nil, err
	}

	result := &ValidateEntityTokenResponseModel{}

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
