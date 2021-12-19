package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ScrapHandler(w http.ResponseWriter, req *http.Request) {
	// get request query
	username := req.URL.Query().Get("u")
	accountUsername := req.URL.Query().Get("au")
	accountPassword := req.URL.Query().Get("p")
	limit := req.URL.Query().Get("limit")
	fmt.Println()
	if username == "" || accountUsername == "" || accountPassword == "" || limit == "" {
		SendErrorResponse(w, ErrReqQueryInvalid)
		return
	}
	parsedLimit, err := strconv.Atoi(limit)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}
	if parsedLimit == 0 {
		limit = "10"
	}

	// run scraper command
	cmd := exec.Command(
		ScrapCommand,
		username,
		ScrapCommandArgAccountUsername,
		accountUsername,
		ScrapCommandArgAccountPassword,
		accountPassword,
		ScrapCommandArgDataDir,
		"data",
		ScrapCommandArgUsernameSubdir,
		ScrapCommandArgProfileMetadata,
		ScrapCommandArgMediaMetadata,
		ScrapCommandArgCookie,
		"cookie",
		ScrapCommandArgLimit,
		limit,
	)
	err = cmd.Run()
	if err != nil {
		LogError(w, err)
		return
	}

	dirPath := "./data/" + username
	dat, err := os.ReadDir(dirPath)
	if err != nil {
		LogError(w, err)
		return
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
		LogError(w, fmt.Errorf("scraper json file does not exist"))
		return
	}

	jsonFilePath := dirPath + "/" + jsonFileName
	fmt.Printf("scraper json file exist, filepath: %s\n", jsonFilePath)

	// open json file
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		LogError(w, err)
		return
	}
	fmt.Println("successfully opened json file")
	defer jsonFile.Close()

	// read json file
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		LogError(w, err)
		return
	}

	// unmarshal the json into struct
	var jsonData JsonData
	err = json.Unmarshal(byteValue, &jsonData)
	if err != nil {
		LogError(w, err)
		return
	}

	// TODO: process data and store it to db
	profileInfo := jsonData.GraphProfileInfo

	SendSuccessResponse(w, "scrap success", &ProfileResponse{
		UserID:      profileInfo.Info.ID,
		FullName:    profileInfo.Info.FullName,
		Username:    profileInfo.Username,
		Biography:   profileInfo.Info.Biography,
		PostsCount:  profileInfo.Info.PostsCount,
		CreatedTime: profileInfo.CreatedTime,
	})
}
