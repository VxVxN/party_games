FROM node:18-alpine AS build

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install --frozen-lockfile

COPY . .

RUN npm run build

FROM nginx:alpine

# 8. Копируем собранные файлы в Nginx
COPY --from=build /app/build /usr/share/nginx/html

# 9. Копируем кастомный конфиг Nginx (если нужно)
COPY nginx.conf /etc/nginx/nginx.conf

# 10. Открываем порт 80
EXPOSE 80

# 11. Запускаем Nginx
CMD ["nginx", "-g", "daemon off;"]
