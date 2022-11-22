package dingrobot

type dingFeedLink struct {
	Title      string `json:"title"`
	PicUrl     string `json:"picURL"`
	MessageUrl string `json:"messageURL"`
}

type dingFeedCard struct {
	Links []*dingFeedLink `json:"links"`
}

type dingFeedCardMsg struct {
	dingMsgType
	FeedCard *dingFeedCard `json:"feedCard"`
}

func NewFeedCardMsg() *dingFeedCardMsg {
	return &dingFeedCardMsg{
		dingMsgType{Msgtype: "feedCard"},
		&dingFeedCard{Links: make([]*dingFeedLink, 0)},
	}
}

func (t *dingFeedCardMsg) AppendLink(title string, picurl string, msgurl string) *dingFeedCardMsg {
	link := &dingFeedLink{
		Title:      title,
		PicUrl:     picurl,
		MessageUrl: msgurl,
	}
	t.FeedCard.Links = append(t.FeedCard.Links, link)
	return t
}
