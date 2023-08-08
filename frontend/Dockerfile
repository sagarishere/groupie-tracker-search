# --- Stage 1: Build stage ---
FROM node:alpine as builder
WORKDIR /groupie
# Copy project files including package.json and package-lock.json
COPY . .

# Install dependencies and build the app
RUN npm install -g npm@9.8.1
RUN npm ci && npm run build

# --- Stage 2: Run stage ---
FROM node:alpine
WORKDIR /app
# Copy only the built application from stage 1
COPY --from=builder /groupie/ ./
# Expose port
EXPOSE 3000
# Run the web service
ENTRYPOINT ["npm", "run", "start"]
