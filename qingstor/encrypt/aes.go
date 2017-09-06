package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"log"
	"os"

	"github.com/yunify/qingstor-sdk-go/service"
	"github.com/zhaohuXing/Go/qingstor"
)

var bucket *service.Bucket
var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

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
}

func main() {

	var safer DataSafer
	file, size := getFile("/home/joe/java/vimconfig.tar.gz")

	safer = &Aes256DataSafer{}
	efile, _, _ := safer.EncryptBody("AES256Key-32Characters1234567890", file)

	// Put object
	output, err := bucket.PutObject(
		"vimconfig.tar.gz",
		&service.PutObjectInput{
			ContentLength: size,
			Body:          efile,
		},
	)
	if err != nil {
		log.Println(err)
		panic("put object failed")
	}

	// Print the HTTP status code.
	// Example: 201
	log.Println(service.IntValue(output.StatusCode))

	output1, err := bucket.GetObject("vimconfig.tar.gz", &service.GetObjectInput{})
	if err != nil {
		log.Println(err)
		panic("get object failed")
	}

	// Print the HTTP status code.
	// Example: 200
	log.Println(service.IntValue(output1.StatusCode))
	datas, _ := safer.DecryptBody("AES256Key-32Characters1234567890",
		output1.Body, *output1.ContentLength)

	// Download file.
	dFile, err := os.Create("vimconfig.tar.gz")
	if err != nil {
		log.Println(err)
	}

	sizes, err := dFile.Write(datas)

	if err != nil {
		log.Println(err)
	}
	log.Println(sizes)

}

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

type DataSafer interface {
	EncryptBody(key string, body io.Reader) (io.Reader, *string, error)
	DecryptBody(key string, body io.Reader, length int64) ([]byte, error)
}

type Aes256DataSafer struct{}

func streamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
func (*Aes256DataSafer) EncryptBody(key string, inBody io.Reader) (io.Reader,
	*string, error) {
	body := streamToByte(inBody)

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, nil, err
	}

	cfb := cipher.NewCFBEncrypter(c, commonIV)
	encryptedData := make([]byte, len(body))
	cfb.XORKeyStream(encryptedData, body)
	r := bytes.NewReader(encryptedData)
	return r, nil, nil
}

func (*Aes256DataSafer) DecryptBody(key string, inBody io.Reader,
	length int64) ([]byte, error) {
	body := streamToByte(inBody)

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	decryptedData := make([]byte, length)
	cfbdec.XORKeyStream(decryptedData, body)
	return decryptedData, nil
}
