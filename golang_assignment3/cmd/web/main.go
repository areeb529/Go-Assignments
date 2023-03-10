package main

import (
	"fmt"
	"log"

	api "github.com/areeb529/downloadAndZip/cmd/api"
)

func main() {

	fileUrl := "https://www.jungleerummy.com/blog/assets/front/img/blog/bsdy4m1efq9i8md9m6gzelymija995khb9kwinei8kjhewbsjjmain.jpg"
	webFile := &api.Web{}
	webFile.NewDownloader()
	webFile.Id = api.WebID
	r1, err := webFile.Download(fileUrl)
	if err != nil {
		fmt.Printf("got error while downloading file with id %d\n", webFile.Id)
		log.Fatal(err)
	}

	fileUrl = "https://humornama.com/wp-content/uploads/2020/10/Miracle-Miracle-meme-template-of-welcome-movie-1024x576.jpg"
	webFile = &api.Web{}
	webFile.NewDownloader()
	webFile.Id = api.WebID
	r2, err := webFile.Download(fileUrl)
	if err != nil {
		fmt.Printf("got error while downloading file with id %d\n", webFile.Id)
	}

	f := &api.FileSystemManager{}
	fileName := "./public/akshay-meme.jpeg"
	_, err = f.Download(fileName)
	if err != nil {
		fmt.Printf("got error while downloading file %s\n", fileName)
	}

	zipper := &api.Zipp{}
	_, err = zipper.Archive([]string{"f1", "f2"}, r1, r2)
	if err != nil {
		panic(err)
	}

}
