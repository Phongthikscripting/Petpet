package petpetsubcommands

import (
	"fmt"
	"petpet/utils"
	"runtime"
	"sync"

	"codeberg.org/lumap/chihuahua"
	"codeberg.org/lumap/dalmatian"
)

type imageResult struct {
	file chihuahua.DiscordFile
	err  error
}

func processImage(interaction *chihuahua.CommandInteraction, untypedImageId string) (chihuahua.DiscordFile, error) {
	imageId, err := chihuahua.StringToSnowflake(untypedImageId)
	if err != nil {
		return chihuahua.DiscordFile{}, fmt.Errorf("Couldn't parse image ID.")
	}

	image := interaction.Data.Resolved.Attachments[imageId]

	isImage, err := utils.IsLinkAnImageURL(image.URL)
	if err != nil {
		return chihuahua.DiscordFile{}, fmt.Errorf("Couldn't check if the URL is an image.")
	}
	if !isImage {
		return chihuahua.DiscordFile{}, fmt.Errorf("The provided URL is not an image.")
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
		return chihuahua.DiscordFile{}, fmt.Errorf("Couldn't parse stretch_hand option.")
	}

	loadedImg, err := utils.LoadImageFromURL(image.URL)
	if err != nil {
		return chihuahua.DiscordFile{}, fmt.Errorf("Failed to load the image.")
	}

	img, err := dalmatian.MakePetImage(loadedImg, speed, width, height, stretchHand)
	if err != nil {
		return chihuahua.DiscordFile{}, err
	}

	return chihuahua.DiscordFile{
		Filename: "petpet.gif",
		Reader:   img,
	}, nil
}

var PetpetImageUpload = chihuahua.Command{
	Name:        "image_upload",
	Description: "Petpet an uploaded image",
	Options:     append(utils.PetpetCommandImageUploadOptions, utils.PetpetCommandOptions...),
	CommandHandler: func(interaction *chihuahua.CommandInteraction) {
		ephemeral, err := interaction.GetBoolOptionValue("ephemeral", false)
		if err != nil {
			interaction.SendSimpleReply("Couldn't parse ephemeral option.", true)
			return
		}

		var imageIDs []string
		for i := 1; i <= 10; i++ {
			optionName := "image_upload"
			if i > 1 {
				optionName += fmt.Sprint(i)
			}
			untypedImage, err := interaction.GetAttachmentOptionId(optionName, "")
			if err != nil || untypedImage == "" {
				continue
			}
			imageIDs = append(imageIDs, untypedImage)
		}

		if len(imageIDs) == 0 {
			interaction.SendSimpleReply("No valid images provided.", true)
			return
		}

		fast := (len(imageIDs) == 1) && (len(interaction.Data.Options) == 1)
		if !fast {
			interaction.Defer(ephemeral)
		}

		results := make([]imageResult, len(imageIDs))
		var wg sync.WaitGroup
		sem := make(chan struct{}, runtime.NumCPU())

		for i, id := range imageIDs {
			wg.Add(1)

			go func(id string, i int) {
				defer wg.Done()

				sem <- struct{}{}
				defer func() { <-sem }()

				file, err := processImage(interaction, id)
				results[i] = imageResult{file: file, err: err}
			}(id, i)
		}

		wg.Wait()

		var files []chihuahua.DiscordFile
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
					}, ephemeral, []chihuahua.DiscordFile{})
				}
				return
			}
			files = append(files, res.file)
		}

		content := "<@" + interaction.GetUser().ID.String() + "> has pet an uploaded image :3"
		if len(imageIDs) > 1 {
			content = "<@" + interaction.GetUser().ID.String() + "> has pet uploaded images :3"
		}

		if fast {
			interaction.SendReply(chihuahua.ResponseMessageData{
				Content:         content,
				AllowedMentions: &chihuahua.AllowedMentions{},
			}, ephemeral, files)
		} else {
			interaction.EditReply(chihuahua.ResponseMessageData{
				Content:         content,
				AllowedMentions: &chihuahua.AllowedMentions{},
			}, ephemeral, files)
		}
	},
}