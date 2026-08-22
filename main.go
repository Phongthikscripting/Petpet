package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"
	"petpet/commands"
	subcommands "petpet/commands/subcommands"

	"codeberg.org/lumap/chihuahua"
	godotenv "github.com/joho/godotenv"
)

func main() {
	chihuahua.LogInfo("Loading environmental variables...")
	if err := godotenv.Load(".env"); err != nil {
		chihuahua.LogError("failed to load env variables", err)
	}

	chihuahua.LogInfo("Creating new client...")
	bot := chihuahua.CreateBot(os.Getenv("DISCORD_BOT_TOKEN"), os.Getenv("DISCORD_PUBLIC_KEY"))

	chihuahua.LogInfo("Registering commands & static components...")
	bot.RegisterCommand(commands.Meow)
	bot.RegisterCommand(commands.Donate)
	bot.RegisterCommand(commands.Support)

	bot.RegisterCommand(commands.Petpet)
	bot.RegisterSubCommand(subcommands.PetpetUser, "petpet")
	bot.RegisterSubCommand(subcommands.PetpetImageURL, "petpet")
	bot.RegisterSubCommand(subcommands.PetpetImageUpload, "petpet")

	bot.RegisterCommand(commands.PetpetUserFromMsgCtx)
	bot.RegisterCommand(commands.PetpetUserCtx)
	bot.RegisterCommand(commands.PetpetImgCtx)

	bot.RegisterCommand(commands.PetpetMessage)
	bot.RegisterCommand(commands.TurnMessageIntoImage)

	if os.Getenv("SYNC_COMMANDS") == "1" {
		chihuahua.LogInfo("Syncing commands with Discord API...")
		if os.Getenv("TEST_SERVER_ID") != "" {
			testServerID, err := chihuahua.StringToSnowflake(os.Getenv("TEST_SERVER_ID"))
			if err != nil {
				chihuahua.LogError("failed to parse TEST_SERVER_ID", err)
			}
			if err = bot.SyncCommandsWithDiscord([]chihuahua.Snowflake{testServerID}); err != nil {
				chihuahua.LogError("failed to sync commands with Discord API", err)
			}
		} else {
			chihuahua.LogInfo("No test server ID provided, syncing commands globally...")
			if err := bot.SyncCommandsWithDiscord([]chihuahua.Snowflake{}); err != nil {
				chihuahua.LogError("failed to sync commands with Discord API", err)
			}
		}
		return
	}

	http.HandleFunc("POST /", bot.DiscordRequestHandler)
	http.HandleFunc("GET /up", bot.UptimeHandler)

	addr := os.Getenv("DISCORD_APP_ADDRESS")
	chihuahua.LogInfo("Serving application at: %s/\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		chihuahua.LogError("something went terribly wrong", err)
	}
}
