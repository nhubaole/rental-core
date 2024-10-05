const express = require('express');
const http = require('http');
const socketIO = require('socket.io');
  
const app = express();
const server = http.createServer(app);
const io = socketIO(server, {
  cors: {
    origin: "*",
    methods: ["GET", "POST"]
  }
});
 

const port = 3000;

app.use(express.json());

let connectedClients = {};

io.on('connection', (socket) => {
  console.log('A user connected:', socket.id);
  
  connectedClients[socket.id] = socket;

  socket.on('disconnect', () => {
    console.log('User disconnected:', socket.id);
    delete connectedClients[socket.id];
  });

  socket.on('send-to-server', (data) => {
    console.log(`Message from ${socket.id}:`, data);
  });
});

app.post('/send-message', (req, res) => {
  const { message, socket_id } = req.body;

  if (!message || !socket_id) {
    return res.status(400).json({ error: 'Message and socket_id are required' });
  }

  const targetSocket = connectedClients[socket_id];

  if (targetSocket) {
    targetSocket.emit('connection', { message });
    console.log(message)
    res.status(200).json({ success: true, message: 'Message sent successfully' });
  } else {
    res.status(404).json({ error: 'Socket not found or user is not connected' });
  }
});

server.listen(port, () => {
  console.log(`Listening on port ${port}`);
});
