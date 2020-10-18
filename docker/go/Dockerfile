FROM golang:latest
#ARG PORT
#ENV PORT $PORT

# Install packages here
RUN mkdir /go/src/myapp && cd /go/src/myapp &&  \
    go get github.com/joho/godotenv

EXPOSE 8070
