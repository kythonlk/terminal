package main

import (
	"fmt"
	"syscall/js"

	"github.com/supabase/postgrest-go"
)

var Client *postgrest.Client

func main() {

	SCHEMA := "public"

	header := map[string]string{
		"apikey":        API_KEY,
		"Authorization": AUTHORIZATION,
	}

	Client = postgrest.NewClient(REST_URL, SCHEMA, header)

	// Register the function that inserts data
	js.Global().Set("insertData", js.FuncOf(insertData))
	c := make(chan struct{}, 0)
	<-c
}

func insertData(this js.Value, p []js.Value) interface{} {
	tableName := p[0].String()
	data := p[1] // This is a JavaScript object
	primaryKey := p[2].String()

	dataMap := make(map[string]interface{})

	// Get keys from the JavaScript object
	keys := js.Global().Get("Object").Call("keys", data)
	for i := 0; i < keys.Length(); i++ {
		key := keys.Index(i).String()
		value := data.Get(key).String()
		dataMap[key] = value
	}

	response, count, err := InsertData(tableName, dataMap, primaryKey)
	if err != nil {
		return js.ValueOf(fmt.Sprintf("Error: %v", err))
	}

	return js.ValueOf(fmt.Sprintf("Response: %s, Count: %d", response, count))
}

func InsertData(tableName string, data map[string]interface{}, primaryKey string) (res string, count int, err error) {
	fmt.Println(data, primaryKey)
	response, responseCount, responseErr := Client.From(tableName).Select("*", "", false).Execute()
	if responseErr != nil {
		return "", 0, responseErr
	}

	return string(response), int(responseCount), nil
}
