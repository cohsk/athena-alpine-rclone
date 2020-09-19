Q: Can we create separate multiple rclone instances in a single cohesity CE cluster or is it going to be a 1:1 effort?

A: Yes, multiple instances can run on the same cluster.  Each instance will be assigned a unique port for GUI access.  Each will run in a separate container for cli access.

There are some considerations regarding how many instances.

When instances are launched, the Cohesity system checks the currently available and allocated system resources (mainly memory and cpu).  New instances will launch based on availability.

Additionally, the value of QoS entered when starting the app will influence how many resources are allocated for that instance.

Version 1.0.80 of rclone browser uses 1 container on 1 node of the system.  Version 1.1.2 requests 1 container for every node on the system.  This will impact resource allocation.

As you can see there are many combinations and options.  I recommend you run some trials on a test system to get a sense of your specific system capacity.

Q: Why am I having trouble mounting disks with rclone browser?

A: This feature is not supported.  Athena-alpine-rclone is built on a container architecture within Cohesity's ecosystem.  Currently there is limited ability to mount to containers natively.  Over time this may evolve to support mounting.
