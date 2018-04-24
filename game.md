# WebRTC

- WebRTC是一个开源项目，旨在使得浏览器能为实时通信(RTC)提供简单的javascript接口。

- WebRTC 
    - 通过一系列信令建立一个浏览器与浏览器之间(peer-to-peer)的信道。
    - 这个信道可以发送任何数据，而不需要经过服务器。
    - 三个接口
        - MediaStream: 通过MediaStream的API能够通过设备的摄像头及话筒获得视频、音频的同步流
        - RTCPeerConnection: RTCPeerConnection是WebRTC用于构建点对点之间稳定、高效的流传输的组件
        - RTCDataChannel: RTCDataChannel使得浏览器之间（点对点）建立一个高吞吐量、低延时的信道，用于传输任意数据
 
 
 - 通过服务器建立信道
    
    - 用户发现以及通信
    - 信令传输
    - NAT/防火墙穿越
 
# STUN, TURN Servers

# SDP(Session Description Protocol)

# Ice Candidates

# Key Api

- GetUserMedia
> Get audio and video streaming

- RTCPeerConnection
> WebRTC peer object

- SDP
> Session Description Protocol

- ICECandidate
> Interactive Connectivity Establishment

# 创建简单的WebRTC的应用

RTCPeerConnection
