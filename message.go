package goicqbot

import "fmt"

//go:generate easyjson -all message.go

// Message represents a text message in ICQ
type Message struct {
	client *Client

	// Id of the message (for editing)
	ID string `json:"msgId"`

	// Id of file to send
	FileID string `json:"fileId"`

	// Text of the message or caption for file
	Text string `json:"text"`

	// Chat where to send the message
	Chat Chat `json:"chat"`

	// Id of replied message
	// You can't use it with ForwardMsgID or ForwardChatID
	ReplyMsgID string `json:"replyMsgId"`

	// Id of forwarded message
	// You can't use it with ReplyMsgID
	ForwardMsgID string `json:"forwardMsgId"`

	// Id of a chat from which you forward the message
	// You can't use it with ReplyMsgID
	// You should use it with ForwardMsgID
	ForwardChatID string `json:"replyChatId"`
}

// Send method sends your message.
// Make sure you have Text or FileID in your message.
func (m *Message) Send() error {
	if (m.client == nil) {
		return fmt.Errorf("client is not inited, create message with constructor NewMessage, NewTextMessage, etc")
	}

	if m.Chat.ID == "" {
		return fmt.Errorf("message should have chat id")
	}

	if m.Text == "" && m.FileID == "" {
		return fmt.Errorf("cannot send message or file without data")
	}

	return m.client.SendMessage(m)
}

// Edit method edits your message.
// Make sure you have ID in your message.
func (m *Message) Edit() error {
	if m.ID == "" {
		return fmt.Errorf("cannot edit message without id")
	}
	return m.client.EditMessage(m)
}

// Delete method deletes your message.
// Make sure you have ID in your message.
func (m *Message) Delete() error {
	if m.ID == "" {
		return fmt.Errorf("cannot delete message without id")
	}

	return m.client.DeleteMessage(m)
}

// Reply method replies to the message.
// Make sure you have ID in the message.
func (m *Message) Reply(text string) error {
	if m.ID == "" {
		return fmt.Errorf("cannot reply to message without id")
	}

	m.ReplyMsgID = m.ID
	m.Text = text

	return m.client.SendMessage(m)
}

// Forward method forwards your message to chat.
// Make sure you have ID in your message.
func (m *Message) Forward(chatID string) error {
	if m.ID == "" {
		return fmt.Errorf("cannot forward message without id")
	}

	m.ForwardChatID = m.Chat.ID
	m.ForwardMsgID = m.ID
	m.Chat.ID = chatID

	return m.client.SendMessage(m)
}