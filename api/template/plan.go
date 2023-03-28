package template

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"server/models"
	"strings"
)

type PlanData struct {
	ID          uint64        `json:"id"`
	Name        string        `json:"name"`
	Width       float32       `json:"width"`
	Height      float32       `json:"height"`
	ImageURL    string        `json:"image_url"`
	ImageHeight int           `json:"image_height"`
	ImageWidth  int           `json:"image_width"`
	Locks       []models.Lock `json:"locks"`
}

type CreatePlanRequest struct {
	Name        string  `json:"name"`
	Width       float32 `json:"width"`
	Height      float32 `json:"height"`
	ImageBase64 string  `json:"image"`
	Image       []byte  `json:""`
}

type AddLockToPlanRequest struct {
	LockID uint64  `json:"lock_id"`
	CoordX float32 `json:"coord_x"`
	CoordY float32 `json:"coord_y"`
}

func (r CreatePlanRequest) Validate() error {
	if r.Name == "" {
		return errors.New("name must not be empty")
	}

	if r.Width == 0 {
		return errors.New("width must not be empty")
	}

	if r.Height == 0 {
		return errors.New("height must not be empty")
	}

	return nil
}

func (r *CreatePlanRequest) SaveImage() (string, image.Image, error) {
	var (
		filename string = strings.ReplaceAll(strings.ToLower(r.Name), " ", "-")
		img      image.Image
	)
	mimeType := strings.Split(r.ImageBase64, ";")[0]
	base64Data := strings.Split(r.ImageBase64, ",")[1]

	imageData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", nil, err
	}
	reader := bytes.NewReader(imageData)

	switch mimeType[5:] {
	case "image/png":
		filename = fmt.Sprintf("%s.png", filename)
		f, err := os.Create(fmt.Sprintf("plans/%s", filename))
		if err != nil {
			return "", nil, err
		}
		defer f.Close()

		img, err = png.Decode(reader)
		if err != nil {
			return "", nil, err
		}

		if err := png.Encode(f, img); err != nil {
			return "", nil, err
		}
	case "image/jpeg":
		filename = fmt.Sprintf("%s.jpeg", filename)
		f, err := os.Create(fmt.Sprintf("plans/%s", filename))
		if err != nil {
			return "", nil, err
		}
		defer f.Close()

		img, err = jpeg.Decode(reader)
		if err != nil {
			return "", nil, err
		}

		if err := jpeg.Encode(f, img, nil); err != nil {
			return "", nil, err
		}
	case "image/jpg":
		filename = fmt.Sprintf("%s.jpg", filename)
		f, err := os.Create(fmt.Sprintf("plans/%s", filename))
		if err != nil {
			return "", nil, err
		}
		defer f.Close()

		img, err = jpeg.Decode(reader)
		if err != nil {
			return "", nil, err
		}

		if err := jpeg.Encode(f, img, nil); err != nil {
			return "", nil, err
		}
	case "image/gif":
		filename = fmt.Sprintf("%s.gif", filename)
		f, err := os.Create(fmt.Sprintf("plans/%s", filename))
		if err != nil {
			return "", nil, err
		}
		defer f.Close()

		img, err = gif.Decode(reader)
		if err != nil {
			return "", nil, err
		}

		options := &gif.Options{NumColors: 256, Quantizer: nil, Drawer: nil}
		if err := gif.Encode(f, img, options); err != nil {
			return "", nil, err
		}
	default:
		return "", nil, errors.New("unsupported mime type")
	}

	return filename, img, nil
}

func (r AddLockToPlanRequest) Validate() error {
	if r.LockID == 0 {
		return errors.New("lock_id must not be empty")
	}

	return nil
}
