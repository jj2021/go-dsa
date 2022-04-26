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

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a signature with a dsa public key.",
	Long:  `Verify a given file and signature with the public key found in the dsa_pub.yaml file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("verify called")

		// read file content
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Printf("Error: No file provided for verification. Use the file option to specify the file to be verified\n")
			return
		}

		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}

		// read public key file
		pubfile := viper.New()
		pubfile.SetConfigType("yaml")
		pubfile.SetConfigFile("./dsa.yaml")

		if err := pubfile.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Printf("public key file not found\n")
				return
			} else {
				fmt.Printf("Could not read public key file\n")
				return
			}
		}

		// read signature file
		sigfile := viper.New()
		sigfile.SetConfigType("yaml")

		signaturePath, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Printf("Error: No signature provided for verification. Use the signature option to specify the signature to be verified\n")
			return
		}
		sigfile.SetConfigFile(signaturePath)

		if err := sigfile.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Printf("Signature file not found\n")
				return
			} else {
				fmt.Printf("Could not read signature file\n")
				return
			}
		}

		// verify
		p := pubfile.GetInt64("p")
		g := pubfile.GetInt64("g")
		q := pubfile.GetInt64("q")
		y := pubfile.GetInt64("y")
		params := dsa.Parameters{P: big.NewInt(p), G: big.NewInt(g), Q: big.NewInt(q)}
		pubkey := dsa.NewPublicKey(big.NewInt(y))

		r := sigfile.GetInt64("r")
		s := sigfile.GetInt64("s")
		signature := dsa.Signature{R: big.NewInt(r), S: big.NewInt(s)}

		valid, err := dsa.Verify(signature, content, pubkey, params)
		if err != nil || !valid {
			fmt.Printf("Signature is invalid. The signer may not be trust worthy, the message may have been tampered with, or the message may have been signed incorrectly.\n")
			return
		}

		fmt.Printf("Signature verified: valid\n")

	},
}

func init() {
	verifyCmd.Flags().StringP("file", "f", "", "The file to be verified.")
	verifyCmd.Flags().StringP("signature", "s", "./dsa_signature.yaml", "The signature file to be verified.")
	rootCmd.AddCommand(verifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
