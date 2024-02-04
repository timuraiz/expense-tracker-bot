package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string

	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	Driver   string

	Messages
}

type Messages struct {
	Responses
	Errors
}

type Responses struct {
	Start             string `mapstructure:"start"`
	AlreadyAuthorized string `mapstructure:"already_authorized"`
	UnknownCommand    string `mapstructure:"unknown_command"`
}

type Errors struct {
	Default      string `mapstructure:"default"`
	UnableToSave string `mapstructure:"unable_to_save"`
}

func Init() (*Config, error) {
	if err := setUpViper(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := fromEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.response", &cfg.Messages.Responses); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.error", &cfg.Messages.Errors); err != nil {
		return err
	}

	return nil
}

func fromEnv(cfg *Config) error {
	if err := viper.BindEnv("token"); err != nil {
		return err
	}
	cfg.TelegramToken = viper.GetString("token")

	if err := viper.BindEnv("host"); err != nil {
		return err
	}
	cfg.Host = viper.GetString("host")

	if err := viper.BindEnv("port"); err != nil {
		return err
	}
	cfg.Port = viper.GetInt("port")

	if err := viper.BindEnv("user"); err != nil {
		return err
	}
	cfg.User = viper.GetString("user")

	if err := viper.BindEnv("password"); err != nil {
		return err
	}
	cfg.Password = viper.GetString("password")

	if err := viper.BindEnv("dbname"); err != nil {
		return err
	}
	cfg.DBName = viper.GetString("dbname")

	if err := viper.BindEnv("driver"); err != nil {
		return err
	}
	cfg.Driver = viper.GetString("driver")

	return nil
}

func setUpViper() error {
	// Set the file name of the configurations file
	viper.SetConfigName("main")
	// Set the path to look for the configurations file
	viper.AddConfigPath("configs")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	// Set the type of the configuration file
	viper.SetConfigType("yaml")

	// Read YAML config first
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
		return err
	}

	// Explicitly specify to use the .env file
	viper.SetConfigFile(".env")
	viper.SetConfigType("env") // if your .env file has specific format, you can specify it here
	viper.AddConfigPath(".")   // the path to look for the .env file

	// Attempt to read the .env config into Viper, overriding any YAML values
	// This is not mandatory and can be removed if you do not wish to use a .env file
	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// .env Config file not found; ignore error if desired
			fmt.Println("No .env file found")
		} else {
			// .env Config file was found but another error was produced
			return err
		}
	}

	return nil
}
