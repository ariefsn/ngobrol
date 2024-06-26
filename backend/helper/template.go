package helper

import (
	"os"
	"path/filepath"

	"github.com/ariefsn/ngobrol/constants"
)

func GetTemplate(tmpKey constants.MailTemplate) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	tmpPath := filepath.Join(wd, "templates", string(tmpKey))

	tmpContent, err := os.ReadFile(tmpPath)
	if err != nil {
		return "", err
	}

	return string(tmpContent), nil
}
