package main

type Seamailer struct {
	apiKey string
}

const BaseURL = "https://api.seamailer.app"

func NewSeamailer(apiKey string) Seamailer {
	return Seamailer{
		apiKey: apiKey,
	}
}

func (seamailer Seamailer) CreateContact(payload ContactCreateAttribute) {
	CreateContact(seamailer.apiKey, payload)
}

func (seamailer Seamailer) UpdateContact(contactID int, payload ContactCreateAttribute) {
	UpdateContact(seamailer.apiKey, contactID, payload)
}
