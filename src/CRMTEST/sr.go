package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {

}

func PorTest(c echo.Context) (r interface{}, err error) {
	log.Debug("Start Process")
	//startProcess := time.Now()
	log.GenTranID()

	log.Debug("ssssssssssssssssssssssssssss")
	log.Debug(config.GetString("sr.category.group.allareas"))
	// fmt.Println("fss", config.Get("sr.category.group.allareas"))
	// fmt.Println("fzzz", config.GetStringMapString("sr.category.group.allareas"))
	// fmt.Println("sssss", config.GetStringMapStringSlice("sr.category.group.allareas")) // value จะเป็น array
	// fmt.Println("ddddd", config.GetStringMap("sr.category.group.allareas"))
	// fmt.Println("wwwww", config.GetStringSlice("sr.category.group.allareas")) //ไม่ออก
	// fmt.Println("dfwese", config.GetString("sr.category.group.allareas")) //ไม่ออก
	fmt.Println(config.GetString("sr.subcategory." + "subcate009"))
	ret := porTestModel("ปัญหา Voice Call")
	fmt.Printf("%T ,%v", ret, ret)
	// for key, value := range config.GetStringMap("sr.category.group.allareas") {
	// 	fmt.Println(fmt.Sprintf("%s : %s", key, value))
	// }
	fmt.Println("-------------------------------------------------------------------------")
	// for key, value := range config.GetStringMapString("sr.category.group.allareas") {
	// 	fmt.Println(fmt.Sprintf("%s : %s", key, value))
	// 	if string(value) == "IT" {
	// 		fmt.Println("ITTTTTTTTTTTTTTTTTT")
	// 		break
	// 	}
	// }

	// configg := config.GetStringMapString("sr.category.group.allareas")
	// for i := 0; i < len(configg); i++ {
	// 	fmt.Printf("%d:%s %s \n", i, configg["key"], configg["value"])
	// 	if string(configg["value"]) == "IT" {
	// 		fmt.Println("ITTTTTTTTTTTTTTTTTT")
	// 		break
	// 	}
	// }

	log.Debug("PorTest111111111111111111")
	// log.Debug(req)
	// req.CallNumber = c.Request().Header.Get("callID")
	// req.CallChannel = c.Request().Header.Get("callChannel")

	r = Response{
		SrNumber:      "",
		OmxTrackingID: "",
		ActivityID:    "",
	}

	return
}

func porTestModel(subcate string) (resp bool) {
	resp = false
	respConfig := config.GetString("sr.subcategory." + subcate)
	if respConfig == "TRUE" {
		resp = true
	}
	return
}
