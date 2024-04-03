# https://medium.com/@shahneel2409/containerize-your-go-application-in-7%EF%B8%8F%E2%83%A3-fun-and-easy-steps-72ea1001390©
FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY ./database/*.go ./database/
COPY ./models/*.go ./models/
COPY ./utils/*.go ./utils/
COPY ./server/ ./server/


USER 1000

EXPOSE 8080

CMD ["/backend"]