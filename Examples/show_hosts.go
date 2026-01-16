package Examples

import (
	"fmt"
	api "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"os"
)

func ShowHosts() {

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

	showHostsRes, err2 := client.ApiQuery("show-hosts", "full", "objects", false, map[string]interface{}{})

	if err2 != nil {
		fmt.Printf("Failed to retrieve the hosts\n")
		return
	}

	//fmt.Println(show_sessions_res.GetData())
	for _, sessionObj := range showHostsRes.GetData() {
		host := sessionObj.(map[string]interface{})
		fmt.Println("-------------\n--------------")
		fmt.Println(host["name"].(string))
		fmt.Println(host["ipv4-address"].(string))
	}

}
