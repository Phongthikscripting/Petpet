package commands

import "codeberg.org/lumap/chihuahua"

var Donate chihuahua.Command = chihuahua.Command{
	Name:             "donate",
	Description:      "Want to support PetPet? This is how!",
	Type:             chihuahua.COMMAND_TYPE_CHAT_INPUT,
	IntegrationTypes: []int{chihuahua.COMMAND_INTEGRATION_TYPE_GUILD, chihuahua.COMMAND_INTEGRATION_TYPE_USER},
	Contexts:         []int{chihuahua.COMMAND_CONTEXT_GUILD, chihuahua.COMMAND_CONTEXT_BOT_DM, chihuahua.COMMAND_CONTEXT_PRIVATE_CHANNEL},
	CommandHandler: func(interaction *chihuahua.CommandInteraction) {
		interaction.SendSimpleReply("Hey there! If you're interested in supporting PetPet, you can do do through this link: https://discord.com/discovery/applications/"+interaction.Bot.ApplicationID.String()+"/store\nPetPet's core commands are and will remain free to use for everyone forever. You are not required to help, but it would be greatly appreciated!", true)
	},
}
