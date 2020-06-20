package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ValidationName(value string) (string, error) {
	if len(value) >= 5 && len(value) <= 40 {
		return value, nil
	} else {
		return "", fmt.Errorf("mins 5 caracter, maks 40 caracter")
	}
}

func ValidationNull(value string) (string, error) {
	if value != "" {
		return value, nil
	}
	return "", fmt.Errorf("column cannot be empty")
}

func ValidationNumber(value string) (int, error) {
	integer, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("column must be number")
	}
	return integer, nil
}

func ValidationDescription(value string) (string, error) {
	if len(value) >= 30 {
		return value, nil
	} else {
		return "", fmt.Errorf("mins 30 caracter")
	}
}

func ValidationFileSize(value int64) (int64, error) {
	if value >= 2003000 {
		return 0, fmt.Errorf("Image size cannot be more than 2 MB")
	}
	return value, nil
}

func ValidationFormatFile(value string) (string, error) {

	formatFile1 := strings.HasSuffix(value, ".png")
	formatFile2 := strings.HasSuffix(value, ".jpg")
	formatFile3 := strings.HasSuffix(value, ".jpeg")

	if formatFile1 == true || formatFile2 == true || formatFile3 == true {
		return value, nil
	}
	return "", fmt.Errorf("Unrecognized file format")
}

func ValidationRollbackImage(values []string) {
	for i := 0; i < len(values); i++ {
		err := os.Remove(values[i])
		if err != nil {
			fmt.Println(err)
		}
	}
}
