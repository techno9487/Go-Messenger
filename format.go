package main

type WebhookEvent struct {
	Object string
	Entry []Event
}

type Event struct {
	Id string
	Time int
	Messaging []Messaging
}

type Messaging struct {
	Sender IdStruct
	Recipient IdStruct `json:"recipient"`
	Message Message `json:"message"`
}

type IdStruct struct {
	Id string `json:"id"`
}

type Message struct {
	Mid string
	Text string `json:"text"`
	Attachments []MessageAttachment
	QuickReply interface{} `json:"quick_reply"`
}

type MessageImage struct {
	Attachment MessageAttachment `json:"attachment"`
}

type MessageAttachment struct {
	Type string `json:"type"`
	Payload interface{} `json:"payload"`
}

type MessageSend struct {
	Text string `json:"text"`
}

type SendMessage struct {
	Recipient IdStruct `json:"recipient"`
	Message MessageSend `json:"message"`
}

type FacebookError struct {
	Message string `json:"message"`
	Type string `json:"type"`
	Code int `json:"code"`
	Subcode int `json:"error_subcode"`
	Title string `json:"error_user_title"`
	UserMessage string `json:"error_user_message"`
	Trace string `json:"fbtrace_id"`
}

type FacebookErrorResponse struct {
	Error FacebookError `json:"error"`
}