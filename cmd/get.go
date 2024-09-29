package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
)

var (
	city  string
	units string
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the current weather for a city",
	Long:  "Fetches the current weather for a specified city. If no city is provided, it uses the default city from the config file.",
	Run:   getWeather,
}

func init() {
	// Register flags
	getCmd.Flags().StringVarP(&city, "city", "c", "", "City name")
	getCmd.Flags().StringVarP(&units, "units", "u", "", "Units of measurement (metric or imperial)")

	// Add the get command to the root command
	rootCmd.AddCommand(getCmd)
}

func getWeather(cmd *cobra.Command, args []string) {
	apiKey := viper.Get("api_key").(string)
	if apiKey == "" {
		fmt.Println("Error: OpenWeatherMap API key not found. Please run 'forecast setup' to configure your API key.")
		return
	}

	if city == "" {
		city = viper.Get("default_city").(string)
		fmt.Printf("No city specified. Using preconfigured default city: %s\n", city)
	}

	if units == "" {
		units = viper.Get("units").(string)
		fmt.Printf("No units specified. Using preconfigured default units: %s\n", units)
	}

	// Build the API request URL
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=%s&appid=%s", city, units, apiKey)

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("City '%s' not found.\n", city)
		return
	} else if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		return
	}
	if err != nil {
		fmt.Printf("Error fetching weather data: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		return
	}

	// Parse the JSON response
	var data map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&data); err != nil {
		fmt.Printf("Error decoding JSON response: %v\n", err)
		return
	}

	// Display the weather information
	displayWeather(data)
}

func displayWeather(data map[string]interface{}) {
	name := data["name"].(string)
	sys := data["sys"].(map[string]interface{})
	country := sys["country"].(string)
	main := data["main"].(map[string]interface{})
	weather := data["weather"].([]interface{})[0].(map[string]interface{})

	fmt.Printf("\n%s weather in %s, %s:\n", color.CyanString("Current"), name, country)
	fmt.Printf("%s: %.2f°\n", color.YellowString("Temperature"), main["temp"].(float64))
	fmt.Printf("%s: %.2f°\n", color.YellowString("Feels Like"), main["feels_like"].(float64))
	fmt.Printf("%s: %s\n\n", color.GreenString("Weather"), weather["description"].(string))
}
