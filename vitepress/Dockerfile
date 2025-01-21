FROM node:18-alpine

WORKDIR /app

COPY . .

RUN npm install && npm run build

RUN npm install -g serve

EXPOSE 5000

CMD ["serve", "-s", "docs/.vitepress/dist", "-l", "5000"]