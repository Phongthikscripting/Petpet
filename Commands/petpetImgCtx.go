package commands

import "codeberg.org/lumap/chihuahua"

var PetpetImgCtx = chihuahua.Command{
	Type:             4,
	Name:             "Petpet this image",
	Description:      "",
	IntegrationTypes: []int{chihuahua.COMMAND_INTEGRATION_TYPE_GUILD, chihuahua.COMMAND_INTEGRATION_TYPE_USER},
	Contexts:         []int{chihuahua.COMMAND_CONTEXT_GUILD, chihuahua.COMMAND_CONTEXT_BOT_DM, chihuahua.COMMAND_CONTEXT_PRIVATE_CHANNEL},
	CommandHandler: func(interaction *chihuahua.CommandInteraction) {
		interaction.SendSimpleReply("This is out? Oh damn. Please inform the developer about this.", true)
	},
}
