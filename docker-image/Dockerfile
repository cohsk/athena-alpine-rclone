#this file is totally notional and incomplete.

#use the alpine image supported by the go team
from golang:alpine

#download and install rclone
RUN apk add curl
RUN curl -O https://downloads.rclone.org/rclone-current-linux-amd64.zip
RUN unzip rclone-current-linux-amd64.zip
RUN cd rclone-*-linux-amd64
RUN cp rclone /usr/bin/
RUN chmod 755 /usr/bin/rclone

#make sure to download the rclone-gui module
#code goes here

#get the program that mounts Cohesity views
RUN apk add git
RUN CD /go/src
RUN go get github.com/cohsk/cohesity-appspec/sampleapp/viewbrowser/viewbrowser
RUN go get gopkg.in/yaml.v2
RUN go get cohsk/viewmounter/view-mounter-exec
RUN go build ...
RUN CP view-mounter-exec /usr/bin
RUN chmod 755 /usr/bin/view-mounter-exec

adjust /etc/init.d files so that certain files run when the container starts
copy rclone-ui.start to /etc/init.d
chmod 755 rclone-ui.start
copy view-mounter.start to /etc/init.d
modify the local policy to use the local startup scripts
rc-update add local default
