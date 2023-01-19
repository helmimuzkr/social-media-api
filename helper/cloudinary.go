package helper

import (
	"context"
	"social-media-app/config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFile(file interface{}, path string) (*uploader.UploadResult, error) {
	// create cloudinary with configuration
	cld, err := cloudinary.NewFromParams(config.CloudinaryName, config.CloudinaryApiKey, config.CloudinaryApiScret)
	if err != nil {
		return nil, err
	}

	// upload file
	uploadResult, err := cld.Upload.Upload(
		context.Background(),
		file,
		uploader.UploadParams{
			ResourceType: "image",
			Folder:       config.CloudinaryUploadFolder + path,
		})
	if err != nil {
		return nil, err
	}

	return uploadResult, nil
}

func DestroyFile(publicID string) error {
	// create cloudinary with configuration
	cld, err := cloudinary.NewFromParams(config.CloudinaryName, config.CloudinaryApiKey, config.CloudinaryApiScret)
	if err != nil {
		return err
	}

	// upload file
	_, err = cld.Upload.Destroy(
		context.Background(),
		uploader.DestroyParams{
			PublicID: publicID,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
