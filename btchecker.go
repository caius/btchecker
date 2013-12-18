package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ServiceResponse struct {
	ServiceLineType        string
	Infinity               bool
	SupportsVision         string
	ExchangeState          string
	ReadyDate              string
	CLEAN_BOTTOM_DOWNSPEED string
	CLEAN_BOTTOM_UPSPEED   string
	CLEAN_TOP_DOWNSPEED    string
	CLEAN_TOP_UPSPEED      string
}

func (service *ServiceResponse) ExchangeEnabled() bool {
	return service.ExchangeState == "E"
}

type CheckerResponse struct {
	Telephone        string
	ExchangeName     string
	ExchangeCode     string
	ServiceLineTypes []ServiceResponse
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("USAGE: btchecker [PHONE NUMBER]")
	}

	phoneNumber := os.Args[1]

	url := "http://www.productsandservices.bt.com/consumerProducts/v1/productAvailability.do?format=json&telephone=" + phoneNumber

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("Request failed; double check the phone number is correct?")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("DEBUG") != "" {
		println(string(body))
	}

	var data CheckerResponse
	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s served by %s (%s)\n", data.Telephone, data.ExchangeName, data.ExchangeCode)
	for _, service := range data.ServiceLineTypes {
		if service.Infinity == true {
			if service.ExchangeState == "P" {
				fmt.Printf("Infinity service pending; ready date is %s\n", service.ReadyDate)
			} else if service.ExchangeState == "E" {
				fmt.Println("Infinity service already enabled!")
			} else {
				fmt.Println("Infinity service status unknown :-(")
			}
		} else {
			fmt.Println("Infinity service status unknown :-(")
		}
		break
	}
}
