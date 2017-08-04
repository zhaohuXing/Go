# QingStor Image Processing Usage Guide

For the user stored in the QingStor object storage on a variety of basic processing, such as format conversion, cutting, flip, watermark and so on.

Currently supported image formats are:
- png
- tiff
- webp
- jpeg
- pdf
- gif
- svg

> Comments are currently not supported on the encrypted image after processing, a single picture up to 10M.

## Request Syntax

```
GET /<object-name>?image&action=<action>
Host: <bucket-name>.pek3a.qingstor.com
Date: <date>
Authorization: <authorization-string>
```

- The action represents a set of operations on a picture.
- The format of action is operation: k_v [, k_v] [| operation: k_v] [, k_v].
- The operation represents the basic operation of the picture, such as crop, watermark and so on. Each operation can be followed by multiple key value pair as a parameter.
- The k is the argument key of operation, and v is the argument value.
- Multiple operations are concatenated as an action, which will sequentially manipulate the image, similar to the pipe.

## Example Request
```
GET /myphoto.jpg?image&action=resize:w_300,h_400|rotate:a_90 HTTP/1.1
Host: mybucket.pek3a.qingstor.com
Date: Sun, 16 Aug 2015 09:05:00 GMT
Authorization: authorization string
```
The above example will be fixed in accordance with 300 * 400 (px) fixed width and height of the abbreviated, and flip 90 degrees.


## Usage
Before using the image service, you need to initialize the [Configuration](https://github.com/yunify/qingstor-sdk-go/blob/master/docs/configuration.md) and [QingStor Service](https://github.com/yunify/qingstor-sdk-go/blob/master/docs/qingstor_service_usage.md).

Here to provide a way to initialize, the other view [Configuration](https://github.com/yunify/qingstor-sdk-go/blob/master/docs/configuration.md) and [QingStor Service](https://github.com/yunify/qingstor-sdk-go/blob/master/docs/qingstor_service_usage.md).
```
//Import the latest version API
import (
	"github.com/yunify/qingstor-sdk-go/config"
	"github.com/yunify/qingstor-sdk-go/helpers"
	qs "github.com/yunify/qingstor-sdk-go/service"
)
```

## Code Snippet

The definition of the structure of the image processing
```
import "github.com/yunify/qingstor-sdk-go/service"
// helpers/image.go
type Image struct {
	ImageOutput *service.ImageProcessOutput
	bucket      *service.Bucket
	uri         *string
	name        *string
	Err         error
}

// About cropping image definition
type CropGravity int
const (
	CropCenter CropGravity = iota
	CropNorth
	CropEast
	CropSouth
	CropWest
	CropAuto
)
type CropParam struct {
	Width   int         `json:"w,omitempty"`
	Height  int         `json:"h,omitempty"`
	Gravity CropGravity `json:"g"`
}

// About rotating image definitions
type RotateParam struct {
	Angle int `json:"a"`
}

// About resizing image definitions
type ResizeMode int
type ResizeParam struct {
	Width  int        `json:"w,omitempty"`
	Height int        `json:"h,omitempty"`
	Mode   ResizeMode `json:"m"`
}

// On the definition of text watermarking
type WaterMarkParam struct {
	Dpi     int     `json:"d,omitempty"`
	Opacity float64 `json:"p,omitempty"`
	Text    string  `json:"t"`
	Color   string  `json:"c"`
}

// On the definition of image watermarking
 type WaterMarkImageParam struct {
	Left    int     `json:"l"`
	Top     int     `json:"t"`
	Opacity float64 `json:"p,omitempty"`
	Url     string  `json:"u"`
}

// About image format conversion definitions
type FormatParam struct {
	Type string `json:"t"`
}

```

Create configuration from Access Key and Initialize the QingStor service with a configuration

```
// Initialize the QingStor service with a configuration
config, _ := config.New("ACCESS_KEY_ID", "SECRET_ACCESS_KEY")
service, _ := qs.Init(config)
```

Initialize a QingStor bucket
```
bucket, _ := service.Bucket("bucketName", "zoneID")
```
Initialize a image 
```
image := helper.InitImage(bucket, "imageName")
```
Crop the image
```
image = image.Crop(&helpers.CropParam{...param_list...})
```
Rotate the image
```
image = image.Rotate(&helpers.RotateParam{...param_list...})
```
Resize the image
```
image = image.Resize(&helpers.ResizeParam{...param_list...})
```
Watermark the image
```
image = image.WaterMark(&helpers.WaterMarkParam{...param_list...})
```
WaterMarkImage the image
```
image = image.WaterMarkImage(&helpers.WaterMarkImageParam{...param_list...})
```
Format the image

```
image = image.Format(&helpers.Format{...param_list...})
```

Include a complete example, but the code needs to fill in your own information
```
package main

import (
	"github.com/yunify/qingstor-sdk-go/config"
	"github.com/yunify/qingstor-sdk-go/helpers"
	qs "github.com/yunify/qingstor-sdk-go/service"
	"log"
)

func main() {
	// Load your configuration
	// Replace here with your key pair
	config, err := config.New("ACCESS_KEY_ID", "SECRET_ACCESS_KEY")
	checkErr(err)

	// Initialize QingStror Service
	service, err := qs.Init(config)
	checkErr(err)

	// Initialize Bucket
	// Replace here with your bucketName and zoneID
	bucket, err := service.Bucket("yourBucketName", "zoneID")
	checkErr(err)

	// Initialize Image
	// Replace here with your your ImageName
	image := helpers.InitImage(bucket, "yourImageName")
	checkErr(image.Err)

	// Because 0 is an invalid parameter, default not modify
	image = image.Crop(&helpers.CropParam{Width: 0})
	checkErr(image.Err)
	testOutput(image.ImageOutput)

	// Rotate the image 90 angles
	image = image.Rotate(&helpers.RotateParam{Angle: 90})
	checkErr(image.Err)
	testOutput(image.ImageOutput)

	// Text watermark, Watermark text content, encoded by base64.
	image = image.WaterMark(&helpers.WaterMarkParam{
		Text: "5rC05Y2w5paH5a2X",
	})
	checkErr(image.Err)
	testOutput(image.ImageOutput)

	// Image watermark, Watermark image url encoded by base64.
	image = image.WaterMarkImage(&helpers.WaterMarkImageParam{
		Url: "aHR0cHM6Ly9wZWszYS5xaW5nc3Rvci5jb20vaW1nLWRvYy1lZy9xaW5jbG91ZC5wbmc=",
	})
	checkErr(image.Err)
	testOutput(image.ImageOutput)

	// Reszie the image with width 300px and height 400 px
	image = image.Resize(&helpers.ResizeParam{
		Width:  300,
		Height: 400,
	})
	checkErr(image.Err)
	testOutput(image.ImageOutput)

	// Swap format to jpeg
	image = image.Format(&helpers.FormatParam{
		Type: "jpeg",
	})
	checkErr(image.Err)
	testOutput(image.ImageOutput)
}

// *qs.ImageProcessOutput: github.com/yunify/qingstor-sdk-go/service/object.go
func testOutput(out *qs.ImageProcessOutput) {
	log.Println(*out.StatusCode)
	log.Println(*out.RequestID)
	log.Println(out.Body)
	log.Println(*out.ContentLength)

}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
```
