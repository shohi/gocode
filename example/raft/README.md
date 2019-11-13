# raft example

raft example based on [juanpabloaj/tick-time-with-raft](https://github.com/juanpabloaj/tick-time-with-raft) with some modifications.

Features
- provide join API - `/join?addr=[raft-addr]&id=[node-id]` (*)
- provide api for checking leader - `/leader` (+)
- send message to cluster - `/apply?msg=[data]` (+)
- provide leave API  - `/leave` (+)

## Usage

```
cd raft/

# 1. start node 1
go run . -dir node1 -haddr :9010 -raddr :9011 -id node1

# 2. start node 2
go run . -id node2 -dir node2 -haddr :9020 -raddr :9021 -join localhost:9010

# That's enough! Watch the console and check the log to figure out how raft works!

```


## Reference

1. <https://github.com/hashicorp/raft>

2. <https://github.com/juanpabloaj/tick-time-with-raft>

