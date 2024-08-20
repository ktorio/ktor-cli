package config

import (
	"encoding/json"
	"io"
	"os"
)

var data map[string]string

func init() {
	data = make(map[string]string)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	f, err := os.Open(ktorConfigPath(homeDir))
	if err != nil {
		return
	}

	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return
	}

	_ = json.Unmarshal(b, &data)
}

func GetValue(key string) (string, bool) {
	v, ok := data[key]
	return v, ok
}

func SetValue(key string, value string) {
	data[key] = value
}

func Commit() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(ktorConfigPath(homeDir), b, 0755)
	if err != nil {
		return err
	}

	return nil
}
