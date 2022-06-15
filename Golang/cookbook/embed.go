package cookbook

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"os"
	"path"

	"github.com/wwqdrh/logger"
)

//go:embed templates
var templates embed.FS

func init() {
	extractTemplates("templates")
}

func extractTemplates(prePath string) error {
	_, err := os.Stat(prePath)
	if !os.IsNotExist(err) {
		return nil
	}

	err = os.MkdirAll(prePath, 0o755)
	if err != nil {
		return err
	}
	items, err := fs.ReadDir(templates, prePath)
	if err != nil {
		return err
	}

	for _, item := range items {
		fullpath := path.Join(prePath, item.Name())
		if item.IsDir() {
			extractTemplates(fullpath)
		} else {
			f, err := templates.Open(fullpath)
			if err != nil {
				logger.DefaultLogger.Error(err.Error())
				continue
			}
			data, err := ioutil.ReadAll(f)
			if err != nil {
				logger.DefaultLogger.Error(err.Error())
				continue
			}
			os.WriteFile(fullpath, data, 0o755)
		}
	}
	return nil
}
