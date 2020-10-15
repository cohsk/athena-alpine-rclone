# About Rclone Browser

With the widespread use of data in the cloud, teams need powerful tools to manage the data.  Object storage is very popular for cloud data.  Rclone (rclone.org) is a popular tool to help manage object based data.  Using Rsync (rsync.samba.org) principles as a starting point, Rclone was developed to sync data between object stores.  As of this writing (October 2020), Rclone supports approximate 30 different object storage systems.  Rclone is an active open source project and new capabilities are regularly added.  Rclone capabilities have evolved well beyond the original Rsync principles.

In 2019, a web based graphical user interface was added to Rclone.  The notion here is that a guided GUI will help users operate Rclone with a lower barrier to learning.  While the GUI is super useful, users may need to refer to the Rclone documentation from time to time for detailed information.

This port of Rclone and the Rclone web based GUI brings the both the CLI and the web GUI to the Cohesity Marketplace (aka Cohesity Athena Framework).
The intent is that administrators can use the GUI to explore configurations and object stores.  Additionally, once Rclone configs are set, administrators
can easily schedule rclone jobs for regular operations.  Note that the CLI does give access to many more options than the GUI.  For example, if an admin
is looking to work with invalid certificates, this is supported in the CLI, but not in the GUI.

There are several use cases for this utility in a Cohesity setting
* Data Migration
* Data Backup
* Archive Migration

# Rclone Browser Basic Operation
* Find the Rclone Browser app in the Cohesity Marketplace (https://marketplace.cohesity.com).  
* Use standard methods to load and run Rclone Browser onto a Cohesity Cluster.
* Access the app home page using the launch link from "My Apps / Rclone Browser"
* From here, use standard Rclone GUI workflows to manage cloud drive data (https://rclone.org/gui/)

As an example, here is a video of setting up a local store and a google cloud store -- https://youtu.be/H6noIywIVd0

Here are some useful references
* FAQ -- https://github.com/cohsk/athena-alpine-rclone/blob/master/FAQ.md
* Lessons Learned -- https://github.com/cohsk/athena-alpine-rclone/blob/master/Lessons-Learned.MD
* Developers -- https://github.com/cohsk/athena-alpine-rclone/blob/master/Dev%20Notes.md
