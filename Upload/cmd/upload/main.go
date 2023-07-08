package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"AKIASOABMIYXN4QNLYVP",
				"HRiewAv09p0o2IFE2Eqt57aHRNN3RzBnQaQ+XSA4",
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}
	s3Client = s3.New(sess)
	s3Bucket = "goexpertbucket"
}

func main() {
	dir, err := os.Open("/tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	uploadControl := make(chan struct{}, 10)
	errorFileUpload := make(chan string, 10)

	go func() {
		for {
			select {
			case filename := <-errorFileUpload:
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(filename, uploadControl, errorFileUpload)
			}
		}
	}()
	for {
		files, err := dir.Readdir(10)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s\n", err.Error())
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
}

func uploadFile(filename string, uploadControl <-chan struct{}, errorUploadFile chan<- string) {
	defer wg.Done()
	completeFileName := fmt.Sprintf("/tmp/%s", filename)
	fmt.Printf("Uploding file %s to bucket %s", completeFileName, s3Bucket)
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error openning file %s\n", completeFileName)
		<-uploadControl
		errorUploadFile <- completeFileName
		return
	}
	defer f.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("Error upload file %s\n", completeFileName)
		<-uploadControl
		errorUploadFile <- completeFileName
		return
	}
	fmt.Printf("File %s upload successful\n", completeFileName)
	<-uploadControl
}
