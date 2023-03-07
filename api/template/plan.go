package template

import (
	"encoding/base64"
	"errors"
	"server/models"
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

func (r *CreatePlanRequest) LoadImage() error {
	r.Image = make([]byte, base64.StdEncoding.DecodedLen(len(r.ImageBase64)))
	n, err := base64.StdEncoding.Decode(r.Image, []byte(r.ImageBase64))
	if err != nil {
		return err
	}

	r.Image = r.Image[:n]

	return nil
}

func (r AddLockToPlanRequest) Validate() error {
	if r.LockID == 0 {
		return errors.New("lock_id must not be empty")
	}

	return nil
}
