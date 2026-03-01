L2 Chat
=======

Experimental chat based on low-level packets bypassing IP stack.


## Dependencies

### gopacket

Use [gopacket](github.com/google/gopacket) for low-level access to network packets.
On Windows [npcap](https://npcap.com) driver should be installed.

### cobra

For CLI is used [cobra](https://github.com/spf13/cobra).
We need to install its cli for autogen:

```bash
go install github.com/spf13/cobra-cli@latest
```