FROM node:22-alpine

WORKDIR /app

COPY package*.json ./
RUN npm install
RUN npm ci
COPY . .
EXPOSE 3000
CMD ["node", "index.js"]