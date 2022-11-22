package dingrobot

type dingBtns struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type actionCard struct {
	Title          string      `json:"title"`
	Text           string      `json:"text"`
	BtnOrientation string      `json:"btnOrientation"`
	Btns           []*dingBtns `json:"btns"`
}

type dingActionCardMsg struct {
	dingMsgType
	ActionCard *actionCard `json:"actionCard"`
}

func NewActionCardMsg() *dingActionCardMsg {
	return &dingActionCardMsg{
		dingMsgType{Msgtype: "actionCard"},
		&actionCard{Btns: make([]*dingBtns, 0)},
	}
}

func (t *dingActionCardMsg) Title(s string) *dingActionCardMsg {
	t.ActionCard.Title = s
	return t
}

func (t *dingActionCardMsg) Text(s string) *dingActionCardMsg {
	t.ActionCard.Text = s
	return t
}

func (t *dingActionCardMsg) BtnOrientation(s string) *dingActionCardMsg {
	t.ActionCard.BtnOrientation = s
	return t
}

func (t *dingActionCardMsg) AppendBtns(title string, actionurl string) *dingActionCardMsg {
	btns := &dingBtns{
		Title:     title,
		ActionURL: actionurl,
	}
	t.ActionCard.Btns = append(t.ActionCard.Btns, btns)
	return t
}
