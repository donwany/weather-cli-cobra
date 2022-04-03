package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"weather/utils"
)
var cityName string
// cityCmd represents the city command
var cityCmd = &cobra.Command{
	Use:   "city",
	Short: "Get temperature of a city",
	Long: `This command fetches the temperature from the openWeatherMap API.`,
	Run: func(cmd *cobra.Command, args []string) {
			var city string
			if cityName != ""{
				city = cityName
			}else {
				city = args[0] // check for validation
			}
			fmt.Printf("Temperature at %s is %.f degrees celcius", city, utils.GetCity(city))
	},
}

func init() {
	rootCmd.AddCommand(cityCmd)
	cityCmd.PersistentFlags().StringVarP(&cityName, "city", "c", "",
		"find temp by city name (required)")
	cityCmd.MarkPersistentFlagRequired("city")
}

