package commands

import (
	"petpet/utils"
	"slices"

	"codeberg.org/lumap/chihuahua"
	"codeberg.org/lumap/dalmatian"
)

var PetpetUserFromMsgCtx = chihuahua.Command{
	Type:             3,
	Name:             "Petpet the message's author",
	Description:      "",
	IntegrationTypes: []int{chihuahua.COMMAND_INTEGRATION_TYPE_GUILD, chihuahua.COMMAND_INTEGRATION_TYPE_USER},
	Contexts:         []int{chihuahua.COMMAND_CONTEXT_GUILD, chihuahua.COMMAND_CONTEXT_BOT_DM, chihuahua.COMMAND_CONTEXT_PRIVATE_CHANNEL},
	CommandHandler: func(interaction *chihuahua.CommandInteraction) {
		messageId := interaction.Data.TargetID
		user := interaction.Data.Resolved.Messages[messageId].Author

		if slices.Contains(utils.BlacklistedUsers, user.ID.String()) {
			interaction.SendSimpleReply("This user is blacklisted, sorry.", true)
			return
		}

		avatar, err := utils.LoadImageFromURL(user.AvatarURL())
		if err != nil {
			interaction.SendSimpleReply("Failed to load user's avatar.", true)
			return
		}

		// interaction.Defer(false)

		img, err := dalmatian.MakePetImage(avatar, 1, 128, 128, false)
		if err != nil {
			interaction.SendSimpleReply(err.Error(), false)
			// interaction.EditReply(chihuahua.ResponseMessageData{
			// 	Content: err.Error(),
			// }, false, nil)
			return
		}

		interaction.SendReply(chihuahua.ResponseMessageData{
			Content:         "<@" + interaction.GetUser().ID.String() + "> has pet <@" + user.ID.String() + "> :33",
			AllowedMentions: &chihuahua.AllowedMentions{},
		}, false, []chihuahua.DiscordFile{
			{
				Filename: "petpet.gif",
				Reader:   img,
			},
		})
	},
}
