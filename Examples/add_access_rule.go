package Examples

import (
	"fmt"
	api "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"os"
)

func AddAccessRule() {

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

	fmt.Printf("Enter the name of the access rule: ")
	var ruleName string
	fmt.Scanln(&ruleName)

	if isFingerPrintTrusted, err := client.CheckFingerprint(); isFingerPrintTrusted == false || err != nil {

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("Could not get the server's fingerprint - Check connectivity with the server.\n")
		}
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

	// add a rule to the top of the "Network" layer
	payload := map[string]interface{}{
		"name":     ruleName,
		"layer":    "Network",
		"position": "top",
	}
	addRuleResponse, err := client.ApiCall("add-access-rule", payload, client.GetSessionID(), false, true)

	if err != nil {
		fmt.Printf("error" + err.Error() + "\n")
	}

	if addRuleResponse.Success {
		fmt.Printf("The rule: " + ruleName + " has been added successfully\n")

		// publish the result
		payload = map[string]interface{}{}

		publishRes, err := client.ApiCall("publish", payload, client.GetSessionID(), false, true)
		if publishRes.Success {
			fmt.Printf("The changes were published successfully.\n")
		} else {
			fmt.Printf("Failed to publish the changes. \n" + err.Error())
		}
	} else {
		fmt.Printf("Failed to add the access-rule: '" + ruleName + "', Error:\n" + addRuleResponse.ErrorMsg)
	}

}
