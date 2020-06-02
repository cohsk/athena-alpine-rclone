from golang:alpine

add code to pull rclone from rclone.org
make sure to run rclone-ui so that the gui lands on the image

RUN apk add git
CD /go/src
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
