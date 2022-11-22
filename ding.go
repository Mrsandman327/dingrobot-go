package dingrobot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type dingMsg interface {
}

type dingMsgType struct {
	Msgtype string `json:"msgtype"`
}

type dingAt struct {
	AtMobiles []string  `json:"atMobiles"`
	AtUserIds []string  `json:"atUserIds"`
	IsAtAll   bool  `json:"isAtAll"`
}

func (t *dingAt) AtMobile(at ...string) *dingAt {
	t.AtMobiles = append(t.AtMobiles, at...)
	return t
}

func (t *dingAt) AtUserID(at ...string) *dingAt {
	t.AtUserIds = append(t.AtUserIds, at...)
	return t
}

func (t *dingAt) IsAtAllUser(b bool) *dingAt {
	t.IsAtAll = b
	return t
}

type Webhook struct {
	AccessToken string
	Secret      string
	Error       error
}

func NewWebhook(token string, secret string) *Webhook {
	return &Webhook{AccessToken: token, Secret: secret}
}

func (t *Webhook) SendMessage(msg dingMsg) {
	b, err := json.Marshal(msg)
	if err != nil {
		t.Error = err
	}
	resp, err := http.Post(t.getURL(), "application/json", bytes.NewBuffer(b))
	if err != nil {
		t.Error = err
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Error = err
	}
}

func (t *Webhook) hmacSha256(stringToSign string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (t *Webhook) getURL() string {
	wh := "https://oapi.dingtalk.com/robot/send?access_token=" + t.AccessToken
	timestamp := time.Now().UnixNano() / 1e6
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, t.Secret)
	sign := t.hmacSha256(stringToSign, t.Secret)
	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", wh, timestamp, sign)
	return url
}
