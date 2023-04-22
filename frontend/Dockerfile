FROM node:18.14.2-alpine3.17

WORKDIR /pear-open-app

COPY . .

RUN npm ci

EXPOSE 3000
CMD ["npm", "start"]
