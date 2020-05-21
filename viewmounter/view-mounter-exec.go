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
//  "bufio"
//  "io/ioutil"
//  "net/http"
  "os"
//  "strconv"
//  "strings"
//  "time"
  "github.com/cohesity/app-sdk-go/appsdk"

)

func PrintUsageAndExit() {
  usage := `view-mounter-exec`
  fmt.Println(usage)
  os.Exit(0)
}

func main() {

  var username = ""
  var password = ""
  var appClient CohesityAppSdk.COHESITYAPPSDK

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

//  hostIp := os.Getenv("HOST_IP")
  endpointIp := os.Getenv("APPS_API_ENDPOINT_IP")
  endpointPort := os.Getenv("APPS_API_ENDPOINT_PORT")
  authToken := os.Getenv("APP_AUTHENTICATION_TOKEN")


// Check to see if there are environment variables for login
// If not, prompt for that info
  if ((len(endpointIp) + len(endpointPort) + len(authToken)) == 0) {
 //   fmt.Println("Please enter HOST IP")
 //   fmt.Scanln(&hostIp)      
    fmt.Println("Please enter ENDPOINT IP")
    fmt.Scanln(&endpointIp)
    fmt.Println("Please enter ENDPOINT PORT")
    fmt.Scanln(&endpointPort)
    fmt.Println("Please enter Username")
    fmt.Scanln(&username)
    fmt.Println("Please enter Password")
    fmt.Scanln(&password)
    appClient = CohesityAppSdk.NewAppSdkClient(authToken, endpointIp, endpointPort)
  } else {
    appClient = CohesityAppSdk.NewAppSdkClient(authToken, endpointIp, endpointPort)
  }

  fmt.Println(endpointIp, endpointPort, authToken, username, password)
  fmt.Println(`Getting List of Views`)
  fmt.Println(`Mounting Views`)
}

