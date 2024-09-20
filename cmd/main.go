package main

import (
	"fmt"
	"myvpn/pkg"
	"myvpn/utils"
)

func main() {
	// response, err := pkg.GetAccessKeys()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// for _, data := range response.AccessKeys {
	// 	fmt.Printf("ID: %s, Name: %s, Method: %s, Port: %d, AccessUrl: %s\n",
	// 		data.ID, data.Name, data.Method, data.Port, data.AccessUrl)
	// }

	response, err := pkg.CreateAccessKey(utils.CreateAccessKeyPayload{
		Method:   "chacha20-ietf-poly1305",
		Name:     "lera",
		Port:     443,
		Limit: utils.Limit{
			Bytes: 1000,
		},
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, data := range response {
		fmt.Printf("ID: %s, Name: %s, Method: %s, Port: %d, AccessUrl: %s\n",
			data.ID, data.Name, data.Method, data.Port, data.AccessUrl)
	}
}
