package main

import (
	"fmt"
	"strconv"
	"strings"
)

type EmailTemplate struct {
	UserID            uint64 `json:"userId"`     // Use uint64 for user ID (likely non-negative)
	TemplateID        uint64 `json:"templateId"` // Consider a unique identifier type later
	TemplateName      string `json:"templateName"`
	SenderName        string `json:"senderName"`
	FromEmail         string `json:"fromEmail"` // TODO: Implement linking to user/project emails
	Subject           string `json:"subject"`
	Type              string `json:"type"`
	EmailHtml         string `json:"emailHtml"`
	EmailDesign       string `json:"emailDesign"`
	IsEditable        bool   `json:"isEditable"`
	IsPublished       bool   `json:"isPublished"`
	IsTemplate        bool   `json:"isTemplate"`
	IsGalleryTemplate bool   `json:"isGalleryTemplate"`
	Tags              string `json:"tags"`
}

type EmailTemplatesResponse struct {
	Templates []EmailTemplate `json:"templates"`
}

type EmailTemplateResponse struct {
	Template EmailTemplate `json:"template"`
}

func GetPublicTemplates(apiKey string, templateIds []uint) {
	// Create a string slice to hold string representations of each uint
	strSlice := make([]string, len(templateIds))
	for i, id := range templateIds {
		strSlice[i] = strconv.Itoa(int(id)) // Convert uint to int (assuming uint is non-negative)
	}

	HttpClient(MethodGet, fmt.Sprintf("%v/connect/v1.0/templates?templateIds=%v", BaseURL, strings.Join(strSlice, ";")), apiKey, nil, nil)
}

func GetOnePublicTemplate(apiKey string, templateId int) {
	HttpClient(MethodGet, fmt.Sprintf("%v/connect/v1.0/templates/%v", BaseURL, templateId), apiKey, nil, nil)
}

func PreviewPublicTemplate(apiKey string, templateId int, variables map[string]interface{}) {
	data := map[string]interface{}{
		"templateId": templateId,
		"variables":  variables,
	}

	reader, err := convertToReader(data)

	if err != nil {
		fmt.Println("Error creating reader:", err)
		return
	}

	HttpClient(MethodPost, fmt.Sprintf("%v/connect/v1.0/templates/preview", BaseURL), apiKey, nil, reader)
}
