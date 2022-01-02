FROM node:16-alpine3.13
WORKDIR /app
COPY . .
RUN npm install && npm run build

FROM nginx:alpine
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*

ENTRYPOINT ["nginx", "-g", "daemon off;"]