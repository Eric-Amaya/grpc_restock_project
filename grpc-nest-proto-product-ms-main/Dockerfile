# Use an official Node.js runtime as the base image
FROM node:20

# Set the working directory in the container to /app
WORKDIR /app

# Copy package.json and package-lock.json into the directory /app in the container
COPY package*.json ./

# Install any needed packages specified in package.json
RUN npm install

# Install protoc (Protocol Buffers compiler)
RUN apt-get update && apt-get install -y \
    protobuf-compiler

# Bundle the app source inside the Docker image
COPY . .

# Run proto install and generate TypeScript files from proto files
RUN npm run proto:install
RUN npm run proto:product

# Make port 5053 available to the world outside this container
EXPOSE 5053

# Compila la aplicación NestJS
RUN npm run build

# Ejecuta la aplicación cuando se inicie el contenedor
CMD ["npm", "run", "start"]