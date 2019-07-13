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
    hornText := HornHookText{ Text: hook.formatEntry(entry) }
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
    if len(hook.Appname) == 0 {
        return fmt.Sprintf("\\[%s] %s", strings.ToUpper(entry.Level.String()), entry.Message)
    } else {
        return fmt.Sprintf("(%s) \\[%s] %s", hook.Appname, strings.ToUpper(entry.Level.String()),
            entry.Message)
    }
}

func (hook *HornHook) Levels() []logrus.Level {
	return hook.levelsToRegister
}
