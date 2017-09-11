package main

import (
	"log"

	"github.com/zhaohuXing/Go/os/basic"
)

func main() {
	file, err := basic.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Write done, so first commented out temporarily
	//	for i := 0; i < 10; i++ {
	//		basic.Write([]byte("Hello World\n"), file)
	//	}

	err = basic.SingleThreadCopy(file, "demo_single_thread_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println(basic.ModeDir)
	//	fmt.Println(basic.ModeAppend)
	//	fmt.Println(basic.ModeSymlink)
	//	fmt.Println(basic.ModeNamedPipe)
	//	fmt.Println(basic.ModeSocket)
	//	fmt.Println(basic.ModeDevice)
	//	fmt.Println(basic.ModeType)
	//
	//	// If the file doesn't exist, create it, or append to the file
	//	// to use OpenFile
	//	f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	// to use Write
	//	if _, err := f.Write([]byte("appended some data\n")); err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fileinfo, err := f.Stat()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	log.Printf("mode: %d\n", fileinfo.Mode())
	//	f.WriteAt([]byte("leave qingstor!"), 100000)
	//	// to use Close
	//	if err := f.Close(); err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	// to use Open
	//	f1, err := os.Open("access.log")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	data := make([]byte, fileinfo.Size())
	//	// to use Open
	//	count, err := f1.Read(data)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	log.Printf("read %d bytes: %q\n", count, data[:count])
	//
	//	// to use seek
	//	ret, err := f1.Seek(10, os.SEEK_CUR)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	log.Println(ret)
	//	f1.Write([]byte("good good study"))
	//	basic.Open()
}
