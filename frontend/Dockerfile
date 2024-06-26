# FROM node:lts-alpine

# # pasang (install) server http sederhana untuk menjalankan static content
# RUN npm install -g http-server

# # buat folder 'app' pada direktori yang sedang dikerjakan
# WORKDIR /app

# # salin 'package.json' dan 'package-lock.json' (jika ada)
# COPY package*.json ./

# # pasang dependecy proyek
# RUN npm install

# # salin berkas-berkas proyek serta folder-foldernya ke direktori yang sedang dikerjakan (misal. folder 'app)
# COPY . .

# # bangun aplikasi untuk produksi dengan minifikasi
# RUN npm run build

# EXPOSE 8080
# CMD [ "http-server", "dist" ]

# tahap pengembangan
FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# tahap produksi
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY --from=build-stage /app/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]