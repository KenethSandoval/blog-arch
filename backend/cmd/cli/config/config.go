package config

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const AppDir = "."

const SendFileName = "data.yml"

type DataSend struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Role     string `yaml:"role"`
}

type ConfigParser struct{}

type Data struct {
	Body DataSend `yaml:"data"`
}

func (parser ConfigParser) getDefaultData() Data {
	return Data{
		DataSend{
			Username: "",
			Password: "",
			Role:     "",
		},
	}
}

func (parser ConfigParser) getDefaultDataYamlContent() string {
	defaultConfig := parser.getDefaultData()
	yaml, _ := yaml.Marshal(defaultConfig)

	return string(yaml)
}

func (parser ConfigParser) writeDefaultConfigContents(newConfigFile *os.File) error {
	_, err := newConfigFile.WriteString(parser.getDefaultDataYamlContent())

	if err != nil {
		return err
	}

	return nil
}

func (parser ConfigParser) createSendFileIfMissing(sendFilePath string) error {
	if _, err := os.Stat(sendFilePath); errors.Is(err, os.ErrNotExist) {
		newSendFile, err := os.OpenFile(sendFilePath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			return err
		}

		defer newSendFile.Close()
		return parser.writeDefaultConfigContents(newSendFile)
	}
	return nil
}

func (parser ConfigParser) getConfigFileOrCreateIfMissing() (*string, error) {
	var err error

	sendFilePath := filepath.Join(AppDir, SendFileName)
	err = parser.createSendFileIfMissing(sendFilePath)

	if err != nil {
		return nil, err
	}
	return &sendFilePath, nil
}

func (parser ConfigParser) readSendFile(path string) (Data, error) {
	dataConfig := parser.getDefaultData()
	data, err := os.ReadFile(path)
	if err != nil {
		return dataConfig, err
	}

	err = yaml.Unmarshal((data), &dataConfig)

	return dataConfig, err
}

func initParser() ConfigParser {
	return ConfigParser{}
}

func ParserConfig() (Data, error) {
	var data Data
	var err error

	parser := initParser()

	configFilePath, err := parser.getConfigFileOrCreateIfMissing()
	if err != nil {
		return data, err
	}

	config, err := parser.readSendFile(*configFilePath)

	if err != nil {
		return data, err
	}

	return config, nil
}
