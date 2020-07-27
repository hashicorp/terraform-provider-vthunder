package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type NatIcmpv6 struct {
	UUID NatIcmpv6Instance `json:"icmpv6,omitempty"`
}
type NatIcmpv6Instance struct {
	RespondToPing int    `json:"respond-to-ping,omitempty"`
	UUID          string `json:"uuid,omitempty"`
}

func PostIpv6NatIcmpv6(id string, inst NatIcmpv6, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpv6NatIcmpv6")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ipv6/nat/icmpv6", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m NatIcmpv6
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetIpv6NatIcmpv6(id string, host string) (*NatIcmpv6, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpv6NatIcmpv6")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ipv6/nat/icmpv6/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m NatIcmpv6
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
