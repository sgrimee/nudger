// Load a config from a json file
// borrowed from github.com/codegangsta/gin

package nudger

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigType struct {
	ItemsDir  string `json:"items_dir"`
	NudgeCmd  string `json:"nudge_cmd"`
	NudgeArgs string `json:"nudge_args"`
}

func LoadConfig(path string) (*ConfigType, error) {
	configFile, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Unable to read configuration file %s", path)
	}

	config := new(ConfigType)

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse configuration file %s", path)
	}

	return config, nil
}
