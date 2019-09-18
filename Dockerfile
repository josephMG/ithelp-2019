FROM golang:alpine
ENV GOBIN /go/bin
ENV GOOGLE_APPLICATION_CREDENTIALS /app/authentication.json

RUN apk add --no-cache git mercurial

WORKDIR /app
COPY get.sh /app

RUN chmod +x /app/get.sh
RUN /app/get.sh
# RUN go get -u -v cloud.google.com/go/vision/apiv1

ADD . /app
RUN cd /app && go build -o app

CMD []
# ENTRYPOINT ./app
