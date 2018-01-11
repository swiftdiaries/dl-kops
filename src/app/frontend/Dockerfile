FROM golang
LABEL maintainer="Adhita Selvaraj <adhita.selvaraj@gmail.com>"

COPY . /go/src/github.com/swiftdiaries/dl-kops

RUN go install github.com/swiftdiaries/dl-kops/src/app/frontend

ENV PORT_1=8081

WORKDIR /go/src/github.com/swiftdiaries/dl-kops

ENTRYPOINT /go/bin/frontend

EXPOSE 8081