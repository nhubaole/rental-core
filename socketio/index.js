import {Server as SocketIOServer} from "socket.io"
const path = require('path');
const yaml = require('node-yaml');
import {server} from "express"
const configPath = path.join(__dirname, '../configs', 'local.yaml');
const io = SocketIOServer(server, {
  cors: {
    origin: "*",
    methods: ["GET", "POST"]
  }
});

yaml.read(configPath)
  .then(config => {
    console.log('Loaded configuration:', config);

    // Sử dụng cấu hình từ file YAML
    const serverPort = config.node_server.port;
    const dbConfig = config.db;
    const redisConfig = config.redis;
    const s3Config = config.s3;

    console.log(`Server running on port: ${serverPort}`);
    console.log(`DB Configuration:`, dbConfig);
    console.log(`Redis Configuration:`, redisConfig);
    console.log(`S3 Configuration:`, s3Config);
  })
  .catch(error => {
    console.error(`Error reading or parsing YAML file: ${error}`);
  });


const port = serverPort || 3001;


app.use(express.json());

let connectedClients = {};

io.on('connection', (socket) => {
  const userId = socket.handshake.query.user_id;
  if (userId) {
    userSocketMap[userId] = socket.id;
    console.log(`User ${userId} connected with socket ID: ${socket.id}`);
  } else {
    console.log(`User ID not found in query parameters for socket ID: ${socket.id}`);
  }
  
  socket.on('disconnect', () => {
    console.log('User disconnected:', socket.id);
    delete connectedClients[socket.id];
  });
});

app.post('/send-message', (req, res) => {
  const { user_id, message } = req.body;

  if (!message || !user_id) {
    return res.status(400).json({ error: 'Message and user_id are required' });
  }

  const socket_id = userSocketMap[user_id];

  if (socket_id) {
    io.to(socket_id).emit('connection', { message });
    console.log(message)
    res.status(200).json({ success: true, message: 'Message sent successfully', data: user_id });
  } else {
    res.status(404).json({ error: 'Socket not found or user is not connected' });
  }
});

server.listen(port, () => {
  console.log(`Listening on port ${port}`);
});
