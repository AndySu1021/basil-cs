FROM node:lts-alpine as build

WORKDIR /app

COPY package*.json ./

RUN yarn

COPY . .

RUN yarn build:dev

FROM nginx:stable-alpine

COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]