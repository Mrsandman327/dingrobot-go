package dingrobot

type dingText struct {
	Content string `json:"content"`
}

type dingTextMsg struct {
	dingMsgType
	At   *dingAt   `json:"at"`
	Text *dingText `json:"text"`
}

func NewTextMsg() *dingTextMsg {
	return &dingTextMsg{
		dingMsgType{Msgtype: "text"},
		&dingAt{},
		&dingText{},
	}
}

func (t *dingTextMsg) Content(s string) *dingTextMsg {
	t.Text.Content = s
	return t
}

func (t *dingTextMsg) AtMobile(at ...string) *dingTextMsg {
	t.At.AtMobile(at...)
	return t
}

func (t *dingTextMsg) AtUserID(at ...string) *dingTextMsg {
	t.At.AtUserID(at...)
	return t
}

func (t *dingTextMsg) IsAtAll(b bool) *dingTextMsg {
	t.At.IsAtAllUser(b)
	return t
}
