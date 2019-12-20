package util

import (
	"io/ioutil"
	"os"
	"path/filepath"

	toml "github.com/pelletier/go-toml"
)

// ParseTOML .
func ParseTOML(path string, in interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = toml.Unmarshal(b, in)
	if err != nil {
		return err
	}
	return nil
}

// GetModeConfig .
func ParseTOMLByMode(mode string, in interface{}) error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	path := filepath.Dir(exe) + "/conf/config." + mode + ".toml"
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = toml.Unmarshal(b, in)
	if err != nil {
		return err
	}
	return nil
}
