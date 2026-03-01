L2 Chat
=======

Experimental chat based on low-level packets bypassing IP stack.


## Plan

- Simple test:
  - [x] start server on selected device (TBD: all devices?)
  - [x] send HELLO message to the networks
  - [x] process HELLO message from other apps (TBD: need UID for each instance)
  - [x] use BPF `ether proto 0xABC`
  - [ ] rework `PacketSource` to `ReadPacketData` and `NewDecodingLayerParser`
- Track other instances presence (by periodic HELLO messages):
  - [x] process HELLO and BYE
  - [ ] make user OFFLINE after timeout
- Encrypt messages (TBD: clients must share the same keys)
- TBD: Store chat history in Redis
- Host simple web server for chat access
- Support `bridge` mode to transfer packets using UDP stack


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