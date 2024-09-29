# Forecast

A simple command-line tool written in Go for fetching weather forecasts for specified locations.

## Features
* Fetch current weather information for any city.	
* Supports metric and imperial units.
* Interactive setup to configure default settings. 
* Shell auto-completion for Bash, Zsh, Fish, and PowerShell.

## Installation

### Prerequisites
* Go (version 1.17 or later)

### Build from Source
```bash
git clone https://github.com/yourusername/forecast.git
cd forecast
go build -o forecast
```
### Add to Your PATH

Move the executable to a directory in your PATH:
```bash
sudo mv forecast /usr/local/bin/
```

## Usage

### Initial Setup

Run the interactive setup to configure your API key and preferences:
```bash
forecast setup
```

### Fetch Weather Data

Get the current weather for the default city:
```bash
forecast get
```

### Get the weather for a specific city:
```bash
forecast get -city "New York"
```

## Configuration

The configuration file is located at ~/.forecast/config.yaml. You can edit it manually to change settings:
```yaml
units: metric
default_city: London
api_key: your_openweathermap_api_key
```

### Obtaining an API Key

Register for a free API key at [OpenWeatherMap](https://openweathermap.org/api).

## License

This project is licensed under the MIT License.