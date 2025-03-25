FROM node:23.10.0-bookworm

WORKDIR /app
COPY . .

RUN npm install
RUN npm run build

CMD ["npm", "start"]