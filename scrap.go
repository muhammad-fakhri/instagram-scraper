package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func ScrapHandler(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	dirPath := "./data/" + username
	dat, err := os.ReadDir(dirPath)
	if err != nil {
		LogAndExit(err.Error())
	}

	// check json file existence
	isFileExist := false
	jsonFileName := ""
	for _, v := range dat {
		if strings.Contains(v.Name(), username+".json") {
			jsonFileName = v.Name()
			isFileExist = true
			break
		}
	}

	if !isFileExist {
		LogAndExit("scraper json file does not exist")
	}

	jsonFilePath := dirPath + "/" + jsonFileName
	fmt.Printf("scraper json file exist, filepath: %s\n", jsonFilePath)

	// open json file
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		LogAndExit(err.Error())
	}
	fmt.Println("successfully opened json file")
	defer jsonFile.Close()

	// read json file
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		LogAndExit(err.Error())
	}

	// unmarshal the json into struct
	var jsonData JsonData
	err = json.Unmarshal(byteValue, &jsonData)
	if err != nil {
		LogAndExit(err.Error())
	}

	// TODO: process data and store it to db
	profileInfo := jsonData.GraphProfileInfo

	SendJsonResponse(w, "scrap success", &ProfileResponse{
		UserID:      profileInfo.Info.ID,
		FullName:    profileInfo.Info.FullName,
		Username:    profileInfo.Username,
		Biography:   profileInfo.Info.Biography,
		PostsCount:  profileInfo.Info.PostsCount,
		CreatedTime: profileInfo.CreatedTime,
	})
}
