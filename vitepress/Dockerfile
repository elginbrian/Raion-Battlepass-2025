# Use Node.js image
FROM node:18-alpine

# Set working directory
WORKDIR /app

# Copy Vitepress files
COPY . .

# Install dependencies and build the site
RUN npm install && npm run build

# Use a lightweight HTTP server to serve static files
RUN npm install -g serve

# Expose HTTP server port
EXPOSE 5000

# Command to serve the built site
CMD ["serve", "-s", "docs/.vitepress/dist", "-l", "5000"]