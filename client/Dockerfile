FROM node:lts-alpine3.14

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY ./src ./
COPY ./public ./
COPY ./package.json ./
RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent
EXPOSE 3000
# start app
CMD ["npm", "start"]