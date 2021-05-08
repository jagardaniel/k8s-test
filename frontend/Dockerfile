# Step 1: Build stage
FROM node:lts AS builder

WORKDIR /build

COPY package.json package-lock.json ./

RUN npm cache clean --force
RUN npm install

COPY . .

RUN npm run build

# Step 2: Build a small image
FROM nginx:1.19-alpine

COPY --from=builder /build/dist/frontend /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
