FROM golang:1.10

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && chmod +x /usr/local/bin/dep
RUN apt-get update

RUN mkdir -p /go/src/github.com/user/mmapi
COPY . /go/src/github.com/user/mmapi
WORKDIR /go/src/github.com/user/mmapi

# Install go dependencies
RUN dep ensure 

# Couldn't get this to load with dep for some reason
RUN go get "github.com/stretchr/testify"

RUN go install .

CMD if [ ${APP_ENV} = "local" ]; \
	then \
    dep ensure && \
	go get github.com/pilu/fresh && \
	fresh; \
	else \
    mmapi; \
	fi