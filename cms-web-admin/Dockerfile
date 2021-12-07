FROM node

WORKDIR /webAdmin
COPY . .

RUN npm install
RUN npm run build

FROM nginx:alpine
LABEL MAINTAINER="10512203@qq.com"

COPY .docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf
COPY --from=0 /webAdmin/dist /usr/share/nginx/html
RUN cat /etc/nginx/nginx.conf
RUN cat /etc/nginx/conf.d/my.conf
RUN ls -al /usr/share/nginx/html
