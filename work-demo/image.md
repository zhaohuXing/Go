# QingStor Image Processing Usage Guide

For processing the image stored in QingStor by a variety of basic operations, such as format, crop, watermark and so on. Please see [QingStor Image API](https://docs.qingcloud.com/qingstor/data_process/image_process/index.html) 

## Usage
Before using the image service, you need to initialize the [Configuration](https://github.com/yunify/qingstor-sdk-go/blob/master/docs/configuration.md) and [QingStor Service](https://github.com/yunify/qingstor-sdk-go/blob/master/docs/qingstor_service_usage.md).

```
//Import the latest version API
import (
	"github.com/yunify/qingstor-sdk-go/config"
	"github.com/yunify/qingstor-sdk-go/helpers"
	qs "github.com/yunify/qingstor-sdk-go/service"
)
```

## Code Snippet

The definition of the structure of the image processing，located in `qingstor-sdk-go/helpers/image.go`.These structs are used as arguments to the operation.

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
	Width   int         `schema:"w,omitempty"`
	Height  int         `schema:"h,omitempty"`
	Gravity CropGravity `schema:"g"`
}

// About rotating image definitions
type RotateParam struct {
	Angle int `schema:"a"`
}

// About resizing image definitions
type ResizeMode int
type ResizeParam struct {
	Width  int        `schema:"w,omitempty"`
	Height int        `schema:"h,omitempty"`
	Mode   ResizeMode `schema:"m"`
}

// On the definition of text watermarking
type WaterMarkParam struct {
	Dpi     int     `schema:"d,omitempty"`
	Opacity float64 `schema:"p,omitempty"`
	Text    string  `schema:"t"`
	Color   string  `schema:"c"`
}

// On the definition of image watermarking
 type WaterMarkImageParam struct {
	Left    int     `schema:"l"`
	Top     int     `schema:"t"`
	Opacity float64 `schema:"p,omitempty"`
	Url     string  `schema:"u"`
}

// About image format conversion definitions
type FormatParam struct {
	Type string `schema:"t"`
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
image = image.Crop(&helpers.CropParam{...param_struct...})
```
Rotate the image
```
image = image.Rotate(&helpers.RotateParam{...param_struct...})
```
Resize the image
```
image = image.Resize(&helpers.ResizeParam{...param_struct...})
```
Watermark the image
```
image = image.WaterMark(&helpers.WaterMarkParam{...param_struct...})
```
WaterMarkImage the image
```
image = image.WaterMarkImage(&helpers.WaterMarkImageParam{...param_struct...})
```
Format the image

```
image = image.Format(&helpers.Format{...param_struct...})
```

Pipline model. The maximum number of operations in the pipeline is 10
```
image = image.Rotate(&helpers.RotateParam{
	... param_struct...	
}).Resize(&helpers.ResizeParam{
	... param_struct...
})
```

Use the original api to rotate the image 90 angles
```
angle := "rotate:a_90"
imageprocessoutput, err := bucket.imageprocess("yourimagename", &qs.imageprocessinput{
	action: &angle})
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

	// Pipline model
	// The maximum number of operations in the pipeline is 10
	image = image.Rotate(&helpers.RotateParam{
		Angle: 270,
	}).Resize(&helpers.ResizeParam{
		Width:  300,
		Height: 300,
	})
	checkErr(image.Err)
	testOutput(image.ImageOutput)

	// Use the original api to rotate the image 90 angles
	angle := "rotate:a_90"
	imageprocessoutput, err := bucket.imageprocess("yourimagename", &qs.imageprocessinput{
		action: &angle})
	checkErr(err)
	testOutput(imageProcessOutput)
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
