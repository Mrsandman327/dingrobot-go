package dingrobot

type dingMarkDown struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

type dingMarkDownMsg struct {
	dingMsgType
	At       *dingAt       `json:"at"`
	MarkDown *dingMarkDown `json:"markdown"`
}

func NewMarkDownMsg() *dingMarkDownMsg {
	return &dingMarkDownMsg{
		dingMsgType{Msgtype: "markdown"},
		&dingAt{},
		&dingMarkDown{},
	}
}

func (t *dingMarkDownMsg) Text(s string) *dingMarkDownMsg {
	t.MarkDown.Text = s
	return t
}

func (t *dingMarkDownMsg) Title(s string) *dingMarkDownMsg {
	t.MarkDown.Title = s
	return t
}

func (t *dingMarkDownMsg) AtMobile(at ...string) *dingMarkDownMsg {
	t.At.AtMobile(at...)
	return t
}

func (t *dingMarkDownMsg) AtUserID(at ...string) *dingMarkDownMsg {
	t.At.AtUserID(at...)
	return t
}

func (t *dingMarkDownMsg) IsAtAll(b bool) *dingMarkDownMsg {
	t.At.IsAtAllUser(b)
	return t
}
