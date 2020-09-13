# athena-alpine-rclone (aka "Rclone Brower")

This repo is for the Cohesity Rclone Browser Marketplace app

# What Rclone Browser does
This app creates a web based user interface to Rclone.  Rclone is a well know utility for managing cloud drive files and objects.  
Read more about Rclone here -- https://rclone.org.  The Rclone gui is useful for building and testing Rclone configs.  Also simple file
operations can be done through the interface.  Additionally the app can be used to perform scheduled operations.  The intent is to run 
rclone (command line version) on a regular schedule once the configs are tested and built.

# Rclone Browser Basic Operation
Find the Rclone Browser app in the Cohesity Marketplace (https://marketplace.cohesity.com).  
Use standard methods to load and run Rclone Browser onto a Cohesity Cluster.
Access the app home page using the launch link from "My Apps / Rclone Browser"
From here, use standard Rclone GUI workflows to manage cloud drive data (https://rclone.org/gui/)

Tip -- When registering a Cohesity S3 server, use these settings
* [svr7admin]
* type = s3
* provider = Other
* access_key_id = ...
* secret_access_key = ...
* endpoint = svr7.cohesity.com:3000
* v2_auth = true

Tip -- basic copy operations can be completed using drag and drop in the side by side view in the Explorer section

Tip -- the app currently does not persist configs between instances.
To save config files to a client workstation use the Explorer section
* Open a pane to the local filesystem (use a local filesystem config)
* Browser to /root/.config/rclone
* Look for and download a file named rclone.config

To load previously saved configs, 

     * ssh into the cohesity cluster
     * move to the bash shell
     * open the firewall to allow client workstation access to the k8s dashboard
     * sudo firewall-cmd --ipset=cluster_ipset --add-entry=1.1.1.1  (of course, replace 1.1.1.1 with the ip address of your client workstation)
     * In your web browser go to "<node_ip>:63773"
     * Use the k8s dashboard to open a terminal windows to the alphine-rclone:latest container
     * Use vi to edit /root/.config/rclone/rclone.config
     * Cut, copy and paste values from the saved config file as needed

Note, the GUI may be a little out of sync when configs are "sideloaded"

Note - cron is available in rclone browser version 1.1 and up.  Version 1.1.2 is due out around 9/15/2020.  Message Steve Klosky if you need it sooner.  Alteratively, load cron by getting to the shell and issuing the "apk update, then apk add cron" commands.

1. To schedule rclone jobs, setup the desired rclone command in the crontab
     * Use Rclone Browser to setup source and target configs
     * Study rclone (cli version) to determine appropriate job command syntax (copy?, sync?, ?)
     * ssh into the cohesity cluster
     * move to the bash shell
     * open the firewall to allow client workstation access to the k8s dashboard
     * sudo firewall-cmd --ipset=cluster_ipset --add-entry=1.1.1.1  (of course, replace 1.1.1.1 with the ip address of your client workstation)
     * In your web browser go to "<node_ip>:63773"
     * Use the k8s dashboard to open a terminal windows to the alphine-rclone:latest container
     * Use vi to edit /etc/crontabs/root
     * Insert new crontab entries or cut, copy and paste values in the crontab file as needed
1. To save the cron config
     * Use Rclone Browser to download the crontab file
     * Setup a config for the local file system
     * Use the Explorer to browse the local system and go to /etc/crontabs
     * Dowload the file named root to a local workstation
     
# Lessons Learned

* Here is a compilation of lessons learned while implementing rclone browser -- https://github.com/cohsk/athena-alpine-rclone/blob/master/Lessons-Learned.MD
     
# Developer notes

Contributions are welcome.  Please contact Steve Klosky -- steve.klosky@cohesity.com for details

This repository holds assets related to rclone gui (https://rclone.org/gui) 
ported to Cohesity's Athena (Marketplace) framework (https://developer.cohesity.com/docs/get-started-apps.html)
production version is here -- https://marketplace.cohesity.com/app-details/rclone-browser

The project used alpine linux "alpine:latest" docker/container image

This is experimental and intended for educational purposes

The respository is to help coordinate versions and issues

How to Build the "Rclone Browser for Cohesity Marketplace" App

You can download the prebuilt package from the Cohesity Marketplace.

If you are interested in learning how it was build, here are instructions.

In order to build the end product, establish an account on https://devportal.cohesity.com
Contact developer@cohesity.com if you need to get access.

Use the instructions in the docker-image directory to build the image

Please name the image file alpine-rclone:latest

Step 1 -- login to devportal.cohesity.com

Step 2 -- Select "Build an App"

Step 3 -- Select "Container App"

Step 4 -- Fill out the fields in this first web form.  Use the rclone-logo1.svg file from the github deployment directory for the app icon.  Use the screen1 to screen3.svg files for the screenshots.  Most fields should be fairly obvious.  For the Min and Max versions, use 6.2 for Min and "all latest versions".  For now, please leave all App Permissions and App Requirements in the unselected position.  No Additional Json is required.

Step 5 -- Select Next

Step 6 -- Fill out the fields in this second web form.  The fields are self-explanitory.

Step 7 -- Select Next

Step 8 -- On the third web form, use the controls to upload the docker image file (alpine-rclone:latest)

Step 9 -- On the third web form, use the controls to upload or copy/paste the alpine-rclone appspec (https://github.com/cohsk/athena-alpine-rclone/blob/master/deployment/alpine-rclone-appspec.yaml) into the webpage

Step 10 -- Select Next

Step 11 -- Accept the Agreement and Select the "Submit for Review" button

Step 12 -- Wait a minute or two

Step 13 -- Move to the next webpage

Step 14 -- Use the "meatball menu" to the right of the app to access and download the app package file

Congratulations, the app package is ready to go

If you have questions, please log an issue here or contact Steve Klosky -- steve.klosky@cohesity.com

Some notes

Currently workng on a Dockerfile to build the container
There is a concept of a dev container (has golang and some other tools)
There is a concept of a prod container (no dev tools, only runtime bits)
