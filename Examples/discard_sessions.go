package Examples

import (
	"fmt"
	api "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"os"
)

func DiscardSessions() {

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

	if x, _ := client.CheckFingerprint(); !x {
		fmt.Printf("Could not get the server's fingerprint - Check connectivity with the server.\n")
		os.Exit(1)
	}

	loginRes, err := client.ApiLogin(username, password, false, "", false, nil)
	if err != nil {
		fmt.Printf("Login error.\n")
		os.Exit(1)
	}

	if !loginRes.Success {
		fmt.Printf("Login failed:\n" + loginRes.ErrorMsg)
		os.Exit(1)
	}

	payload := map[string]interface{}{
		"name":       "fake_name2",
		"ip-address": "1.1.1.2",
	}
	_, err = client.ApiCall("add-host", payload, client.GetSessionID(), false, false)

	if err != nil {
		fmt.Printf("error" + err.Error() + "\n")
	}
	//payload = map[string]interface{} {}
	//client.ApiCall("publish", payload, client.GetSessionID(), false, false)

	showSessionsRes, err := client.ApiQuery("show-sessions", "full", "objects", false, map[string]interface{}{})

	if err != nil {
		fmt.Printf("Failed to retrieve the sessions\n")
		return
	}

	_, err2 := client.ApiQuery("show-hosts", "full", "objects", false, map[string]interface{}{})

	if err2 != nil {
		fmt.Printf("Failed to retrieve the sessions\n")
		return
	}

	//fmt.Println(show_sessions_res.GetData())
	var discardRes api.APIResponse
	for _, sessionObj := range showSessionsRes.GetData() {
		//fmt.Println(sessionObj)
		if sessionObj.(map[string]interface{})["application"].(string) != "WEB_API" {
			continue
		}
		discardRes, _ = client.ApiCall("discard", map[string]interface{}{"uid": sessionObj.(map[string]interface{})["uid"]}, "", false, false)

		if discardRes.Success {
			fmt.Println("Session " + sessionObj.(map[string]interface{})["uid"].(string) + " discarded successfully")
		} else {
			fmt.Println("Session " + sessionObj.(map[string]interface{})["uid"].(string) + " failed to discard")
		}
	}

}
