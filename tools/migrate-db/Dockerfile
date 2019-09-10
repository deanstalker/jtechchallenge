FROM node:carbon

RUN mkdir -p /usr/src/migration
RUN chmod -R 777 /usr/src/migration

WORKDIR /usr/src/waitforit

RUN git clone https://github.com/vishnubob/wait-for-it.git

EXPOSE 8080
CMD ["nodemon","server"]