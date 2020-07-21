# SERVICE can be one of: "orders", "products" or "auth"
FROM golang AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG SERVICE=orders
RUN echo Building image for... $SERVICE
ENV CGO_ENABLED=0
RUN mkdir -p /artifacts
RUN go build -o /artifacts/server ./$SERVICE/cmd/server

FROM scratch
WORKDIR /app/bin
COPY --from=builder /artifacts/server .
CMD [ "/app/bin/server" ]
