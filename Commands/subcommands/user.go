package petpetsubcommands

import (
	"fmt"
	"petpet/utils"
	"runtime"
	"slices"
	"strings"
	"sync"

	"codeberg.org/lumap/chihuahua"
	"codeberg.org/lumap/dalmatian"
)

type userResult struct {
	file chihuahua.DiscordFile
	err  error
}

func processUser(interaction *chihuahua.CommandInteraction, untypedUserId string) (chihuahua.DiscordFile, error) {
	userId, err := chihuahua.StringToSnowflake(untypedUserId)
	if err != nil {
		return chihuahua.DiscordFile{}, fmt.Errorf("Couldn't parse a user ID")
	}

	member := interaction.Data.Resolved.Members[userId]
	user := interaction.Data.Resolved.Users[userId]

	useServerAvatar, err := interaction.GetBoolOptionValue("use_server_avatar", true)
	if err != nil {
		return chihuahua.DiscordFile{}, fmt.Errorf("Couldn't parse use_server_avatar option.")
	}

	avatar := user.AvatarURL()
	if useServerAvatar && member != nil && member.GuildAvatarHash != "" {
		avatar = member.GuildAvatarURL(interaction.GuildID.String(), userId.String())
	}

	avatarImg, err := utils.LoadImageFromURL(avatar)
	if err != nil {
		return chihuahua.DiscordFile{}, fmt.Errorf("Failed to load user's avatar.")
	}

	speed, err := interaction.GetIntOptionValue("speed", 8)
	if err != nil {
		return chihuahua.DiscordFile{}, fmt.Errorf("Couldn't parse speed option.")
	}
	width, err := interaction.GetIntOptionValue("width", 128)
	if err != nil {
		return chihuahua.DiscordFile{}, fmt.Errorf("Couldn't parse width option.")
	}
	height, err := interaction.GetIntOptionValue("height", 128)
	if err != nil {
		return chihuahua.DiscordFile{}, fmt.Errorf("Couldn't parse height option.")
	}
	stretchHand, err := interaction.GetBoolOptionValue("stretch_hand", false)
	if err != nil {
		interaction.SendSimpleReply("Couldn't parse stretch_hand option", true)
	}

	img, err := dalmatian.MakePetImage(avatarImg, speed, width, height, stretchHand)
	if err != nil {
		return chihuahua.DiscordFile{}, err
	}

	return chihuahua.DiscordFile{
		Filename: "petpet.gif",
		Reader:   img,
	}, nil
}

var PetpetUser = chihuahua.Command{
	Name:        "user",
	Description: "Petpet someone's pfp",
	Options:     append(utils.PetpetCommandUserOptions, utils.PetpetCommandOptions...),
	CommandHandler: func(interaction *chihuahua.CommandInteraction) {
		ephemeral, err := interaction.GetBoolOptionValue("ephemeral", false)
		if err != nil {
			interaction.SendSimpleReply("Couldn't parse ephemeral option.", true)
			return
		}
		mentions := []string{}
		if notify, err := interaction.GetBoolOptionValue("notify_users", false); err != nil {
			interaction.SendSimpleReply("Couldn't parse notify_users option.", true)
			return
		} else if notify {
			mentions = append(mentions, "users")
		}

		var (
			blacklistDetected bool
			userIDs           []string
			files             []chihuahua.DiscordFile
		)

		for i := 1; i <= 10; i++ {
			optionName := "user"
			if i > 1 {
				optionName += fmt.Sprint(i)
			}
			user, err := interaction.GetStringOptionValue(optionName, "")
			if err != nil || user == "" {
				continue
			}
			if slices.Contains(utils.BlacklistedUsers, user) {
				blacklistDetected = true
				continue
			}
			userIDs = append(userIDs, user)
		}

		fast := (len(userIDs) == 1) && (len(interaction.Data.Options) == 1)
		if !fast {
			interaction.Defer(ephemeral)
		}

		results := make([]userResult, 10)
		var wg sync.WaitGroup
		sem := make(chan struct{}, runtime.NumCPU())

		for i, user := range userIDs {
			wg.Add(1)

			go func(u string, i int) {
				defer wg.Done()

				sem <- struct{}{}
				defer func() { <-sem }()

				file, err := processUser(interaction, u)
				results[i] = userResult{file: file, err: err}
			}(user, i)
		}

		wg.Wait()

		for _, res := range results {
			emptyFile := chihuahua.DiscordFile{}
			if res.file == emptyFile {
				continue
			}
			if res.err != nil {
				if fast {
					interaction.SendSimpleReply(res.err.Error(), true)
				} else {
					interaction.EditReply(chihuahua.ResponseMessageData{
						Content: res.err.Error(),
					}, true, []chihuahua.DiscordFile{})
				}
				return
			}
			files = append(files, res.file)
		}

		content := "<@" + interaction.GetUser().ID.String() + "> has pet "
		for i, id := range userIDs {
			userIDs[i] = fmt.Sprintf("<@%s>", id)
		}

		mentionedUsers := strings.Join(userIDs, ", ")
		if len(userIDs) > 1 {
			lastComma := strings.LastIndex(mentionedUsers, ", ")
			mentionedUsers = mentionedUsers[:lastComma] + " and" + mentionedUsers[lastComma+1:]
		}
		content += mentionedUsers + " :3"
		if blacklistDetected {
			content += "\n-# A user you tried to petpet has been blacklisted and was ignored."
		}

		if fast {
			interaction.SendReply(chihuahua.ResponseMessageData{
				Content:         content,
				AllowedMentions: &chihuahua.AllowedMentions{Parse: mentions},
			}, ephemeral, files)
		} else {
			interaction.EditReply(chihuahua.ResponseMessageData{
				Content:         content,
				AllowedMentions: &chihuahua.AllowedMentions{Parse: mentions},
			}, ephemeral, files)
		}
	},
}
