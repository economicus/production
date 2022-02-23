package drivers

import (
	"bytes"
	"economicus/config"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"mime/multipart"
	"net/http"
)

type AWS struct {
	Config   *config.AWSConfig
	Uploader *s3manager.Uploader
}

func NewAWS() *AWS {
	awsConfig := config.NewAWSConfig()
	sess, err := session.NewSession(awsConfig.GetAwsConfig())
	if err != nil {
		log.Fatalf("error while connecting aws: %v", err)
	}
	uploader := s3manager.NewUploader(sess)
	return &AWS{
		Uploader: uploader,
		Config:   awsConfig,
	}
}

func (a *AWS) UploadFile(file multipart.File, header *multipart.FileHeader) (*s3manager.UploadOutput, error) {
	filename := fmt.Sprintf("photos/%s", header.Filename)
	buff := make([]byte, header.Size)
	_, err := file.Read(buff)
	if err != nil {
		return nil, err
	}
	up, err := a.Uploader.Upload(&s3manager.UploadInput{
		Bucket:             aws.String(a.Config.Bucket),
		ACL:                a.Config.ACL,
		CacheControl:       a.Config.CacheControl,
		ContentType:        aws.String(http.DetectContentType(buff)),
		Key:                aws.String(filename),
		Body:               bytes.NewBuffer(buff),
		ContentDisposition: a.Config.ContentDisposition,
	})
	if err != nil {
		return nil, fmt.Errorf("error in UploadFile while uploading file: %w", err)
	}
	return up, nil
}

func (a *AWS) GetFilePath(filename string) string {
	return fmt.Sprintf("https://%s/%s", a.Config.GetDomain(), filename)
}
