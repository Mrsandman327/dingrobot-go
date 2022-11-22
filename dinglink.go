package dingrobot

type dingLink struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

type dingLinkMsg struct {
	dingMsgType
	Link *dingLink `json:"link"`
}

func NewLinkMsg() *dingLinkMsg {
	return &dingLinkMsg{
		dingMsgType{Msgtype: "link"},
		&dingLink{},
	}
}

func (t *dingLinkMsg) Text(s string) *dingLinkMsg {
	t.Link.Text = s
	return t
}

func (t *dingLinkMsg) Title(s string) *dingLinkMsg {
	t.Link.Title = s
	return t
}

func (t *dingLinkMsg) PicUrl(s string) *dingLinkMsg {
	t.Link.PicUrl = s
	return t
}

func (t *dingLinkMsg) MessageUrl(s string) *dingLinkMsg {
	t.Link.MessageUrl = s
	return t
}
