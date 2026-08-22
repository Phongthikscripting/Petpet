package utils

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"net/http"
	"strings"
	"time"

	"codeberg.org/lumap/chihuahua"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/webp"
)

func IsLinkAnImageURL(url string) (bool, error) {
	resp, err := http.Head(url)
	if err != nil {
		return false, err
	}

	contentType := resp.Header.Get("Content-Type")

	return strings.HasPrefix(contentType, "image/"), nil
}

func LoadImageFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer chihuahua.CloseBody(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch image: status code %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func DecodeStillImage(data []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	return img, err
}

func ChangeRFC3339Timestamp(ts string) string {
	t, _ := time.Parse(time.RFC3339, ts)
	return t.Format("02/01/2006, 15:04")
}
