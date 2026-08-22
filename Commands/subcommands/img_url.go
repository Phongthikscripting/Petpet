package petpetsubcommands

import (
	"fmt"
	"petpet/utils"
	"runtime"
	"sync"

	"codeberg.org/lumap/chihuahua"
	"codeberg.org/lumap/dalmatian"
)

type urlResult struct {
	file chihuahua.DiscordFile
	err  error
}

func processURL(interaction *chihuahua.CommandInteraction, imageURL string) (chihuahua.DiscordFile, error) {
	isImage, err := utils.IsLinkAnImageURL(imageURL)
	if err != nil {
		return chihuahua.DiscordFile{}, fmt.Errorf("The provided URL is invalid.")
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

	loadedImg, err := utils.LoadImageFromURL(imageURL)
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

var PetpetImageURL = chihuahua.Command{
	Name:        "image_url",
	Description: "Petpet an image (via external URL)",
	Options:     append(utils.PetpetCommandImageURLOptions, utils.PetpetCommandOptions...),
	CommandHandler: func(interaction *chihuahua.CommandInteraction) {
		ephemeral, err := interaction.GetBoolOptionValue("ephemeral", false)
		if err != nil {
			interaction.SendSimpleReply("Couldn't parse ephemeral option.", true)
			return
		}

		var imageURLs []string
		for i := 1; i <= 10; i++ {
			optionName := "image_url"
			if i > 1 {
				optionName += fmt.Sprint(i)
			}
			url, err := interaction.GetStringOptionValue(optionName, "")
			if err != nil || url == "" {
				continue
			}
			imageURLs = append(imageURLs, url)
		}

		if len(imageURLs) == 0 {
			interaction.SendSimpleReply("No valid image URLs provided.", true)
			return
		}

		fast := (len(imageURLs) == 1) && (len(interaction.Data.Options) == 1)
		if !fast {
			interaction.Defer(ephemeral)
		}

		results := make([]urlResult, len(imageURLs))
		var wg sync.WaitGroup
		sem := make(chan struct{}, runtime.NumCPU())

		for i, url := range imageURLs {
			wg.Add(1)

			go func(url string, i int) {
				defer wg.Done()

				sem <- struct{}{}
				defer func() { <-sem }()

				file, err := processURL(interaction, url)
				results[i] = urlResult{file: file, err: err}
			}(url, i)
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

		content := "<@" + interaction.GetUser().ID.String() + "> has pet an image :3"
		if len(imageURLs) > 1 {
			content = "<@" + interaction.GetUser().ID.String() + "> has pet images :3"
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
