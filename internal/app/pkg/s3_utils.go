package pkg

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func UploadS3(key string, fileContent []byte, fileType string) (string, error) {
	region, _ := os.LookupEnv("AWS_REGION")
	bucket, _ := os.LookupEnv("AWS_BUCKET")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return "", err
	}

	svc := s3.New(sess)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(fileContent),
		ContentType: aws.String(fileType),
	})
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return "", err
	}

	return fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", bucket, region, key), nil
}

func GetFileFromS3(key string, fileType string) ([]byte, error) {
	region, _ := os.LookupEnv("AWS_REGION")
	bucket, _ := os.LookupEnv("AWS_BUCKET")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	result, err := svc.GetObject(&s3.GetObjectInput{
		Bucket:              aws.String(bucket),
		Key:                 aws.String(key),
		ResponseContentType: aws.String(fileType),
	})
	if err != nil {
		fmt.Println("Error getting file:", err)
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(result.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
