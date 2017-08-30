package log

import (
	"os"
	"testing"
)

func TestPrintf(t *testing.T) {
	logger := New(os.Stderr, "joe ~ ", LstdFlags)
	logger.Printf("%s %s", "Hello", "World")
	logger.Print("Hello", "World")
	//logger.Fatal("Hello Java")
	logger.Println("Hello", "Go")
	//logger.Fatalf("%s %s", "Hello", "Java")
	//logger.Fatalln("Hello", "World")
	//logger.Panic("Hello, World")
	//logger.Panicf("%s%s", "Hi", "Go")
	//logger.Panicln("Hi, Go")
	logger.Print(logger.Flags())
	logger.SetFlags(5)
	logger.Print(logger.Flags())
	logger.Print(logger.Prefix())
	logger.SetPrefix("go ~")
	logger.Print(logger.Prefix())
	// SetOutput("")

	Flags()
	SetFlags(3)
	logger.Print(Prefix())
	SetPrefix("Go Language")
	Print("Go Language")
	Printf("%s %s", "Go", "Language")
	Println("Learn Go Language")
	//Fatal("Learn Go Language")
	//Fatalf("%s", "Learn Go Language")
	//Fatalln("Learn Go Language")
	//Panic("Learn Go Language")
}
