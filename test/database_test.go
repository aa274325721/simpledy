package test

import (
	"fmt"
	"simpledy/driver"
	"simpledy/repository"
	"testing"
)

//var user = model.User{
//	Username: "testUser",
//	Password: "123456",
//}

func TestDatabase(t *testing.T) {
	driver.InitTable()
	//driver.InitData()
	//fmt.Println(db.RowsAffected)
}

func TestCustomDb(t *testing.T) {
	driver.InitTable()
}
func TestFindNewestVideoTime(t *testing.T) {
	video, num := repository.FindLatestVideo()
	fmt.Println("found :", num)
	fmt.Println(video.CreatedAt)
}
