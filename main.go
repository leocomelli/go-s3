package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var requiredVariables = []string{"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "AWS_S3_REGION", "AWS_S3_BUCKET", "AWS_S3_OBJECT_PATH"}

func main() {

	values := make(map[string]string, len(requiredVariables))

	for _, v := range requiredVariables {
		value := os.Getenv(v)
		if value != "" {
			values[v] = value
			continue
		}

		log.Fatal(fmt.Sprintf("%s is required", v))
	}

	// Upload
	sess, err := session.NewSession(
		&aws.Config{Region: aws.String(values["AWS_S3_REGION"])},
	)

	if err != nil {
		log.Fatal(err)
	}

	uploader := s3manager.NewUploader(sess)

	path := values["AWS_S3_OBJECT_PATH"]
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	out, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(values["AWS_S3_BUCKET"]),
		Key:    aws.String(filepath.Base(path)),
		Body:   bytes.NewReader(file),
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Location: %s\n", out.Location)

	// List
	svc := s3.New(sess)
	input := &s3.ListObjectsInput{Bucket: aws.String(values["AWS_S3_BUCKET"])}

	result, err := svc.ListObjects(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("*** %s ***\n", values["AWS_S3_BUCKET"])
	for _, item := range result.Contents {
		fmt.Printf("-  %s\n", *item.Key)
	}
}
