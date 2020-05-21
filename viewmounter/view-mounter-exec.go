// Copyright 2020 Cohesity Inc.
//
// Author: Steve Klosky
//
// Script to mount currently visible Cohesity views to
// the local linux machine
//
// the intent is to use this program in a docker container
// to augment the rclone utility in the marketplace with
// by giving local file system access to Cohesity views
//

package main

import (

  "fmt"
  "os"
  "bytes"
  "encoding/json"
  "io/ioutil"
  "net/http"
//  "bufio"
//  "strconv"
//  "strings"
//  "time"
)

type authTokenRes struct {
    AccessToken    string   `json:"accessToken"`
    Privileges    []string   `json:"privileges"`
    TokenType    string   `json:"tokenType"`
}

func PrintUsageAndExit() {
  usage := `view-mounter-exec`
  fmt.Println(usage)
  os.Exit(0)
}

func GetAuthToken() string {

  authToken := ""
  endpointIp := ""
  endpointPort := ""
  username := ""
  password := ""
  domain := ""
  var err error
  var body []byte
  var rContainer authTokenRes

  _ = err
  _ = body
  _ = authToken

  // gather user input
  fmt.Println("Please enter ENDPOINT IP")
  fmt.Scanln(&endpointIp)
  fmt.Println("Please enter ENDPOINT PORT")
  fmt.Scanln(&endpointPort)
  fmt.Println("Please enter Username")
  fmt.Scanln(&username)
  fmt.Println("Please enter Password")
  fmt.Scanln(&password)

  // get auth token from cluster
  tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  }
  url := "https://" + endpointIp + ":" + endpointPort + "/irisservices/api/v1/public/accessTokens"
  valuePair := map[string]string{"domain": domain, "password": password, "username": username}
  postData, _ := json.Marshal(valuePair)
  postReader := bytes.NewBuffer(postData)
  client := &http.Client{Transport: tr}
  req, _ := http.NewRequest("POST", url, postReader)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-type", "application/json")
  res, _ := client.Do(req)
  body, _ = ioutil.ReadAll(res.Body)
  err = json.Unmarshal(body, &rContainer)
  authToken = rContainer.AccessToken
  fmt.Println(authToken)
  defer res.Body.Close()
  return authToken
}

func main() {

  var body []byte
  var authToken string

  _ = authToken
  _ = body
  if len(os.Args) > 1  {
    PrintUsageAndExit()
  }

  fmt.Println(`Setting up Cohesity Access`)

  // if we are in an athena setting, we look for these environment
  // variables
  //
  // HOST_IP  # The Host IP on which the container is running.
  // APPS_API_ENDPOINT_IP # Cohesity App Server IP.
  // APPS_API_ENDPOINT_PORT # Cohesity App Server Port.
  // APP_AUTHENTICATION_TOKEN # Authetication Token to make Cohesity App API calls. 

  endpointIp := os.Getenv("APPS_API_ENDPOINT_IP")
  endpointPort := os.Getenv("APPS_API_ENDPOINT_PORT")
  authToken = os.Getenv("APP_AUTHENTICATION_TOKEN")

  // Check to see if there are environment variables for login
  // If not, prompt for access info
  if (len(authToken) == 0) {
    authToken = GetAuthToken()
  }

  fmt.Println(`Getting List of Views`)

  url := "https://" + endpointIp + ":" + endpointPort + "/irisservices/api/v1/public/views"
  client := &http.Client{}
  req, _ := http.NewRequest("GET", url, nil)
  req.Header.Set("Authorization", "Bearer " + authToken)
  req.Header.Add("Accept", "application/json")
  res, _ := client.Do(req)
  body, _ = ioutil.ReadAll(res.Body)
  res.Body.Close()

  fmt.Println(`Mounting Views`)
}
