package api

import (
	"archive/zip"
	"io"
	"os"
)

type Zipp struct{

}
func(z *Zipp) Archive(names []string, readers ...io.Reader)(outR io.Reader,err error){
	result, err := os.Create("./public/result.zip")
	
	if err != nil {
        panic(err)
    }
	defer result.Close()
    zipW := zip.NewWriter(result)
	for i, f := range names{
		w1, err := zipW.Create(f)
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(w1, readers[i])
		if err != nil {
			panic(err)
		}
	}
	zipW.Close()
	return result,nil

	// result, err := os.Create("./public/result.zip")

	// if err != nil {
	// 	panic(err)
	// }

	// defer result.Close()
	// zipW := zip.NewWriter(result)
	// for _, f := range names {
	// 	file, err := os.Open(f)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer file.Close()
	// 	w1, err := zipW.Create(f)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	_, err = io.Copy(w1, file)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// zipW.Close()
	// return result, nil
}