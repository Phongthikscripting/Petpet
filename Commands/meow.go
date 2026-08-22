package commands

import "codeberg.org/lumap/chihuahua"

var Meow chihuahua.Command = chihuahua.Command{
	Name:             "meow",
	Description:      "Meow, meow!",
	Type:             chihuahua.COMMAND_TYPE_CHAT_INPUT,
	IntegrationTypes: []int{chihuahua.COMMAND_INTEGRATION_TYPE_GUILD, chihuahua.COMMAND_INTEGRATION_TYPE_USER},
	Contexts:         []int{chihuahua.COMMAND_CONTEXT_GUILD, chihuahua.COMMAND_CONTEXT_BOT_DM, chihuahua.COMMAND_CONTEXT_PRIVATE_CHANNEL},
	CommandHandler: func(interaction *chihuahua.CommandInteraction) {
		interaction.SendSimpleReply("Mrrowww~", false)
	},
}
