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
  "github.com/cohesity/app-sdk-go/appsdk"
  appModels "github.com/cohesity/app-sdk-go/models"
  CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
  managementModels "github.com/cohesity/management-sdk-go/models"
  )

// View Identification information.
type ViewInfo struct {
  ViewName string `json:"name,omitempty"`
  ViewId   int    `json:"viewId,omitempty"`
}

// Represents the information of views.
type ViewsInformation struct {
  ViewsInfo []*ViewInfo `json:"views,omitempty"`
}

func PrintUsageAndExit() {
  usage := `view-mounter-exec`
  fmt.Println(usage)
  os.Exit(0)
}

func main() {

  var viewsResult *managementModels.GetViewsResult
  var viewNames, viewBoxNames, tenantIds []string
  var matchPartialNames, includeInactive, allUnderHierarchy,
    sortByLogicalUsage, matchAliasNames *bool
  var maxCount, maxViewId *int64
  var viewBoxIds, jobIds []int64

  fmt.Println(`Setting up Cohesity Access`)
  Token := os.Getenv("APP_AUTHENTICATION_TOKEN")
  fmt.Println(Token)
  ClusterVip := os.Getenv("APPS_API_ENDPOINT_IP")
  fmt.Println(ClusterVip)
  ClusterPort := os.Getenv("APPS_API_ENDPOINT_PORT")
  fmt.Println(ClusterPort)
  HostIp := os.Getenv("HOST_IP")
  fmt.Println(HostIp)

  // var Domain string // Set for AD user only.
  fmt.Println(`setting up app client`)
  appClient := CohesityAppSdk.NewAppSdkClient(Token, ClusterVip, ClusterPort)

  fmt.Println(`setting up managementAccessToken`)
  var managementAccessToken managementModels.AccessToken
  fmt.Println(`creating management token`)
  managementTokenResponse, err := appClient.TokenManagement().CreateManagementAccessToken()
  fmt.Println(err)
  fmt.Println(`parsing management token response`)
  managementAccessToken = managementModels.AccessToken{
    AccessToken: managementTokenResponse.AccessToken,
    TokenType:   managementTokenResponse.TokenType,
  }

  // client := CohesityManagementSdk.NewCohesitySdkClient(ClusterVip, Username, Password, Domain)
  fmt.Println(`creating management client`)
  client := CohesityManagementSdk.NewCohesityClientWithToken(HostIp, &managementAccessToken)

  fmt.Println(`Getting and Mounting Views`)
  viewsResult, _ = client.Views().GetViews(viewNames, viewBoxNames,
    matchPartialNames, maxCount, maxViewId, includeInactive, tenantIds,
    allUnderHierarchy, viewBoxIds, jobIds, sortByLogicalUsage, matchAliasNames)
  clusterViews := viewsResult.Views
  clusterViewIDMap := make(map[int]string)

  // ClusterViewsInfo gives information about all the views in a cluster.
  // var clusterViewsInfo ViewsInformation
  var viewsInfo []*ViewInfo
  var myViewName string
  // Iterating over cluster views and storing viewname and id in a map.
  for _, view := range clusterViews {
    clusterViewIDMap[int(*view.ViewId)] = *view.Name
    viewInfo := ViewInfo{
      ViewName: *view.Name,
      ViewId:   int(*view.ViewId),
    }
    viewsInfo = append(viewsInfo, &viewInfo)
    myViewName = *view.Name

    // Name of the directory that is to be created and to be mounted the view on.
    dirName := myViewName + "_dir"
    options := "rw"

    // Options to be specified for the mount api.
    mountOptions := appModels.MountOptions{
    ViewName:      &myViewName,
    DirName:       dirName,
    MountProtocol: appModels.MountProtocol_KNFS,
    MountOptions:  &options,
  }

  createMountParams := appModels.CreateMountParams {
    MountOptions: &mountOptions,
  }

  // Api to mount the view.
  appClient.Mount().CreateMount(&createMountParams)

  }
}
