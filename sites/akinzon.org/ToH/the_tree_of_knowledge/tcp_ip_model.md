# TCP IP model

Up: [Networks](networks)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Networks

### Network Access Layer

Deals with the local link a host is connected to

Each host has a 48-bit [MAC address](mac_address) that is unique in the link.

![Pasted image 20231129184510.png](pasted_image_20231129184510.png)
##### Physical connection standards:
 > LAN (IEEE 802 family) 
 > Ethernet (IEEE 802.3)
 > Wireless LAN (IEEE 802.11)
 > Wireless PAN (IEEE 802.15)
 

### In-between...

ARP (Address Resolution Protocol) - maps IP addresses to MAC addresses (IPv4)

NDP (Neighbour Discovery Protocol) - IPv6 version of ARP

Both operate at the link layer

### Internet layer

**Best efforts** - uses packet (next-hop) routing

**Globally unique addressing** - [IP Addresses|IP addresses](ip_addresses|ip_addresses)

Passes a received packet's payload to the correct transport layer port

Makes it easier to have a connection on different network types

#### ICMP/ICMPv6

Internet control protocol - used to respond to errors in IP operations
	Found in regular IP packet headers
	Used by `ping` and `traceroute` - both of which are good for diagnosing network issues

### Transport layer

Provides **Host-to-host communication** 

#### TCP - Reliability is Key

Connection-oriented: Handles ACK and retransmissions, this is all about how the hosts connect.

Data is guaranteed to arrive in the right order

Good for applications that need reliability

**Flow control** - stops fast sender overwhelming slow receiver
	> Uses sliding windows
	> Sender only sends if there is buffer space to accept data

**Congestion control** - reduces send rate to ease congestion
	The sender starts with small packets, slowly increasing the size until there is packet loss, then restarts transmission with a low packet size

#### UDP - Fire and Forget

Connectionless: Just sends the packets out, no guarantee for a correct packet order, doesn't handle ACKs or retransmissions (the application would have to if needed)

Good for applications that need fast transmission

### Application layer

The software that uses the network

DNS (Domain Name Service) - Maps domain names to an IPv4/6 address
### How naming and addressing works in each layer

Layer | Name/Address type | Example
----------|--------------|---------
Application| Domain name | www.bbc.co.uk
Transport| Port number | 80
Network| IP address | 128.64.56.87
Link layer| MAC address | 00:14:1b:3d:2c:00


Summarise each layer of the TCP/IP stack, from the bottom up #flashcard 
**Network Access Layer** - Deals with the local link a host is connected to and uses MAC addresses
**Internet Layer** - Routes packets and uses [IP addresses](ip_addresses)
**Transport Layer** - Provides host-to-host communication
**Application Layer** - The software that uses the network
<!--ID: 1701371422507-->


What are the main difference between TCP and UDP #flashcard 
TCP focuses on reliability While UDP focuses on speed.
<!--ID: 1701371422513-->


What are flow and congestion control, and how do they work #flashcard 
**Flow control** - Prevents fast sender from overwhelming a slow receiver, data is only sent if there is buffer space to receive it
**Congestion control** - Reduces send rate to ease transmission, slowly increases the packet size until there is too much packet loss and then resets the size. 
<!--ID: 1701371422518-->




























#### Why:
#### How:









