package pkg

import (
	"os"
)

func ReadFile(path string) (string, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func WriteFile(dirName, fileName, data string) error {

	if data != "" {

		directory := dirName + "/_result"
		_ = os.Mkdir(directory, os.ModePerm)
		_ = os.Remove(directory + "/" + fileName)

		f, err := os.Create(directory + "/" + fileName)

		if err != nil {
			return err
		}

		defer f.Close()

		_, err2 := f.WriteString(data)

		if err2 != nil {
			return err
		}
	}
	return nil

}
