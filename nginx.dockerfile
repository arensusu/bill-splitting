
FROM nginx:latest

RUN apt-get update
RUN apt-get install certbot python3-certbot-nginx

COPY nginx.conf /etc/nginx/conf.d/default.conf
