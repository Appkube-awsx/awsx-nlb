package vault

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)


type Response struct {
	Region              string `json:"region,omitempty"`
	AccessKey           string `json:"accessKey,omitempty"`
	SecretKey           string `json:"secretKey,omitempty"`
	CrossAccountRoleArn string `json:"crossAccountRoleArn,omitempty"`
	ExternalId          string `json:"externalId,omitempty"`
}


func GetAccountDetails(accountNo string) (*Response, error) { // vaultUrl string
	log.Println("Calling account details API")
	client := &http.Client{}
	req, err := http.NewRequest("GET", accountNo, nil) //vaultUrl+"?accountNo="+
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	var responseObject Response
	json.Unmarshal(bodyBytes, &responseObject)
	//fmt.Printf("API Response as struct %+v\n", responseObject)
	return &responseObject, nil
	//fmt.Printf("API Response as struct %+v\n", bodyBytes)
}
