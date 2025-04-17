package cetus

import (
	"encoding/json"
	"fmt"

	"github.com/ipoluianov/suigo/client"
)

func GetPosition(objectId string) (*CetusPoolPosition, error) {
	var c CetusPoolPosition

	cl := client.NewClient(client.MAINNET_URL)

	var options client.GetObjectShowOptions
	options.ShowDisplay = true
	options.ShowOwner = true
	options.ShowContent = true

	res, err := cl.GetObject(objectId, options)
	if err != nil {
		return nil, err
	}

	bs, _ := json.MarshalIndent(res.Data.Content, "", "  ")
	//fmt.Println("res:", string(bs))
	json.Unmarshal(bs, &c)

	return &c, nil
}

func GetPool(objectId string) (*CetusPool, error) {
	var c CetusPool

	cl := client.NewClient(client.MAINNET_URL)

	var options client.GetObjectShowOptions
	options.ShowDisplay = true
	options.ShowOwner = true
	options.ShowContent = true

	res, err := cl.GetObject(objectId, options)
	if err != nil {
		return nil, err
	}

	bs, _ := json.MarshalIndent(res.Data.Content, "", "  ")
	fmt.Println("res:", string(bs))
	json.Unmarshal(bs, &c)

	return &c, nil
}
