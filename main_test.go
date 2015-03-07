package uploader

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"os"
	"testing"
	"io/ioutil"
)

func testGetImageByte() []byte {
	fd, err := os.Open("./data/kth.jpg")
	if err != nil {
		panic(err)
	}

	imageData, err := ioutil.ReadAll(fd)
	
	if err != nil {
		panic(err)
	}
	if len(imageData) == 0{
		panic("No data")
	}
	return imageData
}

func testGetFilename() string {
	uu, _ := uuid.NewV4()
	return fmt.Sprintf("%s.jpg", uu.String())
}

func TestMain(m *testing.M) {
	fmt.Println("Test starting")
	retCode := m.Run()
	fmt.Println("Test ending")
	os.Exit(retCode)
}
