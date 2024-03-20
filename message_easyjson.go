// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package botgolang

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson4086215fDecodeGithubComGldmtrBotGolang(in *jlexer.Lexer, out *Message) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ContentType":
			out.ContentType = MessageContentType(in.Uint8())
		case "msgId":
			out.ID = string(in.String())
		case "fileId":
			out.FileID = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "chat":
			(out.Chat).UnmarshalEasyJSON(in)
		case "replyMsgId":
			out.ReplyMsgID = string(in.String())
		case "forwardMsgId":
			out.ForwardMsgID = string(in.String())
		case "forwardChatId":
			out.ForwardChatID = string(in.String())
		case "timestamp":
			out.Timestamp = int(in.Int())
		case "inlineKeyboardMarkup":
			if in.IsNull() {
				in.Skip()
				out.InlineKeyboard = nil
			} else {
				if out.InlineKeyboard == nil {
					out.InlineKeyboard = new(Keyboard)
				}
				easyjson4086215fDecodeGithubComGldmtrBotGolang1(in, out.InlineKeyboard)
			}
		case "parseMode":
			out.ParseMode = ParseMode(in.String())
		case "requestID":
			out.RequestID = string(in.String())
		case "deeplink":
			out.Deeplink = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson4086215fEncodeGithubComGldmtrBotGolang(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ContentType\":"
		out.RawString(prefix[1:])
		out.Uint8(uint8(in.ContentType))
	}
	{
		const prefix string = ",\"msgId\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"fileId\":"
		out.RawString(prefix)
		out.String(string(in.FileID))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"chat\":"
		out.RawString(prefix)
		(in.Chat).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"replyMsgId\":"
		out.RawString(prefix)
		out.String(string(in.ReplyMsgID))
	}
	{
		const prefix string = ",\"forwardMsgId\":"
		out.RawString(prefix)
		out.String(string(in.ForwardMsgID))
	}
	{
		const prefix string = ",\"forwardChatId\":"
		out.RawString(prefix)
		out.String(string(in.ForwardChatID))
	}
	{
		const prefix string = ",\"timestamp\":"
		out.RawString(prefix)
		out.Int(int(in.Timestamp))
	}
	{
		const prefix string = ",\"inlineKeyboardMarkup\":"
		out.RawString(prefix)
		if in.InlineKeyboard == nil {
			out.RawString("null")
		} else {
			easyjson4086215fEncodeGithubComGldmtrBotGolang1(out, *in.InlineKeyboard)
		}
	}
	{
		const prefix string = ",\"parseMode\":"
		out.RawString(prefix)
		out.String(string(in.ParseMode))
	}
	{
		const prefix string = ",\"requestID\":"
		out.RawString(prefix)
		out.String(string(in.RequestID))
	}
	{
		const prefix string = ",\"deeplink\":"
		out.RawString(prefix)
		out.String(string(in.Deeplink))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4086215fEncodeGithubComGldmtrBotGolang(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4086215fEncodeGithubComGldmtrBotGolang(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4086215fDecodeGithubComGldmtrBotGolang(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4086215fDecodeGithubComGldmtrBotGolang(l, v)
}
func easyjson4086215fDecodeGithubComGldmtrBotGolang1(in *jlexer.Lexer, out *Keyboard) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Rows":
			if in.IsNull() {
				in.Skip()
				out.Rows = nil
			} else {
				in.Delim('[')
				if out.Rows == nil {
					if !in.IsDelim(']') {
						out.Rows = make([][]Button, 0, 2)
					} else {
						out.Rows = [][]Button{}
					}
				} else {
					out.Rows = (out.Rows)[:0]
				}
				for !in.IsDelim(']') {
					var v1 []Button
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						in.Delim('[')
						if v1 == nil {
							if !in.IsDelim(']') {
								v1 = make([]Button, 0, 1)
							} else {
								v1 = []Button{}
							}
						} else {
							v1 = (v1)[:0]
						}
						for !in.IsDelim(']') {
							var v2 Button
							(v2).UnmarshalEasyJSON(in)
							v1 = append(v1, v2)
							in.WantComma()
						}
						in.Delim(']')
					}
					out.Rows = append(out.Rows, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson4086215fEncodeGithubComGldmtrBotGolang1(out *jwriter.Writer, in Keyboard) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Rows\":"
		out.RawString(prefix[1:])
		if in.Rows == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Rows {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v5, v6 := range v4 {
						if v5 > 0 {
							out.RawByte(',')
						}
						(v6).MarshalEasyJSON(out)
					}
					out.RawByte(']')
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
