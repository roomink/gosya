package gosya

import (
	"io/ioutil"
	"launchpad.net/goyaml"
	"os"
)

func Merge(settings interface{}, path string, env string) (err error) {
	list_files := []string{
		"settings.yml",
		"settings/" + env + ".yml",
		"settings.local.yml",
		"settings/" + env + ".local.yml",
	}

	for _, file := range list_files {
		filename := path + "/" + file
		if _, err := os.Stat(filename); err != nil {
			if os.IsNotExist(err) == false {
				return err
			}
		} else {
			body, err := ioutil.ReadFile(filename)
			if err != nil {
				return err
			}
			err = goyaml.Unmarshal(body, settings)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
