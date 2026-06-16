package chat

// "fmt" - For formatting text (like "[2024-01-20][Name]: Hello, world!")
// "time" - For getting the current date and time (when messages were sent)
import (
	"fmt"  // Combines text and variables into formatted strings
	"time" // Gets current time and formats it nicely
)

func FormatMessage(sender string, msg string) string {
	
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	
	return fmt.Sprintf("[%s][%s]: %s\n", timestamp, sender, msg)
}

func FormatSystemMessage(msg string) string {
	// Same timestamp logic as FormatMessage.
	// Every message (system or user) gets the same timestamp format.
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	
	
	return fmt.Sprintf("[%s][System]: %s\n", timestamp, msg)
}