FROM node:14

WORKDIR /client

COPY /package*.json ./

RUN npm install

RUN npm install -g redoc-cli

COPY /client ./

EXPOSE 3000

CMD ["npm", "start"]