package commands

import (
	"petpet/utils"
	"slices"

	"codeberg.org/lumap/chihuahua"
	"codeberg.org/lumap/dalmatian"
)

var PetpetUserCtx = chihuahua.Command{
	Type:             2,
	Name:             "Petpet this user",
	Description:      "",
	IntegrationTypes: []int{chihuahua.COMMAND_INTEGRATION_TYPE_GUILD, chihuahua.COMMAND_INTEGRATION_TYPE_USER},
	Contexts:         []int{chihuahua.COMMAND_CONTEXT_GUILD, chihuahua.COMMAND_CONTEXT_BOT_DM, chihuahua.COMMAND_CONTEXT_PRIVATE_CHANNEL},
	CommandHandler: func(interaction *chihuahua.CommandInteraction) {
		userId := interaction.Data.TargetID

		if slices.Contains(utils.BlacklistedUsers, userId.String()) {
			interaction.SendSimpleReply("This user is blacklisted, sorry.", true)
			return
		}

		member := interaction.Data.Resolved.Members[userId]
		user := interaction.Data.Resolved.Users[userId]

		avatar := user.AvatarURL()
		if member != nil && member.GuildAvatarHash != "" && interaction.GuildID != nil {
			avatar = member.GuildAvatarURL(interaction.GuildID.String(), userId.String())
		}

		avatarImg, err := utils.LoadImageFromURL(avatar)
		if err != nil {
			interaction.SendSimpleReply("Failed to load user's avatar.", true)
			return
		}

		// interaction.Defer(false)

		img, err := dalmatian.MakePetImage(avatarImg, 1, 128, 128, false)
		if err != nil {
			interaction.SendSimpleReply(err.Error(), true)
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
