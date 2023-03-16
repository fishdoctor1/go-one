package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Photos []struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type Image struct {
	filePath string
	img      []byte
}

func main() {
	log.SetFlags(log.Ltime)
	dir := "myDownloadImage" + time.Now().Format("15_04_05")
	if _, err := os.Stat(dir); err != nil {
		os.Mkdir(dir, os.ModeDir)
	}
	photos := Photos{}
	err := getJson("https://jsonplaceholder.typicode.com/photos", &photos)
	fmt.Println(err)
	fmt.Println(len(photos)) //data ใน struct

	chImg := make(chan Image, len(photos))
	counter := sync.WaitGroup{}
	for i, v := range photos {
		counter.Add(1)
		v := v //ถ้าไม่มีบรรทัดนี้ go func จะ capture v ของ for เป็นตัวสุดท้ายเสมอ
		i := i
		go func() {
			defer counter.Done()
			img, err := downloadImage(v.ThumbnailURL)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("IND", i, "E")
			// fmt.Println(img) //byteCode
			format, err := decodeImage(img)
			if err != nil {
				log.Fatal(err)
			}
			absFileName := filepath.Join(dir, fmt.Sprintf("%d.%s", v.ID, format))
			// fileName := fmt.Sprintf("%d.%s", v.ID, format)
			chImg <- Image{filePath: absFileName, img: img}
			fmt.Println("GOOOOOOOOOOOOO")
		}()
	}

	go func() {
		counter.Wait() //รอ go routine ทั้งหมดจบ
		close(chImg)
	}()

	for v := range chImg {
		err = saveImage(v.filePath, v.img)
		if err != nil {
			log.Println(err)
		}
	}
	// for range photos { //loop แบบนี้อาจมีปัญหา กรณีโหลดมาได้ไม่ครบตาม length photos โปรแกรมจะค้างเพราะ channel รออยู่
	// 	v := <-chImg
	// 	fmt.Println("saveImage")
	// 	err = saveImage(v.filePath, v.img)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }

}

func saveImage(fileName string, img []byte) error {
	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("saveImage - cannot create file : %v", err)
	}
	defer f.Close()
	_, err = io.Copy(f, bytes.NewReader(img))
	if err != nil {
		return fmt.Errorf("saveImage - cannot save file : %v", err)
	}
	return nil
}

func decodeImage(img []byte) (string, error) {
	// fmt.Println("deCodeImage")
	_, format, _ := image.Decode(bytes.NewReader(img))
	return format, nil
}

func downloadImage(url string) ([]byte, error) {
	fmt.Println("DOWNLOAD")
	errMsg := func(err error) error {
		return fmt.Errorf("downloadImage : %v", err)
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, errMsg(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errMsg(err)
	}
	res.Body.Close()
	fmt.Println("CLOSEEEEEEEEEEEEEE")
	return body, nil
}

func getJson(url string, structType interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	switch v := structType.(type) {
	case *Photos:
		decoder := json.NewDecoder(res.Body)
		photos := structType.(*Photos)
		decoder.Decode(photos) //ได้ data เข้า struct
		return nil
	default:
		return fmt.Errorf("getJson:not support type %v", v)
	}
}
