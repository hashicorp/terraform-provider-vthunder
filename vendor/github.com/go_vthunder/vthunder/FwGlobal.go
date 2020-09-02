package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type FwGlobal struct {
	AlgProcessing FwGlobalInstance `json:"global-instance,omitempty"`
}

type FwGlobalInstance struct {
	AlgProcessing              string                   `json:"alg-processing,omitempty"`
	DisableApplicationProtocol []FwGlobalDisableAppList `json:"disable-app-list,omitempty"`
	DisableIPFwSessions        int                      `json:"disable-ip-fw-sessions,omitempty"`
	ExtendedMatching           string                   `json:"extended-matching,omitempty"`
	ListenOnPortTimeout        int                      `json:"listen-on-port-timeout,omitempty"`
	NatipDdosProtection        string                   `json:"natip-ddos-protection,omitempty"`
	PermitDefaultAction        string                   `json:"permit-default-action,omitempty"`
	RespondToUserMac           int                      `json:"respond-to-user-mac,omitempty"`
	Counters1                  []FwGlobalSamplingEnable `json:"sampling-enable,omitempty"`
	UUID                       string                   `json:"uuid,omitempty"`
}

type FwGlobalDisableAppList struct {
	DisableApplicationCategory string `json:"disable-application-category,omitempty"`
	DisableApplicationProtocol string `json:"disable-application-protocol,omitempty"`
}

type FwGlobalSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostFwGlobal(id string, inst FwGlobal, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwGlobal")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes 0 - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/global", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwGlobal
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetFwGlobal(id string, host string) (*FwGlobal, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwGlobal")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/global/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwGlobal
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			return &m, nil
		}
	}

}
