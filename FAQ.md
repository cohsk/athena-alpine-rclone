Q: Can we create separate multiple rclone instances in a single cohesity CE cluster or is it going to be a 1:1 effort?

A: Yes, multiple instances can run on the same cluster.  Each instance will be assigned a unique port for GUI access.  Each will run in a separate container for cli access.

There are some considerations regarding how many instances.

When instances are launched, the Cohesity system checks the currently available and allocated system resources (mainly memory and cpu).  New instances will launch based on availability.

Additionally, the value of QoS entered when starting the app will influence how many resources are allocated for that instance.

Version 1.0.80 of rclone browser uses 1 container on 1 node of the system.  Version 1.1.2 requests 1 container for every node on the system.  This will impact resource allocation.

As you can see there are many combinations and options.  I recommend you run some trials on a test system to get a sense of your specific system capacity.

Q: Why am I having trouble mounting disks with rclone browser?

A: This feature is not supported.  Athena-alpine-rclone is built on a container architecture within Cohesity's ecosystem.  Currently there is limited ability to mount to containers natively.  Over time this may evolve to support mounting.

Q:  Why did my configs disappear?  Where are my config files?

A:  The early versions of rclone browser do not backup the rclone config files automatically.  Please manually backup the rclone config files in the interim. The rclone config file is located in the container file system.

Normally the config file is in your home directory as a file called .config/rclone/rclone.conf (or .rclone.conf if created with an older version). If $XDG_CONFIG_HOME is set it will be at $XDG_CONFIG_HOME/rclone/rclone.conf.

If there is a file rclone.conf in the same directory as the rclone executable it will be preferred. This file must be created manually for Rclone to use it, it will never be created automatically.

If you run rclone config file you will see where the default location is for you.

Q: Does rclone encrypt data in flight?

A: In flight encryption is a function of the cloud storage used.  Rclone supports about 30 types of cloud storage.  The vast majority communicate over https to ensure encryption in flight. When available, rclone uses https.  So, we’ll need to verify on a case by case basis.  

For S3
“HTTPS is supported by S3, to protect data in transit.
Quoting from the Security section of the S3 FAQs: You can securely upload/download your data to Amazon S3 via SSL endpoints using the HTTPS protocol. If you're using the https:// endpoint for S3, then your data in transit should be encrypted properly.Jun 12, 2017”

For GCS
“	• Figure 1 shows an external path (labeled connection D). Typical cases of this kind of routing request are:

		• From a Compute Engine VM to Google Cloud Storage
		• From a Compute Engine VM to a Machine Learning API
From the VM to the GFE, Google Cloud services support protecting these connections with TLS by default2. The connection is authenticated from the GFE to the service and encrypted if the connection leaves a physical boundary.
"

