FROM node:alpine3.19

WORKDIR /app

COPY frontend/package*.json ./

RUN npm install

COPY frontend/ .

RUN npm run build

EXPOSE 3000

CMD [ "npm", "start" ]