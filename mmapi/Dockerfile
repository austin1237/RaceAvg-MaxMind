FROM golang:1.10

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && chmod +x /usr/local/bin/dep
RUN apt-get update

# Grab csv files
RUN mkdir -p /maxmind
RUN wget -P /maxmind "https://s3.amazonaws.com/maxmind-csvs/GeoLiteCity-Blocks.csv"
RUN wget -P /maxmind "https://s3.amazonaws.com/maxmind-csvs/GeoLiteCity-Location.csv"


RUN mkdir -p /go/src/github.com/user/mmapi
COPY . /go/src/github.com/user/mmapi
WORKDIR /go/src/github.com/user/mmapi

# Install go dependencies
RUN dep ensure 
RUN go install .

CMD if [ ${APP_ENV} = "local" ]; \
	then \
    dep ensure && \
	go get github.com/pilu/fresh && \
	fresh; \
	else \
    mmapi; \
	fi