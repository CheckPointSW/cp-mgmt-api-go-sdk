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
		fmt.Printf("Could not get the server's fingerprint - Check connectivity with the server.\n")
		os.Exit(1)
	}

	loginRes, err := client.ApiLogin(username, password, false, "", false, nil)
	if err != nil {
		fmt.Printf("Login error.\n")
		os.Exit(1)
	}

	if loginRes.Success == false {
		fmt.Printf("Login failed:\n" + loginRes.ErrorMsg)
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
		fmt.Printf("error" + err.Error() + "\n")
	}

	if addHostResponse.Success {
		fmt.Printf("The host: " + hostName + " has been added successfully\n")

		// publish the result
		payload = map[string]interface{}{}

		publish_res, err := client.ApiCall("publish", payload, client.GetSessionID(), true, true)
		if publish_res.Success {
			fmt.Printf("The changes were published successfully.\n")
		} else {
			fmt.Printf("Failed to publish the changes. \n" + err.Error())
		}
	} else {
		fmt.Printf("Failed to add the host: '" + hostName + "', Error:\n" + addHostResponse.ErrorMsg)
	}

}
