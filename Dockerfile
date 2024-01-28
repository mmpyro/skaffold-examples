# Use an official Node.js runtime as a base image
FROM alpine

# Set the working directory in the container
WORKDIR /usr/src/app

RUN apk update && apk add curl npm nodejs

# Copy package.json and package-lock.json to the working directory
COPY ./src/package*.json ./

# Install the application dependencies
RUN npm install

# Copy the application code to the container
COPY ./src/ .

# Expose the port that the application will run on
EXPOSE 3000

# Define the command to run your application
CMD ["node", "app.js"]
