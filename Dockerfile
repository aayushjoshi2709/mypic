FROM node:22-alpine AS build
WORKDIR /app
COPY ./client/package*.json ./
RUN npm install
COPY ./client/ ./
RUN npm run build

# build go server
FROM golang:1.25-alpine AS server-build
WORKDIR /app
COPY ./server/go.mod ./
COPY ./server/go.sum ./
RUN go mod download
COPY ./server/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# final image
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/dist ./public
COPY --from=server-build /app/server ./
EXPOSE 8080
CMD ["./server"]
