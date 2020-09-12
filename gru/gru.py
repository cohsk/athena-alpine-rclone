# Project Gru
#
# A multi thread python controller for rclone daemon (rcd)
# intended to be loaded into a Cohesty EasyScript scheduled task
# intended to get a list of rclone operations, then distribute
# the individaul operations to the collection of local rcd minions
# Read about rclone here -- https://rclone.org
# Read about Cohesity EasyScript here -- https://www.cohesity.com/marketplace/easyscript/
# Read about Rclone Browser here -- https://www.cohesity.com/marketplace/rclone-browser/
# Read about remote controlling rcd here -- https://rclone.org/rc/
# Read about building rcd http requests here -- https://forum.rclone.org/t/python-http-post-to-rcd-examples/19071
# Article on python queues here -- https://www.geeksforgeeks.org/queue-in-python/
#
# September 2020
#
# Contributors:
# Steve Klosky
# {enter names here}
#
#
# Imports
import types

# Start

# In this block, setup interesting parameters
threadsPerMinion=1
username="user"
password="pass"

# rclone details
rcJob = types.SimpleNamespace()
rcJob.operation = "Copy"
#rcJob.operation = "Sync"
rcJob.sourceRemote = "source"
rcJob.sourcePath = "sPath"
rcJob.targetRemote = "target"
rcJob.targetPath = "targetPath"

#output option
reportOut = "email"
reportOut = "file:"
reportOut = "email&file:"
reportOut = ""

# Stand up a fifo job queue

# Stand up a fifo reporting queue

# Get an inventory of rcd minions

# Send an entry to the reporting queue saying that we're starting everything

# Get Big Job details

# Make sure source and target are present (and match) on all minions

# Run rclone in list mode to generate the list of little jobs to run

# Submit all little jobs to the job queue

# Spawn thread(s) to handle jobs for each minion.  Note, number of threads per minion will be variable
# so that the user can adjust to match their environment.

###########
#
# thread logic will be somewhere below
#
##########

# Do while job queue length is greater than 0
# Sleep 1
# end Do

# throw an entry in the job queue showing the big job is done and end time

# output report queue entries to email or file or both or none

# kill all threads (I think this is automatic if threads are daemon threads)
# empty the queues  (I think python handles this automatically)
# terminate all processes

# a simple example of multithreading to help illustrate what we will build
#
#from threading import Thread
#import time
#a=1
#
#def thread1(threadname):
#    global a
#    while True:
#        print(a)
#
#def thread2(threadname):
#    global a
#    while True:
#        a += 1
#
#        time.sleep(1)
#
#thread1 = Thread( target=thread1, args=("Thread-1", ) )
#thread2 = Thread( target=thread2, args=("Thread-2", ) )
#
#thread1.start()
#thread2.start()
#
#thread1.join()
#thread2.join()


### start of thread logic ###

# thread startup
## determine my minion
## authenticate to my minion

#####  while the length of the job queue > 0
# grab a little job from the queue
# submit the little job to my minion
# handle response from my minion
## log error or success in reporting queue
## remove item from the job queue
##### end while

### end of thread logic ###
