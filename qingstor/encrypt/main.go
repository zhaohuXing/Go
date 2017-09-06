package main

import (
	"log"
	"os"

	"github.com/zhaohuXing/Go/qingstor"

	"github.com/yunify/qingstor-sdk-go/client/crypto"
	qs "github.com/yunify/qingstor-sdk-go/service"
	"github.com/yunify/qingstor-sdk-go/utils"
)

// Initialize file path
var filePath string = "/home/joe/Pictures/zw.png"

var server *crypto.EncryptionServer
var bucket *qs.Bucket

func init() {
	// Get QingStor Service
	service, err := qingstor.GetService()
	if err != nil {
		log.Println(err)
		panic("get qingstor service failed")
	}

	// Get QingStor Bucket
	bucket, err = service.Bucket("test002", "pek3a")
	if err != nil {
		log.Println(err)
		panic("get qingstor bucket failed")
	}

	// Initialize encryption server
	algorithm := "AES256"
	key := "AES256Key-32Characters1234567890"
	server, err = crypto.NewEncryptionServer(bucket, &algorithm, &key)
	if err != nil {
		log.Println(err)
		panic("Initialize encryption server failed")
	}
}

func main() {
	//	PutObject()
	//	GetObject()
	//HeadObject()
	//	uploadID := InitiateMultipartUpload()
	//	UploadMultipart(uploadID)
	//	PutObjectCopy()

	//	PutObjectMove()

	//	uploadID, objectKey := InitiateMultipartUpload("sha1.png")
	//	UploadMultipart(uploadID, objectKey)
	//	CompleteMultipartUpload(uploadID, objectKey)

	PutObject()
	uploadID1, objectKey := InitiateMultipartUpload("sha2.png")
	UploadMultipartCopy(uploadID1, objectKey)
}

// GetObject only to test GetObject method of EncryptionServer
func GetObject() {
	output, err := server.GetObject("zw.png", &qs.GetObjectInput{})
	if err != nil {
		log.Println(err)
		panic("get object failed")
	}

	// Print the HTTP status code.
	// Example: 200
	log.Println(qs.IntValue(output.StatusCode))
}

// PutObject only to test PutObject method of EncryptionServer
func PutObject() {

	file, size := getFile(filePath)

	// Put object
	output, err := server.PutObject(
		"zw.png",
		&qs.PutObjectInput{
			ContentLength: size,
			Body:          file,
		},
	)
	defer file.Close()
	if err != nil {
		log.Println(err)
		panic("put object failed")
	}

	// Print the HTTP status code.
	// Example: 201
	log.Println(qs.IntValue(output.StatusCode))
}

// PutObjectCopy only to test PutObjectCopy method of EncryptionServer
func PutObjectCopy() {

	source := utils.URLQueryEscape("/test002/zw.png")
	// Put object copy
	output, err := server.PutObjectCopy(
		"sprint.png",
		&qs.PutObjectInput{
			XQSCopySource: &source,
		})
	if err != nil {
		log.Println(err)
		panic("put object copy failed")
	}

	// Print the HTTP status code.
	// Example: 201
	log.Println(qs.IntValue(output.StatusCode))

}

// PutObjectMove only to test PutObjectMove method of EncryptionServer
func PutObjectMove() {
	source := utils.URLQueryEscape("/test002/zw.png")

	// Put object copy
	output, err := server.PutObjectMove(
		"/test/zw.png",
		&qs.PutObjectInput{
			XQSMoveSource: &source,
		})
	if err != nil {
		log.Println(err)
		panic("put object copy failed")
	}

	// Print the HTTP status code.
	// Example: 201
	log.Println(qs.IntValue(output.StatusCode))

}

// Notice: the way in action not imitate
// Here's just for me to test the feature
func getFile(path string) (*os.File, *int64) {
	// Open file operartion
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		panic("open file failed")
	}

	fileInfo, err := file.Stat()
	if err != nil {
		log.Println(err)
		panic("get fileinfo failed")
	}
	size := fileInfo.Size()
	return file, &size
}

// HeadObject only to test HeadObject method of EncryptionServer
func HeadObject() {
	output, err := server.HeadObject("zw.png", &qs.HeadObjectInput{})
	if err != nil {
		log.Println(err)
		panic("get head object failed")
	}

	// Print the HTTP status code.
	// Example: 200
	log.Println(qs.IntValue(output.StatusCode))
}

// InitiateMultipartUpload only to test InitializeMultipartUpload
// method of EncryptionServer
func InitiateMultipartUpload(objectKey string) (*string, string) {
	output, err := server.InitiateMultipartUpload(objectKey,
		&qs.InitiateMultipartUploadInput{})
	if err != nil {
		log.Println(err)
		panic("initialize multipart upload failed")
	}

	// Print the HTTP status code.
	// Example: 200
	log.Println(qs.IntValue(output.StatusCode))
	return output.UploadID, *output.Key
}

// UploadMultipart only to test UploadMultipart method of EncryptionServer
func UploadMultipart(uploadID *string, objectKey string) {
	number := 2
	file, _ := getFile(filePath)
	output, err := server.UploadMultipart(objectKey,
		&qs.UploadMultipartInput{
			UploadID:   uploadID,
			PartNumber: &number,
			Body:       file,
		})
	defer file.Close()
	if err != nil {
		log.Println(err)
		panic("upload multipart failed")
	}

	// Print the HTTP status code.
	// Example: 201
	log.Println(qs.IntValue(output.StatusCode))
}

func UploadMultipartCopy(uploadID *string, objectKey string) {
	number := 2
	file, _ := getFile(filePath)
	source := utils.URLQueryEscape("/test002/zw.png")
	output, err := server.UploadMultipartCopy(objectKey,
		&qs.UploadMultipartInput{
			UploadID:      uploadID,
			PartNumber:    &number,
			XQSCopySource: &source,
		})
	defer file.Close()
	if err != nil {
		log.Println(err)
		panic("upload multipart copy failed")
	}

	// Print the HTTP status code.
	// Example: 201
	log.Println(qs.IntValue(output.StatusCode))
}

func CompleteMultipartUpload(uploadID *string, objectKey string) {

	aOutput, _ := bucket.CompleteMultipartUpload(
		objectKey,
		&qs.CompleteMultipartUploadInput{
			UploadID: uploadID,
			ObjectParts: []*qs.ObjectPartType{{
				PartNumber: qs.Int(2),
			}},
		})

	// Print the HTTP status code.
	// Example: 200
	log.Println(qs.IntValue(aOutput.StatusCode))
}
