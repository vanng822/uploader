package mongodb

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"os"
	"testing"
	"io/ioutil"
	//"gopkg.in/mgo.v2"
)


var testHost = "localhost:27017"
var testDbname = "uploader_unittest_db"

var testUrl = fmt.Sprintf("%s/%s", testHost, testDbname)
var testPrefix = "uploader_test_collection"

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
	session := getSession(testUrl)
	defer session.Close()
	
	retCode := m.Run()
	session.DB(testDbname).DropDatabase()
	
	fmt.Println("Test ending")
	os.Exit(retCode)
}

