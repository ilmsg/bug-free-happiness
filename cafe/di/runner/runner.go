package runner

import (
	"errors"
	"io"

	"github.com/ilmsg/cafe/di/storage"
)

var (
	errUsage = errors.New(`usage:
	set <key> <value> Set specified key, and value
	get <key>         Get specified key`)
)

type runner struct {
	database storage.Storage
}

func (r *runner) Run(output io.StringWriter, args []string) error {
	if len(args) < 3 {
		return errUsage
	}
	switch args[1] {
	case "set":
		if len(args) < 4 {
			return errUsage
		}
		if err := r.database.Set(args[2], args[3]); err != nil {
			return err
		}
	case "get":
		v, err := r.database.Get(args[2])
		if err != nil {
			return err
		}
		output.WriteString(v + "\n")
	default:
		return errUsage
	}

	return nil
}

func NewRunner(db storage.Storage) *runner {
	return &runner{db}
}
