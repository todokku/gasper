package enrai

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/sdslabs/gasper/configs"
	"github.com/sdslabs/gasper/lib/redis"
	"github.com/sdslabs/gasper/lib/utils"
	"github.com/sdslabs/gasper/types"
)

func handleError(err error) {
	utils.Log("Failed to update Record Storage", utils.ErrorTAG)
	utils.LogError(err)
}

// filterValidInstances filters the instances and returns
// valid instances i.e which is in the form of IP:Port
func filterValidInstances(reverseProxyInstances []string) []string {
	filteredInstances := make([]string, 0)
	for _, instance := range reverseProxyInstances {
		if strings.Contains(instance, ":") {
			filteredInstances = append(filteredInstances, instance)
		} else {
			utils.LogError(fmt.Errorf("Instance %s is of invalid format", instance))
		}
	}
	return filteredInstances
}

// Updates the reverse proxy record storage periodically
func updateStorage() {
	apps, err := redis.FetchAllApps()
	if err != nil {
		handleError(err)
		return
	}

	updateBody := make(map[string]string)
	appInfoStruct := &types.InstanceBindings{}

	// Create entries for applications
	for name, data := range apps {
		resultByte := []byte(data)
		if err = json.Unmarshal(resultByte, appInfoStruct); err != nil {
			handleError(err)
			continue
		}
		updateBody[name] = appInfoStruct.Server
	}

	// Create enrties for Kaze in the load balancer
	kazeInstances, err := redis.FetchServiceInstances(types.Kaze)
	if err != nil {
		utils.Log("Failed to fetch kaze instances", utils.ErrorTAG)
	} else {
		kazeBalancer.Update(filterValidInstances(kazeInstances))
	}
	storage.Replace(updateBody)
}

// ScheduleUpdate runs updateStorage on given intervals of time
func ScheduleUpdate() {
	interval := configs.ServiceConfig.Enrai.RecordUpdateInterval * time.Second
	scheduler := utils.NewScheduler(interval, updateStorage)
	scheduler.RunAsync()
}
