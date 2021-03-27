FROM node:latest AS builder
WORKDIR /build
COPY . .
RUN npm run build

FROM nginx:latest
COPY --from=builder /build/dist /usr/share/nginx/html
