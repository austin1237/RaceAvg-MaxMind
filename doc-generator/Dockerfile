FROM node:8.6

# Create app directory
WORKDIR /app

# For npm@5 or later, copy package-lock.json as well
# COPY package.json package-lock.json ./
RUN npm install apidoc -g

# Bundle app source
COPY ./mmapi .