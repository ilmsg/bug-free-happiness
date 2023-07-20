package filedatabase

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type fileDatabase struct {
	file string
}

func (db fileDatabase) Set(key, value string) error {
	f, err := os.OpenFile(db.file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("%s:%s\n", key, value)); err != nil {
		return err
	}

	return nil
}

func (db fileDatabase) Get(key string) (string, error) {
	f, err := os.OpenFile(db.file, os.O_RDONLY, 0600)
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var last string
	for scanner.Scan() {
		row := scanner.Text()
		parts := strings.Split(row, ":")
		if len(parts) < 2 {
			return "", errors.New("invalid record")
		}
		if parts[0] == key {
			last = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return "", errors.Wrap(err, "scanner.Err")
	}

	if last != "" {
		return last, nil
	}

	return "", errors.New("not found")
}

func NewFileDatabase(path string) *fileDatabase {
	return &fileDatabase{path}
}
