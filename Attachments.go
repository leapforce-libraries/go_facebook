package facebook

type Attachments struct {
	Data []Attachment `json:"data"`
}

type Attachment struct {
	Description string               `json:"description"`
	Media       AttachmentMediaImage `json:"media"`
	Target      AttachmentTarget     `json:"target"`
	Title       string               `json:"title"`
	Type        string               `json:"type"`
	URL         string               `json:"urlV16"`
}

type AttachmentMedia struct {
	Image AttachmentMediaImage `json:"image"`
}

type AttachmentMediaImage struct {
	Height int64  `json:"height"`
	Source string `json:"src"`
	Width  string `json:"width"`
}

type AttachmentTarget struct {
	URL string `json:"urlV16"`
}
