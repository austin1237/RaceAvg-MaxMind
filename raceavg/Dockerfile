FROM golang:1.10


RUN apt-get update

RUN mkdir -p /go/src/github.com/user/raceavg
COPY . /go/src/github.com/user/raceavg
WORKDIR /go/src/github.com/user/raceavg


# Couldn't get this to load with dep for some reason
RUN go get "github.com/stretchr/testify"

RUN go install .

CMD if [ ${APP_ENV} = "local" ]; \
	then \
	go get github.com/pilu/fresh && \
	fresh; \
	else \
    raceavg; \
	fi