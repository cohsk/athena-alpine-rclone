This a collection of lessons learned from testing Rclone browser in the Cohesity Marketplace

1.  Make sure to download and save your configuration file.  The configuration does not persist from pod instance to pod instance
1.  When writing to Cohesity S3, set authorization to v2
1.  When writing to Cohesity S3, don't forget the bucket name
1.  When setting up a local file system as a remote, don't go directly to the explorer.  Setup a remote under the config menu
1.  When setting up GCS, make sure to include the project number from GCP
1.  Here is a video of setting up a local store and a google cloud store -- https://youtu.be/H6noIywIVd0


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

Note - Because of the container based architecture, the mount functions of rclone and rclone browser will not mount
cloud based file systems to the rclone container.

Note - the GUI may be a little out of sync when configs are "sideloaded"

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
     
