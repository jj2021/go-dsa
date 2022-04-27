/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/hex"
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
		phex := privfile.GetString("p")
		ghex := privfile.GetString("g")
		qhex := privfile.GetString("q")
		xhex := privfile.GetString("x")

		p, err := hex.DecodeString(phex)
		if err != nil {
			fmt.Printf("Error decoding private key file")
			return
		}
		g, err := hex.DecodeString(ghex)
		if err != nil {
			fmt.Printf("Error decoding private key file")
			return
		}
		q, err := hex.DecodeString(qhex)
		if err != nil {
			fmt.Printf("Error decoding private key file")
			return
		}
		x, err := hex.DecodeString(xhex)
		if err != nil {
			fmt.Printf("Error decoding private key file")
			return
		}

		params := dsa.Parameters{P: new(big.Int).SetBytes(p), G: new(big.Int).SetBytes(g), Q: new(big.Int).SetBytes(q)}
		privkey := dsa.NewPrivateKey(new(big.Int).SetBytes(x))
		// fmt.Printf("%+v\n", params)
		// fmt.Printf("%+v\n", privkey)

		signature := dsa.Sign(content, privkey, params)

		// write signature file

		sigfile := viper.New()
		sigfile.SetConfigType("yaml")

		rhex := hex.EncodeToString(signature.R.Bytes())
		shex := hex.EncodeToString(signature.S.Bytes())

		sigfile.Set("r", rhex)
		sigfile.Set("s", shex)

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
