FROM golang:1.22

WORKDIR /app
RUN go mod init user
RUN  go get github.com/gin-gonic/gin 
RUN go get gorm.io/driver/sqlite
RUN  go get gorm.io/gorm
COPY . .
EXPOSE 8080
CMD ["go", "run", "main.go"]