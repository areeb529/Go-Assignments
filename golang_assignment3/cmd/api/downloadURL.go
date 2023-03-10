package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var WebID = 0

type Web struct {
	Id int
}

func (v Web) NewDownloader() {
	WebID++
}

func (v *Web) Download(url string) (io.Reader, error) {

	// get the data
	resp, err := http.Get(url)
	if err != nil {
		return resp.Body, err
	}

	//defer resp.Body.Close()

	// Create the file
	filepath := "f"
	out, err := os.Create(fmt.Sprintf("./public/"+filepath+"%d", WebID))

	if err != nil {
		fmt.Println("Error creating file")
		return resp.Body, err
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	out.Seek(0, io.SeekStart)

	return out, err
}
