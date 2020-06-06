
# This is a dev / test utility for view-mounter
# It loads some programs into the golang:alpine container
# Nano, Git, Curl
# It pulls the source code for view-mounter
# It logs into the Cohesity Cluster to generate an Access Token
# It puts the access token and several other variables into
# Environment variables to simulate Athena loading
#
# This program assumes that prior to running, an golang:alpine
# container has been loaded into the Athena/k8s system on the cluster
#
# Some helpful commands for loading golang:alpine on Athena
# 
# sudo firewall-cmd --ipset=cluster_ipset --add-entry=$(echo $SSH_CLIENT | awk '{print $1}')
# 172.16.3.101:63773
# kubectl.sh run -i --tty --rm rclone --image=golang:alpine --restart=Never -- sh
# sudo docker container ls
# sudo docker exec -it {container id} sh
#

apk update
apk add nano
apk add git
apk add curl

curl -X POST -k --url 'https://172.16.3.101/irisservices/api/v1/public/accessTokens' -H 'Accept: application/json' -H 'Content-type: application/json' --data-raw '{"password": "admin","username": "admin"}'
