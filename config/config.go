package config

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

func InitConfig() {
	viper.SetConfigName("config")          // name of config file (without extension)
	viper.SetConfigType("yaml")            // or json, toml, etc.
	viper.AddConfigPath("$HOME/.forecast") // path to look for the config file

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; create default config
			err := CreateDefaultConfig()
			if err != nil {
				fmt.Printf("Error creating default config file: %v\n", err)
			}
		} else {
			// Config file was found but another error was produced
			fmt.Printf("Error reading config file: %v\n", err)
		}
	}
}

func CreateDefaultConfig() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Let's set up your default configuration. You can exit at any time by typing 'exit'.")

	// Prompt for units
	var unitsInput string
	for {
		var existingConfigValue = viper.Get("units").(string)
		if existingConfigValue == "" {
			existingConfigValue = "metric"
		}
		fmt.Printf("Choose units (metric/imperial) [%s]: ", existingConfigValue)
		unitsInput, _ = reader.ReadString('\n')
		unitsInput = strings.TrimSpace(unitsInput)

		if unitsInput == "" {
			unitsInput = existingConfigValue
			break
		} else if strings.EqualFold(unitsInput, "exit") {
			fmt.Println("Setup exited.")
			os.Exit(0)
		} else if unitsInput == "metric" || unitsInput == "imperial" {
			break
		} else {
			fmt.Println("Invalid input. Please enter 'metric' or 'imperial'.")
		}
	}

	// Prompt for default city
	var cityInput string
	for {
		var existingConfigValue = viper.Get("default_city").(string)
		if existingConfigValue == "" {
			existingConfigValue = "London"
		}
		fmt.Printf("Enter your default city [%s]: ", existingConfigValue)

		cityInput, _ = reader.ReadString('\n')
		cityInput = strings.TrimSpace(cityInput)

		if cityInput == "" {
			cityInput = existingConfigValue
			break
		} else if strings.EqualFold(cityInput, "exit") {
			fmt.Println("Setup exited.")
			os.Exit(0)
		} else {
			break
		}
	}

	// Prompt for API key
	var apiKeyInput string
	for {
		var existingApiKey = viper.Get("api_key").(string)
		if existingApiKey == "" {
			fmt.Print("Enter your OpenWeatherMap API key: ")
		} else {
			fmt.Printf("Enter your OpenWeatherMap API key [%.6s....]: ", existingApiKey)
		}

		apiKeyInput, _ = reader.ReadString('\n')
		apiKeyInput = strings.TrimSpace(apiKeyInput)

		if strings.EqualFold(apiKeyInput, "exit") {
			fmt.Println("Setup exited.")
			os.Exit(0)
		} else if apiKeyInput == "" && existingApiKey != "" {
			fmt.Printf("Re-using existing API key %.6s....", existingApiKey)
			apiKeyInput = existingApiKey
			break
		} else if apiKeyInput == "" {
			fmt.Println("API key cannot be empty. Please enter a valid API key.")
		} else {
			break
		}
	}

	// Set configuration values
	viper.Set("units", unitsInput)
	viper.Set("default_city", cityInput)
	viper.Set("api_key", apiKeyInput)

	// Determine the config file path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configDir := filepath.Join(homeDir, ".forecast")

	// Create the directory if it doesn't exist
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err = os.Mkdir(configDir, 0755)
		if err != nil {
			return err
		}
	}

	// Define the full path to the config file
	configFile := filepath.Join(configDir, "config.yaml")

	// Write the config file
	err = viper.WriteConfigAs(configFile)
	if err != nil {
		return err
	}

	fmt.Printf("Configuration saved to %s\n", configFile)
	return nil
}
