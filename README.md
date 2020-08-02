# athena-alpine-rclone (aka "Rclone Brower")

This repo is for the Cohesity Rclone Browser Marketplace app

# What Rclone Browser does
This app creates a web based user interface to Rclone.  Rclone is a well know utility for managing cloud drive files and objects.  
Read more about Rclone here -- https://rclone.org.  The Rclone gui is useful for building and testing Rclone configs.  Also simple file
operations can be done through the interface.  Additionally the app can be used to perform scheduled operations.  The intent is to run 
rclone (command line version) on a regular schedule once the configs are tested and built.

# developer notes
This repository holds assets related to rclone gui (https://rclone.org/gui) 
ported to Cohesity's Athena (Marketplace) framework (https://developer.cohesity.com/docs/get-started-apps.html)

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
