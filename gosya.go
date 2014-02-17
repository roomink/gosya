package gosya

import (
	"io/ioutil"
	"launchpad.net/goyaml"
	"log"
	"os"
)

func Basta() int {
	log.Println("Begin settings")
	return 1
}

func Merge(settings interface{}, path string, env string) (err error) {
	// goyaml.Unmarshal([]byte("advert: 11"), settings)

	// env_file := strings.Join([]string{"settings/", env, ".yml"}, "")
	list_files := []string{"settings.yml", "settings/" + env + ".yml", "settings.local.yml"}
	// log.Printf("%#v", list_files)
	for _, file := range list_files {
		filename := path + "/" + file
		// log.Println(filename)
		if _, err := os.Stat(filename); err != nil {
			if os.IsNotExist(err) == false {
				return err
			}
		} else {
			body, err := ioutil.ReadFile(filename)
			if err != nil {
				return err
			}
			// log.Printf("%#v", string(body))
			err = goyaml.Unmarshal(body, settings)
			if err != nil {
				return err
			}
			// log.Printf("%#v", settings)
		}
	}

	return nil
}
