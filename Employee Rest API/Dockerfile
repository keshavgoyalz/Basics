FROM golang:1.16-alpine
WORKDIR /app
RUN go mod download
COPY . .
RUN go build -o EmployeeRestApi .
EXPOSE 8080
CMD ["./EmployeeRestApi"]

