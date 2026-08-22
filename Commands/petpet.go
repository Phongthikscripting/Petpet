package commands

import "codeberg.org/lumap/chihuahua"

var Petpet chihuahua.Command = chihuahua.Command{
	Name:             "petpet",
	Description:      "Petpet someone. Easy.",
	IntegrationTypes: []int{chihuahua.COMMAND_INTEGRATION_TYPE_GUILD, chihuahua.COMMAND_INTEGRATION_TYPE_USER},
	Contexts:         []int{chihuahua.COMMAND_CONTEXT_GUILD, chihuahua.COMMAND_CONTEXT_BOT_DM, chihuahua.COMMAND_CONTEXT_PRIVATE_CHANNEL},
}
