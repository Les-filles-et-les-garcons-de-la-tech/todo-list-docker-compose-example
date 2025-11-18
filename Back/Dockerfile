FROM golang:1.25.4-alpine

WORKDIR /app

COPY app/go.mod app/go.sum ./
RUN go mod download

# On copie les sources Go
COPY app/*.go ./

# Build binaire
RUN go build -o webapp .

ENV PORT=81
EXPOSE 81

CMD ["./webapp"]
