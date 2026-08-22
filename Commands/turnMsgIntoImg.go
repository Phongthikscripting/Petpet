package commands

import (
	"bytes"
	"fmt"

	"petpet/utils"
	"slices"

	"codeberg.org/lumap/chihuahua"
	"codeberg.org/lumap/labrador"
)

var TurnMessageIntoImage = chihuahua.Command{
	Type:             3,
	Name:             "Turn this message into an image",
	Description:      "",
	IntegrationTypes: []int{chihuahua.COMMAND_INTEGRATION_TYPE_GUILD, chihuahua.COMMAND_INTEGRATION_TYPE_USER},
	Contexts:         []int{chihuahua.COMMAND_CONTEXT_GUILD, chihuahua.COMMAND_CONTEXT_BOT_DM, chihuahua.COMMAND_CONTEXT_PRIVATE_CHANNEL},
	CommandHandler: func(interaction *chihuahua.CommandInteraction) {
		hasPremium := interaction.DoesUserHavePremium(false)
		if !hasPremium {
			interaction.SendSimpleReply("Sorry, this command requires Premium. Purchase it at https://discord.com/discovery/applications/"+interaction.Bot.ApplicationID.String()+"/store", true)
			return
		}
		messageId := interaction.Data.TargetID
		msg := interaction.Data.Resolved.Messages[messageId]
		if msg.Type != nil && *msg.Type != chihuahua.MESSAGE_TYPE_DEFAULT && *msg.Type != chihuahua.MESSAGE_TYPE_REPLY && *msg.Type != chihuahua.MESSAGE_TYPE_CHAT_INPUT_COMMAND && *msg.Type != chihuahua.MESSAGE_TYPE_CONTEXT_MENU_COMMAND {
			interaction.SendSimpleReply("Nice try! Unfortunately, this type of message cannot get pet. If you think this is a mistake, contact the developer via Codeberg", true)
			return
		}

		msgAuthor := interaction.Data.Resolved.Messages[messageId].Author
		if slices.Contains(utils.BlacklistedUsers, msgAuthor.ID.String()) {
			interaction.SendSimpleReply("This user is blacklisted, sorry.", true)
			return
		}

		interaction.Defer(false)

		msgData := labrador.MessageData{
			Avatar:        msgAuthor.AvatarURL(),
			Timestamp:     utils.ChangeRFC3339Timestamp(msg.Timestamp),
			Content:       *msg.Content,
			UsernameColor: "#FFFFFF",
		}

		if msgAuthor.GlobalName != "" {
			msgData.Username = msgAuthor.GlobalName
		} else {
			msgData.Username = msgAuthor.Username
		}

		if interaction.GuildID != nil {
			member, err := chihuahua.MakeGetRequestToDiscord[chihuahua.Member]("/guilds/" + interaction.GuildID.String() + "/members/" + msgAuthor.ID.String())
			if err == nil && member != nil {
				if member.Nickname != "" {
					msgData.Username = member.Nickname
				}
				memberRoles := []chihuahua.Role{}
				guildRoles, err := chihuahua.MakeGetRequestToDiscord[[]chihuahua.Role]("/guilds/" + interaction.GuildID.String() + "/roles")
				if err == nil && guildRoles != nil {
					for _, roleId := range member.Roles {
						for _, guildRole := range *guildRoles {
							if guildRole.ID == roleId {
								memberRoles = append(memberRoles, guildRole)
								break
							}
						}
					}
					if len(memberRoles) > 0 {
						slices.SortFunc(memberRoles, func(a, b chihuahua.Role) int {
							return b.Position - a.Position
						})
						filteredRoles := []chihuahua.Role{}
						for _, role := range memberRoles {
							if role.Colors.PrimaryColor != 0 {
								filteredRoles = append(filteredRoles, role)
							}
						}
						if len(filteredRoles) > 0 {
							msgData.UsernameColor = fmt.Sprintf("#%x", filteredRoles[0].Colors.PrimaryColor)
						}
						filteredRoles = []chihuahua.Role{}
						for _, role := range memberRoles {
							if role.Icon != nil {
								filteredRoles = append(filteredRoles, role)
							}
						}
						if len(filteredRoles) > 0 {
							msgData.RoleIconURL = filteredRoles[0].GetIconURL()
						}
					}
				}
			}
		}

		mentions := []string{}
		for _, user := range msg.UserMentions {
			mentions = append(mentions, user.ID.String()+":"+user.Username)
		}
		msgData.Mentions = mentions

		attachments := []string{}
		for _, attachment := range msg.Attachments {
			attachments = append(attachments, attachment.URL)
		}
		msgData.Attachments = attachments

		img, err := labrador.GenerateDiscordMessage(msgData)
		if err != nil {
			interaction.EditReply(chihuahua.ResponseMessageData{
				Content: "Failed to generate message :(",
			}, false, nil)
		}

		interaction.EditReply(chihuahua.ResponseMessageData{
			AllowedMentions: &chihuahua.AllowedMentions{},
		}, false, []chihuahua.DiscordFile{
			{
				Filename: "message.png",
				Reader:   bytes.NewReader(img),
			},
		})
	},
}
