package commands

import (
	"github.com/spf13/cobra"
	"go-api-template/services/country_info"
	"log"
)

var CapitalsCmd = &cobra.Command{
	Use:   "capitals",
	Short: "Start a country capitals guessing game",
	Run: func(cmd *cobra.Command, args []string) {
		c := country_info.NewCountryInfoClient()
		err := c.GetCapital("BR")
		if err != nil {
			log.Println("Ooops", err)
		}
	},
}