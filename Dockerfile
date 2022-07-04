# Building the binary of the App
FROM golang:alpine
RUN apk add build-base


WORKDIR /src

# Copy all the code to compile everything
COPY . .

# Downloads all the dependencies in advance
RUN go mod download

WORKDIR /src/cmd/web


# Builds the application as statically linked
RUN CGO_ENABLED=1 GOOS=linux go build -o app -a -ldflags '-linkmode external -extldflags "-static"' .


# Expose port 4000
EXPOSE 4000

CMD ["./app"]