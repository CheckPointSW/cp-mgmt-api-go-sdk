package Examples

import (
	"fmt"
	api "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"os"
	"strconv"
	"time"
)

// This program demonstrates auto publish feature
// Publish will run automatically on every 'X' number of api calls, set 'autoPublishBatchSize' to -1 (or any negative number) to disable the feature
// In this program, 10 threads runs in parallel, each thread create 20 group objects, auto publish runs on every 100 created objects
func AutoPublish() {
	numOfThreads := 10                               // Total number of threads
	numOfObjectsToCreate := 20                       // Total number of objects to create by each thread
	autoPublishBatchSize := api.AutoPublishBatchSize // Call publish on every 100 api calls
	threadNamePrefix := "auto-publish-thread"
	var apiServer string
	var username string
	var password string

	fmt.Printf("Enter server IP address or hostname: ")
	fmt.Scanln(&apiServer)

	fmt.Printf("Enter username: ")
	fmt.Scanln(&username)

	fmt.Printf("Enter password: ")
	fmt.Scanln(&password)

	args := api.APIClientArgs(api.DefaultPort, "", "", apiServer, "", -1, "", false, false, "deb.txt", api.WebContext, api.TimeOut, api.SleepTime, "", "", autoPublishBatchSize)

	client := api.APIClient(args)

	if ok, _ := client.CheckFingerprint(); !ok {
		fmt.Println("Could not get the server's fingerprint - Check connectivity with the server")
		os.Exit(1)
	}

	loginRes, err := client.ApiLogin(username, password, false, "", false, nil)
	if err != nil {
		fmt.Println("Login error")
		os.Exit(1)
	}

	if !loginRes.Success {
		fmt.Println("Login failed: " + loginRes.ErrorMsg)
		os.Exit(1)
	}

	fmt.Println("Start auto publish program. Number of threads " + strconv.Itoa(numOfThreads))
	for i := 0; i < numOfThreads; i++ {
		go run(threadNamePrefix+strconv.Itoa(i), numOfObjectsToCreate, client)
	}

	time.Sleep(time.Minute * 3)

	_, _ = client.ApiCallSimple("publish", map[string]interface{}{}) // publish leftovers if exists

	fmt.Println("Finished to create all objects")

	deleteObjects(client, threadNamePrefix)

	_, _ = client.ApiCallSimple("logout", map[string]interface{}{})

	fmt.Println("Auto publish example finished successfully")
}

func run(threadName string, numOfObjectsToCreate int, apiClient interface{}) {
	fmt.Println("Start thread -> " + threadName)
	client := apiClient.(*api.ApiClient)
	for i := 0; i < numOfObjectsToCreate; i++ {
		groupName := threadName + "-group" + strconv.Itoa(i)
		res, err := client.ApiCallSimple("add-group", map[string]interface{}{"name": groupName})
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		if !res.Success {
			fmt.Println("Failed to add group: " + res.ErrorMsg)
		}
	}
	fmt.Println(threadName + " finished")
}

func deleteObjects(client *api.ApiClient, objPrefix string) {
	res, _ := client.ApiQuery("show-groups", "standard", "", true, map[string]interface{}{"filter": objPrefix})
	if res.Success {
		client.ResetTotalCallsCounter()
		objects := res.GetData()["objects"].([]map[string]interface{})
		fmt.Println("Delete " + strconv.Itoa(len(objects)) + " objects...")
		if len(objects) > 0 {
			for i := 0; i < len(objects); i++ {
				client.ApiCallSimple("delete-group", map[string]interface{}{"uid": objects[i]["uid"]})
			}
			client.ApiCallSimple("publish", map[string]interface{}{}) // publish leftovers if exists
			fmt.Println("Groups deleted")
		} else {
			fmt.Println("Not found groups to delete")
		}
	}
}
