package pkg

import "os"

func ReadTextFiles(path string) (string, error) {
	file, errOnRead := os.ReadFile(path)

	if errOnRead != nil {
		return "", errOnRead
	}

	buffer := string(file)
	return buffer, nil
}
