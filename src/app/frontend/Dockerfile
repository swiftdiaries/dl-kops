FROM golang:1.10 as builder

LABEL maintainer="Adhita Selvaraj <adhita.selvaraj@gmail.com>"

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/swiftdiaries/dl-kops/src/app/frontend
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o dl-kops .

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/swiftdiaries/dl-kops/src/app/frontend/ /app
ENV PORT_1=6060
EXPOSE 6060
CMD [ "/app/dl-kops" ]