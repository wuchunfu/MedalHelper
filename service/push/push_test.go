package push

import (
	"testing"

	"medalhelper/util"
)

func TestPushDeerPush(t *testing.T) {
	push := PushDeerPush{
		util.Endpoint{
			Name:  "push",
			Type:  "push_deer",
			URL:   "http://api2.pushdeer.com/message/push",
			Token: "<YOUR-TOKEN-HERE>",
		},
	}
	if err := push.Submit(Data{
		Title:   "# Just for test",
		Content: "Good Morning!",
	}); err != nil {
		t.Fatalf("Push error: %v", err)
	}
}

func TestPushPlusPush(t *testing.T) {
	push := PushDeerPush{
		util.Endpoint{
			Name:  "push",
			Type:  "push_plus",
			URL:   "http://www.pushplus.plus/send",
			Token: "<YOUR-TOKEN-HERE>",
		},
	}
	if err := push.Submit(Data{
		Title:   "# Just for test",
		Content: "Good Morning!",
	}); err != nil {
		t.Fatalf("Push error: %v", err)
	}
}

func TestTelegramPush(t *testing.T) {
	push := TelegramPush{
		util.Endpoint{
			Name:  "push",
			Type:  "telegram",
			URL:   "https://api.telegram.org/bot<YOUR-BOT-TOKEN-HERE>/sendMessage",
			Token: "<YOUR-TELEGRAM-CHATID>",
		},
	}
	if err := push.Submit(Data{
		Title:   "# Just for test",
		Content: "Good Morning!",
	}); err != nil {
		t.Fatalf("Push error: %v", err)
	}
}
