package php

import (
	"bufio"
	"bytes"
	"os"
)

// assertIsPHPScript checks that the composer file is indeed a phar/PHP script (not a .bat file)
func assertIsPHPScript(path string) bool {
	if path == "" {
		return false
	}
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	byteSlice, _, err := reader.ReadLine()
	if err != nil {
		return false
	}

	if bytes.Equal(byteSlice, []byte("<?php")) {
		return true
	}

	return bytes.HasPrefix(byteSlice, []byte("#!/")) && bytes.HasSuffix(byteSlice, []byte("php"))
}
