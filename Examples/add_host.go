package Examples

import (
	"fmt"
	api "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"os"
)

func AddHost() {
	var apiServer string
	var username string
	var password string

	fmt.Printf("Enter server IP address or hostname: ")
	fmt.Scanln(&apiServer)

	fmt.Printf("Enter username: ")
	fmt.Scanln(&username)

	fmt.Printf("Enter password: ")
	fmt.Scanln(&password)

	args := api.APIClientArgs(api.DefaultPort, "", "", apiServer, "", -1, "", false, false, "deb.txt", api.WebContext, api.TimeOut, api.SleepTime, "", "", -1)

	client := api.APIClient(args)

	if x, _ := client.CheckFingerprint(); x == false {
		print("Could not get the server's fingerprint - Check connectivity with the server.\n")
		os.Exit(1)
	}

	loginRes, err := client.Login(username, password,false, "", false, "")
	if err != nil {
		print("Login error.\n")
		os.Exit(1)
	}

	if loginRes.Success == false {
		print("Login failed:\n" + loginRes.ErrorMsg)
		os.Exit(1)
	}

	fmt.Printf("Enter the name of the host: ")
	var hostName string
	fmt.Scanln(&hostName)

	fmt.Printf("Enter the ip of the host: ")
	var hostIp string
	fmt.Scanln(&hostIp)

	// add host
	payload := map[string]interface{}{
		"name":       hostName,
		"ip-address": hostIp,
	}
	addHostResponse, err := client.ApiCall("add-host", payload, client.GetSessionID(), false, true)

	if err != nil {
		print("error" + err.Error() + "\n")
	}

	if addHostResponse.Success {
		print("The host: " + hostName + " has been added successfully\n")

		// publish the result
		payload = map[string]interface{}{}

		publish_res, err := client.ApiCall("publish", payload, client.GetSessionID(), true, true)
		if publish_res.Success {
			print("The changes were published successfully.\n")
		} else {
			print("Failed to publish the changes. \n" + err.Error())
		}
	} else {
		print("Failed to add the host: '" + hostName + "', Error:\n" + addHostResponse.ErrorMsg)
	}

}
