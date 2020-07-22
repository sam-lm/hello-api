# TEST

FROM docker.io/library/golang:1.13
WORKDIR /work
COPY . /work/
RUN go build -mod=vendor
CMD ["go","test","-mod=vendor","-v","./..."]

