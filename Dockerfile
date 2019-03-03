FROM golang:stretch
RUN go get -v -t -d github.com/andrzejd-pl/SimpleRestBlogBackend
WORKDIR /go/src/github.com/andrzejd-pl/SimpleRestBlogBackend
RUN go build github.com/andrzejd-pl/SimpleRestBlogBackend \
    && go install github.com/andrzejd-pl/SimpleRestBlogBackend
CMD ["SimpleRestBlogBackend"]