FROM node:carbon-alpine
LABEL maintainer="@douglaspands"
LABEL version="1.0.0"
ENV NODE_ENV=production
COPY . /var/www
WORKDIR /var/www
RUN npm i
CMD npm start