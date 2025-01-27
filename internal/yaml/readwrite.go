package yaml

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/aceberg/ForAuth/internal/check"
)

// Read - read .yaml file to struct
func Read(path string) map[string]string {

	file, err := os.ReadFile(path)
	check.IfError(err)

	var items map[string]string
	err = yaml.Unmarshal(file, &items)
	check.IfError(err)

	return items
}

// Write - write struct to  .yaml file
func Write(path string, items map[string]string) {

	yamlData, err := yaml.Marshal(&items)
	check.IfError(err)

	err = os.WriteFile(path, yamlData, 0644)
	check.IfError(err)

	log.Println("INFO: writing new tagrets to", path)
}
