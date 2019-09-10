package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/rdegges/go-ipify"
)

//ConfigLocation is the 'config.json' location
var ConfigLocation = "config.json"

//Config is the setting which fetch from 'config.json'
var Config = config{Duration: 30 * time.Second}

//IP record address
var IP string
var getIP func() (string, error) = ipify.GetIp

type config struct {
	APIKey   string
	Domain   string
	Duration time.Duration
}
type configReq struct {
	APIKey   string `json:"api_key"`
	Domain   string `json:"domain"`
	Duration string `json:"duration"`
}

func updateConfig() error {
	b, err := ioutil.ReadFile(ConfigLocation)
	if err != nil {
		return err
	}
	c := configReq{}
	err = json.Unmarshal(b, &c)
	if err != nil {
		return err
	}
	if c.APIKey != "" {
		Config.APIKey = c.APIKey
	}
	if c.Domain != "" {
		Config.Domain = c.Domain
	}
	if c.Duration != "" {
		t, err := time.ParseDuration(c.Duration)
		if err != nil {
			return err
		}
		if t > 30*time.Second {
			Config.Duration = t
		}
	}
	return nil
}
func updateDNS() error {
	b := bytes.NewBuffer([]byte(`{"rrset_ttl": 10800,"rrset_values":["` + IP + `"]}`))
	req2, err := http.NewRequest(http.MethodPut, "https://dns.api.gandi.net/api/v5/domains/"+Config.Domain+"/records/%40/A", b)
	req2.Header.Set("X-Api-Key", Config.APIKey)
	req2.Header.Add("Content-Type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req2)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
func main() {
	for {
		err := updateConfig()
		if err != nil {
			log.Println(err)
		}
		ipAddress, err := getIP()
		if err != nil {
			log.Println("Couldn't get my IP address:", err)
		} else {
			log.Println("My IP address is:", ipAddress)
			if IP != ipAddress {
				IP = ipAddress
			}
			updateDNS()
		}
		time.Sleep(Config.Duration)
	}
}
