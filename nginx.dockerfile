
FROM nginx:latest

RUN apt-get update -y
RUN apt-get install -y certbot python3-certbot-nginx

COPY nginx.conf /etc/nginx/conf.d/default.conf
