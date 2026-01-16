package Examples

import (
	"fmt"
	api "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"os"
)

func DupIp() {

	fmt.Printf("Enter server IP address or hostname: ")
	var apiServer string
	fmt.Scanln(&apiServer)

	fmt.Printf("Enter username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Printf("Enter password: ")
	var pass string
	fmt.Scanln(&pass)

	args := api.APIClientArgs(api.DefaultPort, "", "", apiServer, "", -1, "", false, false, "deb.txt", api.WebContext, api.TimeOut, api.SleepTime, "", "", -1)

	client := api.APIClient(args)

	if x, _ := client.CheckFingerprint(); x == false {
		fmt.Printf("Could not get the server's fingerprint - Check connectivity with the server.\n")
		os.Exit(1)
	}

	loginRes, err := client.ApiLogin(username, pass, false, "", false, nil)
	if err != nil {
		fmt.Printf("Login error.\n")
		os.Exit(1)
	}

	if loginRes.Success == false {
		fmt.Printf("Login failed:\n" + loginRes.ErrorMsg)
		os.Exit(1)
	}

	showHostsRes, err := client.ApiQuery("show-hosts", "full", "objects", false, map[string]interface{}{})

	if err != nil {
		fmt.Printf("Failed to retrieve the hosts\n")
		return
	}

	// objDictionary - for a given IP address, get an array of hosts (name, unique-ID) that use this IP address.
	var objDictionary = map[string][]map[string]string{}

	//dupIpSlice - a collection of the duplicate IP addresses in all the host objects.
	var dupIpSlice []string

	for _, host := range showHostsRes.GetData() {
		ipaddr := host.(map[string]interface{})["ipv4-address"].(string)
		if ipaddr == "" {
			fmt.Printf(host.(map[string]interface{})["name"].(string) + " has no IPv4 address. Skipping...")
			continue
		}
		hostData := map[string]string{"name": host.(map[string]interface{})["name"].(string), "uid": host.(map[string]interface{})["uid"].(string)}

		//ok will be set to true if ipaddr is actually in objDictionary
		if _, ok := objDictionary[ipaddr]; ok {
			ipExists := false
			for _, ip := range dupIpSlice {
				if ip == ipaddr {
					ipExists = true
					break
				}
			}
			if !ipExists {
				dupIpSlice = append(dupIpSlice, ipaddr)
			}
			objDictionary[ipaddr] = append(objDictionary[ipaddr], hostData)
		} else {

			objDictionary[ipaddr] = []map[string]string{}
			objDictionary[ipaddr] = append(objDictionary[ipaddr], hostData)
		}

	}

	//print list of duplicate IP addresses to the console
	fmt.Println("List of Duplicate IP addresses: ")
	fmt.Println("-------------------------------")

	if len(dupIpSlice) == 0 {
		fmt.Println("No hosts with duplicate IP addresses")
	}

	//for every duplicate ip - print hosts with that ip:
	for _, dupIp := range dupIpSlice {

		fmt.Println("\nIP Address: " + dupIp + "")
		fmt.Println("-------------------------------")

		for _, hostData := range objDictionary[dupIp] {

			fmt.Println("host name: " + hostData["name"] + " host uid: " + hostData["uid"])
			//fmt.Println(hostData[1])

		}

	}

}
