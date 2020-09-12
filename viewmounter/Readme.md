6/20/2020 -- This directory is currently old and not used
Currently putting a programatic loading of a view mounter on hold
Learned that the Cohesity App API gives limited, read only access to views
For this utility, read-only gives limited access

9/12/20 -- for now I've been using the local file system on the docker container
Also, in the docker container filesystem, there is an ephemeral mounted view under
/cohesity/mounts/ephemeral.  Somewhat useful.  As time permits, will look at using
persistent views and athena mount.  May or may not be needed.
