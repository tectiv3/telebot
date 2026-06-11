package telebot

import "encoding/json"

// InputRichMessage describes a rich message to be sent.
// Exactly one of HTML or Markdown must be set.
type InputRichMessage struct {
	HTML                string `json:"html,omitempty"`
	Markdown            string `json:"markdown,omitempty"`
	IsRTL               bool   `json:"is_rtl,omitempty"`
	SkipEntityDetection bool   `json:"skip_entity_detection,omitempty"`
}

// RichMessage represents a received rich formatted message.
type RichMessage struct {
	Blocks json.RawMessage `json:"blocks"`
	IsRTL  bool            `json:"is_rtl,omitempty"`
}
