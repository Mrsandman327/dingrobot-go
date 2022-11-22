package dingrobot

import "testing"

var (
	webhook *Webhook
	token   = "fff686bfb4ba22d5915f66e6430b50495bd45169d8b5442a2f48090890a038f3"
	secret  = "SEC56de0cad9a8975f19af8b4ad112d2cb21abc38930ffa3bf84418c2a15cf8c4c0"
)

func init() {
	webhook = NewWebhook(token, secret)
}

func Test_SendText(t *testing.T) {
	msg := NewTextMsg().
		Content("我就是我, @XXX 是不一样的烟火").
		IsAtAll(false).
		AtMobile("180xxxxxx").
		AtUserID("user123")
	webhook.SendMessage(msg)
}

func Test_SendLink(t *testing.T) {
	msg := NewLinkMsg().
		Text("这个即将发布的新版本，创始人xx称它为红树林。而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是红树林").
		Title("时代的火车向前开").
		PicUrl("").
		MessageUrl(`https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&
	sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&
	isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI`)
	webhook.SendMessage(msg)
}

func Test_SendMarkDown(t *testing.T) {
	msg := NewMarkDownMsg().
		Title("杭州天气").
		Text(`#### 杭州天气 @150XXXXXXXX 
		> 9度，西北风1级，空气良89，相对温度73%
		> ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)
	    > ###### 10点20分发布 [天气](https://www.dingtalk.com)`).
		IsAtAll(true)

	webhook.SendMessage(msg)
}

func Test_SendFreeCard(t *testing.T) {
	msg := NewFeedCardMsg().AppendLink("时代的火车向前开1",
		"https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
		"https://www.dingtalk.com/").AppendLink("时代的火车向前开2",
		"https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
		"https://www.dingtalk.com/")

	webhook.SendMessage(msg)
}

func Test_SendActionCard(t *testing.T) {
	msg := NewActionCardMsg().
		Title("我 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身").
		Text(`![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)
#### 乔布斯 20 年前想打造的苹果咖啡厅\
		Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划`).
		BtnOrientation("0").
		AppendBtns("内容不错",
			"https://www.dingtalk.com/").
		AppendBtns("不感兴趣",
			"https://www.dingtalk.com/")

	webhook.SendMessage(msg)
}
