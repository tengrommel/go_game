# RTC Peer Connection

    const peerConnection = new RTCPeerConnection();
    
    navigator.mediaDevices.getUserMedia({
        audio: true,
        video: true
    })
    .then(function(stream){
        peerConnection.addStream(stream);
        peerConnection.createOffer()
        .then(sdp => peerConnection.setLocalDescription(sdp))
        .then(function() {
            socket.emit('offer', peerConnection.localDescription)
        });
    })
    
    socket.on('answer', function (message){
        peerConnection.setRemoteDescription(message)
    })

# createOffer()

- createOffer creates SDP
- It's promise returns the SDP

## SDP

- SDP is short for Session Description Protocol
- SDP contains everything the peer needs to connect

# setLocalDescription()

- setLocalDescription sets the peerConnection's SDP

# Local Description & SDP

- SDP is short for Session Description Protocol
- SDP contains everything the peer needs to connect
- SDP contains a lot of information

It's important that we get that info over to the peer's browser.

# Signaling

- Signaling is whatever mechanism you use to transport the Local Description & SDP
- WebRTC intentionally does not specify what signaling method to use
- Valid methods include
- SIP over Websockets
- XMPP/Jabber message
- Facebook Messenger
- Email
- Snail Mail
- etc

# I thought WebRTC is P2P?

- The media streaming will be completely peer-to-peer
- How the two peers discover each other is not defined
- Possible you have a novel way for transmitting SDP?

# Browser B

    const peerConnection = new RTCPeerConnection();
    socker.on('offer', function (message){
        peerConnection.setRemoteDescription(message)
        .then(() => peerConnection.createAnswer())
        .then(funtion() {
            socket.emit('answer', peerConnection.localDescription);
        });
    });
    
    peerConnection.onaddstream = function (event) {
        video.srcObject = event.stream
    }
    
# WebRTC in the wild:

- Your device may not have a public IP Address
- STUN: Public Services to ask "What IP and Port am I"
- "Session Traversal Utilities for NAT"
- TURN: STUN server + Media Relay (see RFC5766)
- "Traversal Using Relay around NAT"
- ICE Candidata: Description of connectable IP and Port
- "Interactive Connection Establishment"

# STUN Configuration

    const config = {
        'iceServers': [{
            'urls': ['stun:stun.l.google.com:19302']
        }]
    }
    
# TURN Configuration

    const config = {
        'iceServers': [{
            'urls': ['turn: 54.149.135.227:3478'],
            'username': 'basscord',
            'credential': 'limpbizkitrulez1998',
            'credentialType': 'passwd'
        }]
    }
    
# RTC Configuration
    
    const config = {
        'iceServers': [{
            'urls': ['stun: stun.l.google.com:19302']
        }]
    }
    
    const peerConnection = new RTCPeerConnection(config)
    peerConnection.onicecandidata = function(event) {
        if (event.candidate) {
            socket.emit('candidate', event.candidate)
        }
    };
    
    socket.on('candidate', function (candidate) {
        const c = new RTCIceCandidate(candidate)
        peerConnection.addIceCandidate(c)
    })