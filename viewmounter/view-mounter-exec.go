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
  Username := "admin"
  Password := "admin"
  ClusterVip := "172.16.3.101"
  var Domain string // Set for AD user only.
  client := CohesityManagementSdk.NewCohesitySdkClient(ClusterVip, Username, Password, Domain)
  //CohesityManagementSdk.NewCohesityClientWithToken(hostIp, &managementAccessToken)
  
  fmt.Println(`Getting List of Views`)
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
    fmt.Println(myViewName)
  }

  fmt.Println(`Mounting Views`)

}
