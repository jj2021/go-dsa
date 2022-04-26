/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"godsa/pkg/dsa"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// keygenCmd represents the keygen command
var keygenCmd = &cobra.Command{
	Use:   "keygen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating Keys... This may take a minute")
		pair := dsa.GenerateKeyPair()
		/*
			fmt.Printf("params: %+v\n", pair.Params)
			fmt.Printf("Private: %+v\n", pair.Private)
			fmt.Printf("Public: %+v\n", pair.Public)
		*/

		// write public key file
		pubfile := viper.New()
		pubfile.SetConfigType("yaml")

		pubfile.Set("p", pair.Params.P)
		pubfile.Set("q", pair.Params.Q)
		pubfile.Set("g", pair.Params.G)
		pubfile.Set("y", pair.Public.Int)

		pubfile.WriteConfigAs("./dsa_pub.yaml")

		// write private key file
		privfile := viper.New()
		privfile.SetConfigType("yaml")

		privfile.Set("p", pair.Params.P)
		privfile.Set("q", pair.Params.Q)
		privfile.Set("g", pair.Params.G)
		privfile.Set("x", pair.Private.Int)

		privfile.WriteConfigAs("./dsa.yaml")
	},
}

func init() {
	rootCmd.AddCommand(keygenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// keygenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keygenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
