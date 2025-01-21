# Stage 1: Build the Vitepress site
FROM node:18-alpine AS build

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM node:18-alpine

WORKDIR /app

RUN npm install -g serve

COPY --from=build /app/docs/.vitepress/dist /app/docs/.vitepress/dist

EXPOSE 5000

CMD ["serve", "-s", "docs/.vitepress/dist", "-l", "5000"]