package main

import "fmt"

type EmailMessage struct {
	From               EmailRecipient         `json:"from"`
	To                 []EmailRecipient       `json:"to"`
	CC                 []EmailRecipient       `json:"cc,omitempty"`
	Bcc                []EmailRecipient       `json:"bcc,omitempty"`
	ReplyTo            *EmailRecipient        `json:"replyTo,omitempty"`
	Subject            *string                `json:"subject,omitempty"`
	TextPart           *string                `json:"textPart,omitempty"`
	HtmlPart           *string                `json:"htmlPart,omitempty"`
	TemplateId         int                    `json:"templateId,omitempty"`
	TemplateLanguage   bool                   `json:"templateLanguage,omitempty"`
	Variables          map[string]interface{} `json:"variables,omitempty"`
	Priority           int                    `json:"priority,omitempty"`
	Attachments        []EmailAttachment      `json:"attachments,omitempty"`
	InlinedAttachments []EmailAttachment      `json:"inlinedAttachments,omitempty"`
}

type EmailRecipient struct {
	Email string  `json:"email"`
	Name  *string `json:"name,omitempty"`
}

type EmailAttachment struct {
	Filename      string `json:"filename"`
	ContentType   string `json:"contentType"`
	Base64Content string `json:"base64Content"`
}

type ContactRecipient struct {
	Address string  `json:"address"`
	Name    *string `json:"name,omitempty"`
}

type SendMailMultipleEmail struct {
	Contacts []ContactRecipient `json:"contacts"`
	Data     EmailMessage       `json:"data"`
}

type CampaignEmailMessage struct {
	From               EmailRecipient         `json:"from"`
	Contacts           []EmailRecipient       `json:"contacts"`
	ReplyTo            *EmailRecipient        `json:"replyTo,omitempty"`
	Subject            *string                `json:"subject,omitempty"`
	TextPart           *string                `json:"textPart,omitempty"`
	HtmlPart           *string                `json:"htmlPart,omitempty"`
	TemplateId         int                    `json:"templateId,omitempty"`
	TemplateLanguage   bool                   `json:"templateLanguage,omitempty"`
	Variables          map[string]interface{} `json:"variables,omitempty"`
	Priority           int                    `json:"priority,omitempty"`
	Attachments        []EmailAttachment      `json:"attachments,omitempty"`
	InlinedAttachments []EmailAttachment      `json:"inlinedAttachments,omitempty"`
}

func SendEmail(apiKey string, payload EmailMessage) {
	reader, err := convertToReader(payload)

	if err != nil {
		fmt.Println("Error creating reader:", err)
		return
	}

	HttpClient(MethodPost, fmt.Sprintf("%v/connect/v1.0/email/send", BaseURL), apiKey, nil, reader)
}

func SendMultipleMail(apiKey string, payload SendMailMultipleEmail) {
	reader, err := convertToReader(payload)

	if err != nil {
		fmt.Println("Error creating reader:", err)
		return
	}

	HttpClient(MethodPost, fmt.Sprintf("%v/connect/v1.0/email/send-multiple", BaseURL), apiKey, nil, reader)
}

func SendCampaignMail(apiKey string, payload CampaignEmailMessage) {
	reader, err := convertToReader(payload)

	if err != nil {
		fmt.Println("Error creating reader:", err)
		return
	}

	HttpClient(MethodPost, fmt.Sprintf("%v/connect/v1.0/email/send-campaign", BaseURL), apiKey, nil, reader)
}
