package services

import (
	"os"

	"gopkg.in/yaml.v3"
)

// JiracliAppInfo represents metadata for this project (mostly generated from CI pipeline).
type JiracliAppInfo struct {
	Version string `yaml:"version"`
}

// Read unmarshalls the `app-info.yml` file into JiracliAppInfo struct. The `app-info.yml`
// is located at `src/main/app-info.yml`. However since this project relies heavily on Docker
// the actual path to the file in the respective container (dev container, go wrapper container
// and jiracli image) is `/app/app-info.yml`.
func (info *JiracliAppInfo) Read() (JiracliAppInfo, error) {
	path := "/app/app-info.yml"

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return JiracliAppInfo{}, err
	}

	err = yaml.Unmarshal(yamlFile, info)
	if err != nil {
		return JiracliAppInfo{}, err
	}

	return *info, nil
}
