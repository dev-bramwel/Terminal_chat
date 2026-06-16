package chat

import (
	"fmt"
	"time"
)

type Message struct {
	Sender   string
	Text     string
	Time     time.Time
	IsSystem bool
}

func NewChatMessage(sender string, text string) Message {
	return Message{
		Sender: sender,
		Text:   text,
		Time:   time.Now(),
	}
}

func NewSystemMessage(text string) Message {
	return Message{
		Sender:   "System",
		Text:     text,
		Time:     time.Now(),
		IsSystem: true,
	}
}

func (m Message) Format() string {
	if m.IsSystem {
		return formatWithTime(m.Time, "System", m.Text)
	}
	return formatWithTime(m.Time, m.Sender, m.Text)
}

func FormatMessage(sender string, msg string) string {
	return NewChatMessage(sender, msg).Format()
}

func FormatSystemMessage(msg string) string {
	return NewSystemMessage(msg).Format()
}

func Prompt(sender string) string {
	return fmt.Sprintf("[%s][%s]: ", time.Now().Format("2006-01-02 15:04:05"), sender)
}

func formatWithTime(t time.Time, sender string, text string) string {
	return fmt.Sprintf("[%s][%s]: %s\n", t.Format("2006-01-02 15:04:05"), sender, text)
}
