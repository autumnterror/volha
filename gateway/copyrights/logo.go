package copyrights

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Info() error {
	file, err := os.Open(filepath.Join(".", "copyrights", "logo.txt"))
	if err != nil {
		return err
	}

	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(data[:n]))
	}
	fmt.Println("\nCreated by \"BreeZy Innovations technologies Russia\". All rights secured. \n www.breezyinnovations.su")
	if err = file.Close(); err != nil {
		return err
	}

	return nil
}
