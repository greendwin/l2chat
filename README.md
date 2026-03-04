L2 Chat
=======

Experimental chat based on low-level packets bypassing IP stack.


## Plan

- ~~Simple test:~~
  - [x] start server on selected device (TBD: all devices?)
  - [x] send HELLO message to the networks
  - [x] process HELLO message from other apps (TBD: need UID for each instance)
  - [x] use BPF `ether proto 0xABC`
  - [x] rework `PacketSource` to `ReadPacketData` and `NewDecodingLayerParser`

- Split CLI to multiple independant options:
  - [ ] `--device ID` - listen L2 network
  - [ ] `--user NAME` - send `HELLO` and participate in chat (chat is optional, we can work as a simple bridge)
  - [ ] `--bridge IP` - resend all packets to other instance on `IP:PORT` (which port?) 
  - [ ] `--web PORT` - host web server with chat iface (TBD: does it force `--user` option? can we change user name or enable chat dinamically? technically we can)

- Web Server:
  - [ ] host basic client
  - [ ] enable / disable online (aka start `HELLO` or `BYE`)
  - [ ] rename our user
  - [ ] send messages
  - [ ] show chat and system events 
  - [ ] show available users in the chat

- Bridge:
  - TBD: all instances must listen specific port, but will ignore OS stack anyway, so it can be any port and UDP packet with magic number to make sure that its our packets only
  - [ ] check `UDP` packets on `7890`
  - [ ] user extra layer with magic number for `UDP` tunneling
  - [ ] filter `UDP` packets with magic number (TBD: do we need to fix port?)
  - [ ] support bridge handshake (we must known that other host exists before routing all trafic there)
  - [ ] resend all L2 Chat traffic to other bridge
  - [ ] resend all traffic to bridge that connected to us
  - [ ] we must stop send traffic when other host down

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