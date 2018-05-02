'use strict';

var ws = new WebSocket("wss://localhost:8000/wss")
const peerConnection = new RTCPeerConnection();

navigator.mediaDevices.getUserMedia({
    audio: false,
    video: true
})
    .then(function(stream){
        peerConnection.addStream(stream);
        peerConnection.createOffer()
            .then(sdp => peerConnection.setLocalDescription(sdp))
            .then(function() {
                ws.send(JSON.stringify({"type": "call", "content": peerConnection.localDescription}))
            });
    })

ws.onmessage = function(e) {
    console.log("Received Message: " + e.data);
};



const video_array = document.querySelectorAll('video')

for (let i=0;i<2;i++) {
    console.log(video_array[i])
    video_array[i].srcObject
}