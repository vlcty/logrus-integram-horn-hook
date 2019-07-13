# logrus-integram-horn-hook

Logrus hook to send log messages to [integram](https://integram.org/) using their webhook service.   
All you need to do is start a chat with [https://t.me/bullhorn_bot](@bullhorn_bot). The bot will greet you with a message containing your webhook ID. For example:

```
https://integram.org/webhook/supersecret123
```

In this example the webhook ID is `supersecret123`. Save it because you need it later!

## Example

Simple example to send a telegram message if an error is logged:

```go
package main

import (
    "github.com/sirupsen/logrus"
    hornhook "github.com/vlcty/logrus-integram-horn-hook"
)

func main() {
    hook := hornhook.New("supersecret123s") // Set your webhook ID here
    // hook.Appname = "Test"
    hook.AddLevel(logrus.ErrorLevel)

    logrus.AddHook(hook)

    logrus.Error("test123")
}
```

In the example you will receive a Telegram message like this:

```
[ERROR] test123
```

If you specify the app name the message changes slightly:

```
(Test) [ERROR] test123
```

## Levels

By default the hook doesn't register any log levels. You have to add them yourself before calling `logrus.AddHook`!
