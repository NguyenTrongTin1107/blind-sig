package cmd

import "github.com/spf13/cobra"

// Receive userId and ticket path
var spend = &cobra.Command{
	Use: "spend",
	Run: func(cmd *cobra.Command, args []string) {
		// key, err := rsa.GenerateKey(rand.Reader, KeySize)
		// if err != nil {
		// 	fmt.Println("Failed to generate key:", err)
		// 	return
		// }
		// err = saveKeyToFile(key, PrivKeyFile)
		// if err != nil {
		// 	fmt.Println("Failed to save private key:", err)
		// 	return
		// }
		//
		// err = savePublicKeyToFile(&key.PublicKey, PubKeyFile)
		// if err != nil {
		// 	fmt.Println("Failed to save public key:", err)
		// 	return
		// }
	},
}
