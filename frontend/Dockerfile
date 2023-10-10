FROM node:16-alpine as react-build

WORKDIR /app

COPY . ./
RUN yarn
RUN yarn build

CMD ["yarn", "start"]
# FROM nginx:alpine
# COPY nginx.conf /etc/nginx/conf.d/default.conf
# COPY --from=react-build /app/build /usr/share/nginx/html
# EXPOSE 80
# CMD ["nginx", "-g", "daemon off;"]
# FROM nginx:1.19-alpine
# COPY --from=build-step /usr/src/app/build /usr/share/nginx/html
# EXPOSE 80

# CMD ["nginx", "-g", "daemon off;"]