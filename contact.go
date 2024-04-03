package main

import "fmt"

type ContactCreateRequestBody struct {
	Contacts []ContactCreateAttribute `json:"contacts"`
}

type ContactCreateAttribute struct {
	Email        string  `json:"email"`
	FirstName    *string `json:"firstName,omitempty"`
	LastName     *string `json:"lastName,omitempty"`
	PhoneCountry *string `json:"phoneCountry,omitempty"`
	PhoneNumber  *string `json:"phoneNumber,omitempty"`
}

type ContactResponseData struct {
	Contact Contact `json:"contact"`
}

type ContactResponse struct {
	Code       int                 `json:"code"`
	Data       ContactResponseData `json:"data"`
	Message    string              `json:"message"`
	DevMessage string              `json:"devMessage"`
	Success    bool                `json:"success"`
}

type Contact struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"userId"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func CreateContact(apiKey string, payload ContactCreateAttribute) {
	reader, err := convertToReader(payload)

	if err != nil {
		fmt.Println("Error creating reader:", err)
		return
	}

	HttpClient(MethodPost, fmt.Sprintf("%v/connect/v1.0/contacts", BaseURL), apiKey, nil, reader)
}

func UpdateContact(apiKey string, contactID int, payload ContactCreateAttribute) {
	reader, err := convertToReader(payload)

	if err != nil {
		fmt.Println("Error creating reader:", err)
		return
	}

	HttpClient(MethodPatch, fmt.Sprintf("%v/connect/v1.0/contacts/%v", BaseURL, contactID), apiKey, nil, reader)
}
