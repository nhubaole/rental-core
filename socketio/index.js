const app = require('express')();
const server = require('http').createServer(app);
const io = require('socket.io')(server, {
  transports: ['polling'],  // Cấu hình server để hỗ trợ polling
  cors: {
    origin: "*",
    methods: ["GET", "POST"]
  }
});

const port = 3000;

io.on('connection', (socket) => {
    socket.emit('connection', {message: 'hi'})
    socket.on('connection', (data) => {
        console.log(data)
    })
    socket.on('disconnect', function () {
        console.log('user disconnected');
      });
})

server.listen(port, function() {
  console.log(`Listening on port ${port}`);
});