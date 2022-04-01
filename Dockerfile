FROM golang:1.17-buster
WORKDIR /src
COPY tf-linter /usr/bin/tf-linter
ENTRYPOINT ["/usr/bin/tf-linter"]
CMD ["./..."]
