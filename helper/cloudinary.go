package helper

import (
	"context"
	"log"
	"social-media-app/config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func Upload(file interface{}, path string) (string, error) {
	// create cloudinary with configuration
	cld, err := cloudinary.NewFromParams(config.CloudinaryName, config.CloudinaryApiKey, config.CloudinaryApiScret)
	if err != nil {
		return "", err
	}

	// upload file
	uploadResult, err := cld.Upload.Upload(
		context.Background(),
		file,
		uploader.UploadParams{
			UniqueFilename: api.Bool(true),
			ResourceType:   "image",
			Folder:         config.CloudinaryUploadFolder + path,
		})
	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
		return "", err
	}

	return uploadResult.SecureURL, nil
}
