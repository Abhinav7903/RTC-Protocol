const SIGNAL_SERVER_URL = "http://localhost:8194";

const pc = new RTCPeerConnection({
  iceServers: [{ urls: ["stun:stun1.l.google.com:19302"] }],
});

let localStream = null;
let remoteStream = new MediaStream();

let senderId = Math.floor(Math.random() * 100000); // Changed to integer
let roomId = null;

document.getElementById("webcamButton").onclick = async () => {
  try {
    localStream = await navigator.mediaDevices.getUserMedia({ video: true, audio: true });
    localStream.getTracks().forEach((track) => pc.addTrack(track, localStream));
    document.getElementById("webcamVideo").srcObject = localStream;
    document.getElementById("remoteVideo").srcObject = remoteStream;

    pc.ontrack = (event) => {
      remoteStream.addTrack(event.track);
    };

    pc.onicecandidate = async (event) => {
      if (event.candidate && roomId) {
        await sendSignal("ice", event.candidate);
      }
    };

    document.getElementById("callButton").disabled = false;
    document.getElementById("answerButton").disabled = false;
  } catch (error) {
    console.error("Error accessing webcam:", error);
  }
};

// Create call and room
document.getElementById("callButton").onclick = async () => {
  const roomName = document.getElementById("createRoomId").value.trim();
  if (!roomName) return alert("Please enter a Room Name to create a call.");

  // Step 1: Create room via backend
  const res = await fetch(`${SIGNAL_SERVER_URL}/room`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      name: roomName,
      room_type: "single",
    }),
  });

  const roomResponse = await res.json();
  roomId = roomResponse?.data?.data?.id; // Extract numeric ID

  if (!roomId) return alert("Failed to create room");

  // Step 2: Continue signaling
  const offer = await pc.createOffer();
  await pc.setLocalDescription(offer);
  await sendSignal("offer", offer);
  pollSignals();
};

// Answer call
document.getElementById("answerButton").onclick = async () => {
  const roomIdInput = document.getElementById("callInput").value.trim();
  if (!roomIdInput) return alert("Please enter a Room ID to join a call.");

  roomId = parseInt(roomIdInput); // ensure it's numeric
  const signals = await fetchSignals();

  const offerSignal = signals.find(s => s.signal_type === "offer");
  if (!offerSignal) return alert("No offer found for this Room ID.");

  await pc.setRemoteDescription(new RTCSessionDescription(offerSignal.payload));

  const answer = await pc.createAnswer();
  await pc.setLocalDescription(answer);
  await sendSignal("answer", answer);
  pollSignals();
};

// Send signal to Go backend
async function sendSignal(type, payload) {
  if (!roomId) {
    console.warn("No room ID set for signaling.");
    return;
  }

  await fetch(`${SIGNAL_SERVER_URL}/signal`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      room_id: parseInt(roomId),
      sender_id: senderId,
      signal_type: type,
      payload,
    }),
  });
}

// Poll for signals from other peer
function pollSignals() {
  setInterval(async () => {
    if (!roomId) return;

    const signals = await fetchSignals();
    for (const signal of signals) {
      if (signal.signal_type === "answer" && !pc.currentRemoteDescription) {
        await pc.setRemoteDescription(new RTCSessionDescription(signal.payload));
      } else if (signal.signal_type === "ice") {
        await pc.addIceCandidate(new RTCIceCandidate(signal.payload));
      }
    }
  }, 1000); // Poll every second
}

async function fetchSignals() {
  const res = await fetch(
    `${SIGNAL_SERVER_URL}/signal?room_id=${roomId}&sender_id=${senderId}`
  );
  const json = await res.json();
  return json.data || [];
}
