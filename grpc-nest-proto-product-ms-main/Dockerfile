# Use an official Node.js runtime as the base image
FROM node:18

# Set the working directory in the container to /app
WORKDIR /usr/src/app

# Copy package.json and package-lock.json into the directory /app in the container
COPY package*.json ./

# Añadir las claves públicas antes de apt-get update
RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 0E98404D386FA1D9 6ED0E7B82643E131 F8D2585B8783D481 54404762BBB6E853 BDE6D2B9216EC7A8
RUN apt-get update && apt-get install -y protobuf-compiler

# Install any needed packages specified in package.json
RUN npm install

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
CMD [ "node", "dist/main.js" ]
