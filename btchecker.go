package main

import (
	"encoding/json"
	"net/http"
	"os"
	"log"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("USAGE: btchecker [PHONE NUMBER]")
	}

	phoneNumber := os.Args[1]

	// url := "http://www.productsandservices.bt.com/consumerProducts/v1/productAvailability.do?format=json&telephone=" + phoneNumber
	url := "http://dev.hentan.caius.name/bt.json?format=json&telephone=" + phoneNumber

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("Request failed; double check the phone number is correct?")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body);
	if err != nil {
		log.Fatal(err)
	}

	// println(string(body))

	var data map[string]interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	println(string(data["exchangeName"]))
}

/*
{"salesFilter1Date":"01-03-2013","exchangeName":"MARKET DRAYTON","salesFilter1Switch":"ON","telephone":"01630658343","serviceLineTypes":[{"supportsMulticast":"RED","CLEAN_TOP_DOWNSPEED":"40000","readyDate":"31-05-2014","exchangeState":"P","supportsVision":"RED","infinity":true,"CLEAN_BOTTOM_DOWNSPEED":"39000","CLEAN_BOTTOM_UPSPEED":"9000","capacityDate":"","CLEAN_TOP_UPSPEED":"10000","bbUpsellAvailable":"N","bb":"GREEN","serviceLineType":"BBSLT_UPTO_40M_10M_FTTCWBC"},{"supportsMulticast":"RED","readyDate":"","speed":"7M","capacityDate":"","exchangeState":"E","minRangeSpeed":"4.5M","bbUpsellAvailable":"N","maxRangeSpeed":"11.5M","bb":"GREEN","supportsVision":"GREEN","infinity":false,"serviceLineType":"BBSLT_UPTO_8M_WBC"},{"supportsMulticast":"RED","CLEAN_TOP_DOWNSPEED":"80000","readyDate":"31-05-2014","exchangeState":"P","supportsVision":"RED","infinity":true,"CLEAN_BOTTOM_DOWNSPEED":"79900","CLEAN_BOTTOM_UPSPEED":"20000","capacityDate":"","CLEAN_TOP_UPSPEED":"20000","bbUpsellAvailable":"N","bb":"GREEN","serviceLineType":"BBSLT_UPTO_80M_20M_FTTCWBC"},{"supportsMulticast":"RED","CLEAN_TOP_DOWNSPEED":"40000","readyDate":"31-05-2014","exchangeState":"P","supportsVision":"RED","infinity":true,"CLEAN_BOTTOM_DOWNSPEED":"39000","CLEAN_BOTTOM_UPSPEED":"1900","capacityDate":"","CLEAN_TOP_UPSPEED":"2000","bbUpsellAvailable":"N","bb":"GREEN","serviceLineType":"BBSLT_UPTO_40M_FTTCWBC"}],"exchangeCode":"WNMD"}
*/
