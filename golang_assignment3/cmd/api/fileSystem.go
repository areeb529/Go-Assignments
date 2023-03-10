package api

import (
	"fmt"
	"io"
	"os"
)

type FileSystemManager struct{

}

func (f* FileSystemManager)Download(sourceFilePath string)(io.Reader, error){
	file, err := os.Open(sourceFilePath)
	if err != nil{
		fmt.Println("Error opening source file")
		return file, err
	}
	destFilePath := fmt.Sprintf("./public/akshay-meme"+"copy" +".jpeg")
	out, err := os.Create(destFilePath)
	//defer out.Close()
	 if err != nil{
		fmt.Println("Error creating destination file")
		return file, err
	}

	_, err = io.Copy(out, file)
	return file, nil


}