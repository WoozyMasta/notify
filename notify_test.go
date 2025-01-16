package notify

import (
	"os"
	"testing"
	"time"
)

// TestSendEditDeleteDiscord tests the Send, Edit, and Delete functions for Discord.
func TestSendEditDeleteDiscord(t *testing.T) {
	discordID, ok := os.LookupEnv("DISCORD_ID")
	if !ok {
		t.Error("Discord webhook ID must be set in the environment variable 'DISCORD_ID'")
	}
	discordToken, ok := os.LookupEnv("DISCORD_TOKEN")
	if !ok {
		t.Error("Discord webhook token must be set in the environment variable 'DISCORD_TOKEN'")
	}

	msg := "This is a **bold**, _italic_, [inline URL](http://www.example.com/), `code`."
	discordClient := NewDiscord(discordID, discordToken)

	// Send messages
	id, err := discordClient.Send(msg)
	if err != nil {
		t.Errorf("Failed to send Discord message: %v", err)
	}

	// Wait for messages to be sent
	time.Sleep(3 * time.Second)

	// Edit messages
	if err := discordClient.Edit(id, "Redacted Discord Message"); err != nil {
		t.Errorf("Failed to edit Discord message: %v", err)
	}

	// Wait for messages to be edited
	time.Sleep(3 * time.Second)

	// Delete messages
	if err := discordClient.Delete(id); err != nil {
		t.Errorf("Failed to delete Discord message: %v", err)
	}
}

// TestSendEditDeleteTelegram tests the Send, Edit, and Delete functions for Telegram.
func TestSendEditDeleteTelegram(t *testing.T) {
	telegramID, ok := os.LookupEnv("TELEGRAM_ID")
	if !ok {
		t.Error("Telegram chat ID must be set in the environment variable 'TELEGRAM_ID'")
	}
	telegramToken, ok := os.LookupEnv("TELEGRAM_TOKEN")
	if !ok {
		t.Error("Telegram bot token must be set in the environment variable 'TELEGRAM_TOKEN'")
	}

	msg := "This is a **bold**, _italic_, [inline URL](http://www.example.com/), `code`."
	telegramClient := NewTelegram(telegramID, telegramToken)

	// Send messages
	id, err := telegramClient.Send(msg)
	if err != nil {
		t.Errorf("Failed to send Telegram message: %v", err)
	}

	// Wait for messages to be sent
	time.Sleep(3 * time.Second)

	// Edit messages
	if err := telegramClient.Edit(id, "Redacted Telegram Message"); err != nil {
		t.Errorf("Failed to edit Telegram message: %v", err)
	}

	// Wait for messages to be edited
	time.Sleep(3 * time.Second)

	// Delete messages
	if err := telegramClient.Delete(id); err != nil {
		t.Errorf("Failed to delete Telegram message: %v", err)
	}
}
