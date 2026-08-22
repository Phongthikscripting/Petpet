package commands

import "codeberg.org/lumap/chihuahua"

var Support chihuahua.Command = chihuahua.Command{
	Name:             "support",
	Description:      "Have a question? Want help? This is the way",
	Type:             chihuahua.COMMAND_TYPE_CHAT_INPUT,
	IntegrationTypes: []int{chihuahua.COMMAND_INTEGRATION_TYPE_GUILD, chihuahua.COMMAND_INTEGRATION_TYPE_USER},
	Contexts:         []int{chihuahua.COMMAND_CONTEXT_GUILD, chihuahua.COMMAND_CONTEXT_BOT_DM, chihuahua.COMMAND_CONTEXT_PRIVATE_CHANNEL},
	CommandHandler: func(interaction *chihuahua.CommandInteraction) {
		interaction.SendSimpleReply("Help can be provided by opening an issue at https://codeberg.org/lumap/petpet/issues/new. If you don't want to use Codeberg, send the developer an email directly at <lumap@duck.com>!", true)
	},
}
