FROM nginx:1.25-alpine3.18

RUN rm /etc/nginx/conf.d/default.conf
COPY ./nginx/nginx.conf /etc/nginx/conf.d

COPY ./nginx/certs /etc/nginx/certs
