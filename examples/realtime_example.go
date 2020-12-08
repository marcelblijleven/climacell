package main

import (
	"fmt"
	"github.com/marcelblijleven/climacell"
	"os"
)

func main() {
	apikey := os.Getenv("CLIMACELL_API_KEY")
	c, err := climacell.NewClient(apikey, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	latitude := 52.369069057354665
	longitude := 4.896479268175967

	resp, err := c.Realtime(
		latitude,
		longitude,
		climacell.Si,
		climacell.Temperature,
		climacell.CloudBase,
		climacell.CloudCeiling,
		climacell.TreePollen,
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	temperature := resp.Temperature
	cloudBase := resp.CloudBase
	cloudCeiling := resp.CloudCeiling

	fmt.Printf("temperature: %v\ncloud_base: %v\ncloud_ceiling: %v\n", temperature, cloudBase, cloudCeiling)
}
