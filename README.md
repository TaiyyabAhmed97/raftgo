# MY attempt at implementing raft in process

for learning purposes

some docs:
## Go Raft implementation

* I want to implement raft locally in golang
* to do this, I want to start with nodes as processes instead of other servers/computers over the network.


### Plan
* RAFT-0 - set up empty go project :woohoo
* RAFT-1 - spawn multiple processes of the raft binary by manually running the binary in different terminal windows :woohoo
* RAFT-2 - send messages viw RPC in a all -> all fashion. e.g if process one sends a message, each process recieved and logs the message.
* RAFT-3 - Implement leader election (1 of the MAIN stories of this epic).