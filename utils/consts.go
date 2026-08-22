package utils

import "codeberg.org/lumap/chihuahua"

var BlacklistedUsers = []string{
	"1118171067463241868",
}

var PetpetCommandOptions = []chihuahua.CommandOption{
	{
		Type:        4,
		Name:        "width",
		Description: "The width of the gif (default is 128)",
		Required:    false,
		MinValue:    8,
		MaxValue:    1024,
	},
	{
		Type:        4,
		Name:        "height",
		Description: "The height of the gif (default is 128)",
		Required:    false,
		MinValue:    8,
		MaxValue:    1024,
	},
	{
		Type:        4,
		Name:        "speed",
		Description: "How fast the petting is",
		Required:    false,
		Choices: []chihuahua.CommandOptionChoice{
			{
				Name:  "Fastest",
				Value: 2,
			},
			{
				Name:  "Faster",
				Value: 4,
			},
			{
				Name:  "Fast",
				Value: 6,
			},
			{
				Name:  "Default",
				Value: 8,
			},
			{
				Name:  "Slow",
				Value: 10,
			},
			{
				Name:  "Slower",
				Value: 12,
			},
			{
				Name:  "Even Slower",
				Value: 15,
			},
			{
				Name:  "Snail",
				Value: 50,
			},
			{
				Name:  "Is it even moving?",
				Value: 1000,
			},
		},
	},
	{
		Type:        5,
		Name:        "ephemeral",
		Description: "Whether or not to make the message ephemeral (default is false)",
		Required:    false,
	},
	{
		Type:        5,
		Name:        "stretch_hand",
		Description: "Whether to stretch the hand or not to fit the custom width/height (defaults to false)",
		Required:    false,
	},
}

var PetpetCommandUserOptions = []chihuahua.CommandOption{
	{
		Type:        6,
		Name:        "user",
		Description: "The user to petpet",
		Required:    true,
	},
	{
		Type:        6,
		Name:        "user2",
		Description: "The second user to petpet",
		Required:    false,
	},
	{
		Type:        6,
		Name:        "user3",
		Description: "The third user to petpet",
		Required:    false,
	},
	{
		Type:        6,
		Name:        "user4",
		Description: "The fourth user to petpet",
		Required:    false,
	},
	{
		Type:        6,
		Name:        "user5",
		Description: "The fifth user to petpet",
		Required:    false,
	},
	{
		Type:        6,
		Name:        "user6",
		Description: "The sixth user to petpet",
		Required:    false,
	},
	{
		Type:        6,
		Name:        "user7",
		Description: "The seventh user to petpet",
		Required:    false,
	},
	{
		Type:        6,
		Name:        "user8",
		Description: "The eighth user to petpet",
		Required:    false,
	},
	{
		Type:        6,
		Name:        "user9",
		Description: "The ninth user to petpet",
		Required:    false,
	},
	{
		Type:        6,
		Name:        "user10",
		Description: "The tenth user to petpet",
		Required:    false,
	},
	{
		Type:        5,
		Name:        "use_server_avatar",
		Description: "Whether to use the server avatars (default is true; applies to all users)",
		Required:    false,
	},
	{
		Type:        5,
		Name:        "notify_users",
		Description: "Whether to notify the users that they've been petpet (default is false)",
		Required:    false,
	},
}

var PetpetCommandImageURLOptions = []chihuahua.CommandOption{
	{
		Type:        3,
		Name:        "image_url",
		Description: "The image's URL",
		Required:    true,
	},
	{
		Type:        3,
		Name:        "image_url2",
		Description: "The second image's URL",
		Required:    false,
	},
	{
		Type:        3,
		Name:        "image_url3",
		Description: "The third image's URL",
		Required:    false,
	},
	{
		Type:        3,
		Name:        "image_url4",
		Description: "The fourth image's URL",
		Required:    false,
	},
	{
		Type:        3,
		Name:        "image_url5",
		Description: "The fifth image's URL",
		Required:    false,
	},
	{
		Type:        3,
		Name:        "image_url6",
		Description: "The sixth image's URL",
		Required:    false,
	},
	{
		Type:        3,
		Name:        "image_url7",
		Description: "The seventh image's URL",
		Required:    false,
	},
	{
		Type:        3,
		Name:        "image_url8",
		Description: "The eighth image's URL",
		Required:    false,
	},
	{
		Type:        3,
		Name:        "image_url9",
		Description: "The ninth image's URL",
		Required:    false,
	},
	{
		Type:        3,
		Name:        "image_url10",
		Description: "The tenth image's URL",
		Required:    false,
	},
}

var PetpetCommandImageUploadOptions = []chihuahua.CommandOption{
	{
		Type:        11,
		Name:        "image_upload",
		Description: "The image to petpet",
		Required:    true,
	},
	{
		Type:        11,
		Name:        "image_upload2",
		Description: "The second image to petpet",
		Required:    false,
	},
	{
		Type:        11,
		Name:        "image_upload3",
		Description: "The third image to petpet",
		Required:    false,
	},
	{
		Type:        11,
		Name:        "image_upload4",
		Description: "The fourth image to petpet",
		Required:    false,
	},
	{
		Type:        11,
		Name:        "image_upload5",
		Description: "The fifth image to petpet",
		Required:    false,
	},
	{
		Type:        11,
		Name:        "image_upload6",
		Description: "The sixth image to petpet",
		Required:    false,
	},
	{
		Type:        11,
		Name:        "image_upload7",
		Description: "The seventh image to petpet",
		Required:    false,
	},
	{
		Type:        11,
		Name:        "image_upload8",
		Description: "The eighth image to petpet",
		Required:    false,
	},
	{
		Type:        11,
		Name:        "image_upload9",
		Description: "The ninth image to petpet",
		Required:    false,
	},
	{
		Type:        11,
		Name:        "image_upload10",
		Description: "The tenth image to petpet",
		Required:    false,
	},
}
