package commands

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"go-api-template/services/country_info"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var CapitalsCmd = &cobra.Command{
	Use:   "capitals",
	Short: "Start a country capitals guessing game",
	Long: "In this game, you need to choose the correct country of a capital city. You can choose the number of options that will be displayed.",
	Example: "capitals 5",
	Run: func(cmd *cobra.Command, args []string) {
		size := 2

		if len(args) > 0 {
			val, err := strconv.Atoi(args[0])
			if err == nil && val > 2 {
				size = val
			}
		}


		c := country_info.NewCountryInfoClient()

		fmt.Println("Loading...\n")
		countries, err := c.GetCountries()
		if err != nil {
			log.Println("Ooops", err)
		}

		var points int
		var tries int
		for  {
			selectedCountries := RandomCountries(countries, size)
			correctCountry := RandomCountry(selectedCountries)

			capital, err := c.GetCapital(correctCountry.Code)
			if err != nil {
				log.Println("Ooops", err)
			}


			fmt.Println("\nOptions:")
			for _, c := range selectedCountries {
				fmt.Println(c.Code, "-", c.Name)
			}

			fmt.Println("\nType END to quit\n")

			code := strings.ToUpper(Ask("\nWhat is the country code of the capital " + capital + "? "))

			if code == "END" {
				break
			}

			if code == correctCountry.Code {
				points++
				fmt.Println("Your anwser is correct!\n\n")
			} else {
				fmt.Println("Oops... The correct country was " + correctCountry.Name)
			}

			fmt.Println("\n---------------------------------\n")
			Ask("Press ENTER to continue")
			tries++
		}

		fmt.Println("\n\nYou answered", points, "correctly of", tries, "questions.")
		if points == 0 {
			fmt.Println("It was BAAAAD")
		}
	},
}


func Ask(question string) string {
	fmt.Print(question)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	return input.Text()
}

func RandomCountries(countries []country_info.Country, size int) []country_info.Country {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(countries), func(i, j int) {
		countries[i], countries[j] = countries[j], countries[i]
	})

	return countries[0:size]
}

func RandomCountry(countries []country_info.Country) country_info.Country {
	return RandomCountries(countries, 1)[0]
}
