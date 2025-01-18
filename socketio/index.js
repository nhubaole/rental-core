import { Server as SocketIOServer } from "socket.io";
import express from 'express';
import http from 'http';
import loadConfig from './config/config.js';
import AWS from 'aws-sdk';
import {query } from "./config/database.js";
import multer from 'multer';

const app = express();
const server = http.createServer(app);  

let config;
let connectedClients = {};
loadConfig()
  .then((loadedConfig) => {
    config = loadedConfig;
    const serverPort = config.node_server.port || 3001;
    
    const io = new SocketIOServer(server, {
      cors: {
        origin: "*",
        methods: ["GET", "POST"],
      },
    });

    // Configure S3 client
    const s3 = new AWS.S3({
      accessKeyId: config.s3.aws_access_key_id,
      secretAccessKey: config.s3.aws_secret_access_key,
      region: config.s3.region,
    });
    const upload = multer({ storage: multer.memoryStorage() });

    console.log(s3)
    io.on('connection', (socket) => {
      const userId = socket.handshake.query.userID;
      if (userId) {
        connectedClients[userId] = socket.id;
        console.log(`User ${userId} connected with socket ID: ${socket.id}`);
      } else {
        console.log(`User ID not found in query parameters for socket ID: ${socket.id}`);
      }
      
      socket.on('sendMessage', async (message,callback) => {
        console.log(message)
        const { sender_id, receiver_id, conversation_id, file, content, type, rent_auto_content } = message;
      
        try {
          let savedContent = content;

          // If type is 3, handle image upload to S3
          if (type === 3 && file) {
            const buffer = Buffer.from(file.data, 'base64');
            const fileName = `message-images/${Date.now()}-${sender_id}${file.extension}`;
    
             const s3Params = {
              Bucket: "smartrental",
              Key: fileName,
              Body: buffer,
              ContentType: file.mimeType,
            };
    
            const s3Response = await s3.upload(s3Params).promise();
           
            savedContent = s3Response.Location; // URL of the uploaded file
            console.log(savedContent)
          }
          const normalizedRentAutoContent = rent_auto_content || {};
      
          const result = await query(
            'INSERT INTO messages (sender_id, conversation_id, content, type, rent_auto_content) VALUES ($1, $2, $3, $4, $5) RETURNING *',
            [sender_id, conversation_id, savedContent, type, normalizedRentAutoContent]
          );
      
          const savedMessage = result.rows[0];
          console.log(savedMessage);
      
          io.to(connectedClients[receiver_id]).emit('receiveMessage', savedMessage);
          io.to(connectedClients[sender_id]).emit('receiveMessage', savedMessage);
      

          
          await query(
            'UPDATE conversations SET last_message_id = $1 WHERE id = $2',
            [savedMessage.id, savedMessage.conversation_id]
          );

          if (callback) {
            callback({ success: true, message: savedMessage });
          }
        } catch (error) {
          console.error(`Failed to save message from ${sender_id} to ${receiver_id}:`, error);
        }
      }); 
      
      socket.on('disconnect', () => {
        console.log('User disconnected:', socket.id);
        delete connectedClients[socket.id];
      });
    });

    server.listen(serverPort, () => {
      console.log(`Listening on port ${serverPort}`);
    });
  })
  .catch((error) => {
    console.error(`Failed to load configuration: ${error}`);
  });

app.use(express.json());
