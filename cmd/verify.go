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

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a signature with a dsa public key.",
	Long:  `Verify a given file and signature with the public key found in the dsa_pub.yaml file.`,
	Run: func(cmd *cobra.Command, args []string) {

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
		pubfile.SetConfigFile("./dsa_pub.yaml")

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

		signaturePath, err := cmd.Flags().GetString("signature")
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

		// decode public key values
		phex := pubfile.GetString("p")
		ghex := pubfile.GetString("g")
		qhex := pubfile.GetString("q")
		yhex := pubfile.GetString("y")

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
		y, err := hex.DecodeString(yhex)
		if err != nil {
			fmt.Printf("Error decoding private key file")
			return
		}

		params := dsa.Parameters{P: new(big.Int).SetBytes(p), G: new(big.Int).SetBytes(g), Q: new(big.Int).SetBytes(q)}
		pubkey := dsa.NewPublicKey(new(big.Int).SetBytes(y))

		// decode signature values
		rhex := sigfile.GetString("r")
		shex := sigfile.GetString("s")

		r, err := hex.DecodeString(rhex)
		if err != nil {
			fmt.Printf("Error decoding private key file")
			return
		}
		s, err := hex.DecodeString(shex)
		if err != nil {
			fmt.Printf("Error decoding private key file")
			return
		}

		signature := dsa.Signature{R: new(big.Int).SetBytes(r), S: new(big.Int).SetBytes(s)}

		// verify
		valid, err := dsa.Verify(signature, content, pubkey, params)
		if err != nil || !valid {
			fmt.Printf("Signature is invalid. The signer may not be trust worthy, the message may have been tampered with, or the message may have been signed incorrectly.\n%v\n", err.Error())
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
