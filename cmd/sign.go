/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"godsa/pkg/dsa"
	"io/ioutil"
	"math/big"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign a file with a dsa private key",
	Long:  `Signs a specified file with the dsa private key found in the dsa.yaml file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sign called")

		// read file content
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Printf("Error: No file provided for signing. Use the file option to specify the file to be signed\n")
			return
		}

		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}

		// read dsa private key file
		privfile := viper.New()
		privfile.SetConfigType("yaml")
		privfile.SetConfigFile("./dsa.yaml")

		if err := privfile.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Printf("Private key file not found\n")
				return
			} else {
				fmt.Printf("Could not read private key file\n")
				return
			}
		}

		// sign content
		p := privfile.GetInt64("p")
		g := privfile.GetInt64("g")
		q := privfile.GetInt64("q")
		x := privfile.GetInt64("x")
		params := dsa.Parameters{P: big.NewInt(p), G: big.NewInt(g), Q: big.NewInt(q)}
		privkey := dsa.NewPrivateKey(big.NewInt(x))
		signature := dsa.Sign(content, privkey, params)

		// write signature file
		sigfile := viper.New()
		sigfile.SetConfigType("yaml")

		sigfile.Set("r", signature.R)
		sigfile.Set("s", signature.S)

		sigfile.WriteConfigAs("./dsa_signature.yaml")
	},
}

func init() {
	signCmd.Flags().StringP("file", "f", "", "The file to be signed.")
	rootCmd.AddCommand(signCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
