package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

const (
	paramCity = "city"
)

func main() {
	app := cli.NewApp()
	app.Name = "weather"
	app.Usage = "Get current weather"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  paramCity,
			Value: "berlin",
			Usage: "city to get weather for",
		},
	}
	app.Action = func(c *cli.Context) {
		client := NewClient()
		city := c.String(paramCity)
		weather, err := client.Request(city)
		if err != nil {
			fmt.Printf("Error: %q", err)
		} else {
			fmt.Printf("Weather in %s: %.1f\u2103 \n", city, weather.Temperature)
		}
	}

	app.Run(os.Args)
}
