package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type Location struct {
	Name string `json:"name"`
}

type Weather struct {
	Temperature string `json:"temperature"`
}

type Response struct {
	Location Location `json:"location"`
	Current  Weather  `json:"current"`
}

func getTemperature(str string) string {
	httpClient := http.Client{}
	req, err := http.NewRequest("GET", "http://api.weatherstack.com/current", nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("acces_key", "0eef548ef1ee12f7d680b1c34c2476b0")
	q.Add("query", str)
	req.URL.RawQuery = q.Encode()

	res, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}

func main() {

	var cmdTemperature = &cobra.Command{
		Use:   "temperatura [cadena a imprimir]",
		Short: "Muestra la temperatura de un lugar.",
		Long:  `Muestra la temperatura de un lugar especificando su nombre como parámetro del comando.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Mostrar Temperatura")
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdTemperature)
	rootCmd.Execute()
}
