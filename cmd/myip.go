/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// myipCmd represents the myip command
var myipCmd = &cobra.Command{
	Use:   "myip",
	Short: "Get my Public ip Status & Data",
	Long: `Get More Information from my public Address like
	Country_Code & ip`,
	Run: func(cmd *cobra.Command, args []string) {
		getMyIpAddress()
	},
}

func init() {
	rootCmd.AddCommand(myipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// myipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// myipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Ip struct {
	Country_Code string `json:"Country_code"`
	Ip           string `json:"ip"`
}

func getMyIpAddress() {
	baseUrl := "https://ifconfig.io/all.json"

	res, err := http.Get(baseUrl)
	if err != nil {
		log.Printf("err in Create Request %v", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("err in get Response %v", err)
	}

	ip := Ip{}
	if err := json.Unmarshal(body, &ip); err != nil {
		log.Printf("error in Parse Json Response %v", err)
	}

	fmt.Printf("IP Addres: %v\n", ip.Ip)
	fmt.Printf("Country Code: %v\n", ip.Country_Code)

}
