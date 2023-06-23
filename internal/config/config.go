package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

const JwtSecret = "AboutVaccine"

const SuccessStatus = 1
const FailureStatus = -1

const UserClaimCookie = "user_token"

type AllConfig struct {
	Server   *Server   `yaml:"server"`
	Database *Database `yaml:"database"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Database struct {
	Driver     string `yaml:"driver"`
	Connection string `yaml:"connection"`
}

func SaveConfig(conf *AllConfig) error {
	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}
	// 检查文件是否存在
	_, err = os.Stat("/config.yaml")
	if os.IsNotExist(err) {
		file, err := os.Create("/config.yaml")
		defer file.Close()
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	err = ioutil.WriteFile("/config.yaml", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadConfig() (*AllConfig, error) {
	conf := &AllConfig{}
	data, err := ioutil.ReadFile("/config.yaml")
	if err != nil {
		return conf, err
	}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
