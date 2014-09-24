package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/alexcesaro/log"
	"github.com/alexcesaro/log/stdlog"
	"github.com/longnguyen11288/go-transmission/transmission"
)

//Config holds all config data
type Config struct {
	Log                log.Logger                       `json:"-"`
	TransmissionClient *transmission.TransmissionClient `json:"-"`
	Rss                []string                         `json:"rss"`
	Shows              []string                         `json:"shows"`
	Transmission       Tranmission                      `json:"tranmission"`
	OutputDir          string                           `json:"output_dir"`
	DownloadDir        string                           `json:"download_dir"`
}

//Tranmission client for handling torrents
type Tranmission struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//ParseConfigFile parses config file
func ParseConfigFile(file string) (*Config, error) {
	var config Config
	configByte, err := ioutil.ReadFile(file)
	if err != nil {
		return &config, err
	}
	err = json.Unmarshal(configByte, &config)
	if err != nil {
		return &config, err
	}

	client := transmission.New(config.Transmission.URL,
		config.Transmission.Username, config.Transmission.Password)
	config.TransmissionClient = &client
	logger := stdlog.GetFromFlags()
	config.Log = logger

	return &config, nil

}
