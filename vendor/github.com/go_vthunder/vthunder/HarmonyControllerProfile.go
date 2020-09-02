package go_vthunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"util"
)

type Profile struct {
	Host ProfileInstance `json:"profile,omitempty"`
}
type ThunderMgmtIP struct {
	IPAddress string `json:"ip-address,omitempty"`
}
type ProfileInstance struct {
	Host             string        `json:"host,omitempty"`
	Port             int           `json:"port,omitempty"`
	UserName         string        `json:"user-name,omitempty"`
	SecretValue      string        `json:"secret-value,omitempty"`
	Provider2        string        `json:"provider,omitempty"`
	IPAddress        ThunderMgmtIP `json:"thunder-mgmt-ip,omitempty"`
	Action           string        `json:"action,omitempty"`
	UseMgmtPort      int           `json:"use-mgmt-port,omitempty"`
	Region           string        `json:"region,omitempty"`
	AvailabilityZone string        `json:"availability-zone,omitempty"`
}

func GetProfile(id string, host string) (*ProfileInstance, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/harmony-controller/profile", nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ProfileInstance
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("Profile instance from Read")
			logger.Println(string(data))
			logger.Println("[INFO] GET REQ RES..........................", m)
			return &m, nil
		}
	}
}

func PostProfile(id string, vc Profile, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(vc)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/harmony-controller/profile", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var v ProfileInstance
		erro := json.Unmarshal(data, &v)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
		}
	}

}
