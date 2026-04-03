package client

type ImgAttachment struct {
	Type   string `json:"type"`
	Src    string `json:"src"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type AudioAttachment struct {
	Type   string `json:"type"`
	Src    string `json:"src"`
	Origin string `json:"origin"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Genre  string `json:"genre"`
}

type Attachment struct {
	Type   string `json:"type"`
	Src    string `json:"src"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
	Origin string `json:"origin,omitempty"`
	Artist string `json:"artist,omitempty"`
	Title  string `json:"title,omitempty"`
	Genre  string `json:"genre,omitempty"`
}
