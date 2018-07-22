FROM golang

# Fetch dependencies
RUN go get github.com/tools/godep

# Add project directory to Docker image.
ADD . /go/src/github.com/cwise/tfsa

ENV USER cwise
ENV HTTP_ADDR :8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET xsOV25ZJi0fO33ls

# Replace this with actual PostgreSQL DSN.
ENV DSN postgres://cwise@localhost:5432/tfsa?sslmode=disable

WORKDIR /go/src/github.com/cwise/tfsa

RUN godep go build

EXPOSE 8888
CMD ./tfsa