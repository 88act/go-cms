FROM nginx:alpine
LABEL MAINTAINER="10512203@qq.com"
 
# 将dist文件中的内容复制到 /usr/share/nginx/html/ 这个目录下面
COPY dist/  /usr/share/nginx/html/
#用本地的 default.conf 配置来替换nginx镜像里的默认配置
COPY .docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/default.conf

RUN cat /etc/nginx/nginx.conf
RUN cat /etc/nginx/conf.d/default.conf
RUN ls -al /usr/share/nginx/html


#FROM node

#WORKDIR /go-cms-admin/
#COPY . .
 
#RUN npm config set registry https://registry.npm.taobao.org 
#RUN npm install
#RUN npm run build
