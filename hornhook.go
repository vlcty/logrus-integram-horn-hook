package hornhook

import (
    "fmt"
    "net/http"
    "strings"
    "encoding/json"

    "github.com/sirupsen/logrus"
)

type HornHook struct {
    levelsToRegister []logrus.Level
    WebhookID string
    Appname string
}

type HornHookText struct {
    Text string `json:"text"`
    Markdown bool `json:"mrkdwn"`
}

func New(webkookID string) *HornHook {
    return &HornHook{
        levelsToRegister: make([]logrus.Level, 0),
        WebhookID: webkookID,
    }
}

func (hook *HornHook) AddLevel(level logrus.Level) {
    hook.levelsToRegister = append(hook.levelsToRegister, level)
}

func (hook *HornHook) Fire(entry *logrus.Entry) error {
    hornText := HornHookText{
        Text: hook.formatEntry(entry),
        Markdown: false }

    builder := &strings.Builder{}

    encodererr := json.NewEncoder(builder).Encode(hornText)

    if encodererr != nil {
        return encodererr
    }

    _, posterr := http.Post("https://integram.org/webhook/" + hook.WebhookID,
        "application/json",
        strings.NewReader(builder.String()))

    if posterr != nil {
        return posterr
    }

    return nil
}

func (hook *HornHook) formatEntry(entry *logrus.Entry) string {
    value, _ := entry.String()

    if len(hook.Appname) == 0 {
        return fmt.Sprintf("%s", value)
    } else {
        return fmt.Sprintf("(%s) %s", hook.Appname, value)
    }
}

func (hook *HornHook) Levels() []logrus.Level {
	return hook.levelsToRegister
}
