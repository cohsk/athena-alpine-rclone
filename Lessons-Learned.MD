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
     
In some cases, you may want to access the k8s cluster from an external machine.  In order to do this, open the firewall and use port 25689 to hit
the k8s endpoint
* open the firewall to allow client workstation access to the k8s dashboard
* sudo firewall-cmd --ipset=cluster_ipset --add-entry=1.1.1.1  (of course, replace 1.1.1.1 with the ip address of your client workstation)
* install kubectl on your workstation -- https://kubernetes.io/docs/tasks/tools/install-kubectl/
* note that you have to check the version of k8s server running.  Do this from a cluster node -- kubectl.sh version
* follow this article and adjust your $HOME/.kube/config file point at the Cohesity k8s cluster
* https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/
* your config file should look something like this
C:\Users\Administrator>kubectl version
Client Version: version.Info{Major:"1", Minor:"17", GitVersion:"v1.17.0", GitCommit:"70132b0f130acc0bed193d9ba59dd186f0e634cf", GitTreeState:"clean", BuildDate:"2019-12-07T21:20:10Z", GoVersion:"go1.13.4", Compiler:"gc", Platform:"windows/amd64"}
Unable to connect to the server: dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.

C:\Users\Administrator>kubectl config set-cluster cohk8s --server=172.16.3.101:25689
Cluster "cohk8s" set.

C:\Users\Administrator>kubectl config set-context cohcontext --cluster=cohk8s
Context "cohcontext" created.

C:\Users\Administrator>kubectl config use-context cohcontext
Switched to context "cohcontext".

Don’t forgot to open firewall

C:\Users\Administrator>ssh cohesity@cohesity-01.talabs.local
The authenticity of host 'cohesity-01.talabs.local (172.16.3.101)' can't be established.
ECDSA key fingerprint is SHA256:F9Qu09+Dg4B0Ev4UjkmZ1QcVjv7yKe/m2K0uGiNFv0w.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added 'cohesity-01.talabs.local,172.16.3.101' (ECDSA) to the list of known hosts.
*****************  Welcome to Cohesity ******************

WARNING: Unauthorized access to this system is forbidden.
By accessing this system, you agree that your actions
may be monitored if unauthorized usage is suspected.

*********************************************************
cohesity@cohesity-01.talabs.local's password:
Last login: Fri May 29 08:54:05 2020 from 10.0.2.2

Welcome to Cohesity OS!

This Linux is carefully configured and tuned to work with the Cohesity
software. Due to the distributed nature of the Cohesity product, all
nodes are managed by a central configuration manager. Configuration
changes must be done only through the Cohesity UI or CLI, iris_cli.

DO NOT make changes to the Linux OS, including but not limited to:
   - the disk subsystem
   - the Linux kernel
   - the Linux configuration
   - the files under /etc

Any manual changes may cause PERFORMANCE PROBLEMS, CLUSTER FAILURE,
and/or DATA LOSS!

PLEASE CONTACT COHESITY SUPPORT IF YOU FEEL CHANGES ARE NECESSARY.

Version: 6.5.0a_p1_release-20200529_511a508f
 Host: cohesity-01-005056010101-node-1, 172.16.3.101
 Type 'h' for help.
cohesity_shell# s
Enter password for user 'cohesity'
Password:
Last login: Sat Oct 17 07:52:28 PDT 2020 from 172.16.3.9 on pts/0
 Running bash from SSH_CONNECTION: 172.16.3.9 50014 172.16.3.101 22
 Commands are being logged. Precede log comments with '#'.
[cohesity@cohesity-01-005056010101-node-1 ~]$ sudo firewall-cmd --ipset=cluster_ipset --add-entry=172.16.3.9
success
[cohesity@cohesity-01-005056010101-node-1 ~]$ sudo firewall-cmd --ipset=cluster_ipset --add-entry=172.16.3.10
success
[cohesity@cohesity-01-005056010101-node-1 ~]$ exit
logout
 Exiting bash from SSH_CONNECTION: 172.16.3.9 50014 172.16.3.101 22
cohesity_shell# exit
Connection to cohesity-01.talabs.local closed.

C:\Users\Administrator>kubectl get pods -A
NAMESPACE     NAME                                    READY   STATUS    RESTARTS   AGE
kube-system   kubernetes-dashboard-54fb766c84-x5g7w   1/1     Running   2          86d

C:\Users\Administrator>type c:\users\administrator\.kube\config
apiVersion: v1
clusters:
- cluster:
    server: 172.16.3.101:25689
  name: cohk8s
contexts:
- context:
    cluster: cohk8s
    user: ""
  name: cohcontext
current-context: cohcontext
kind: Config
preferences: {}
users: null

C:\Users\Administrator>


