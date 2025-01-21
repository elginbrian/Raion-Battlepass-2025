FROM node:18-alpine

WORKDIR /app

COPY . .

RUN npm install vitepress --save-dev

RUN npm install

EXPOSE 5173

CMD ["npm", "run", "docs:dev", "--", "--host"]