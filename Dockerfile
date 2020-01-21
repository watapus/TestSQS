FROM golang:1.13.6-stretch

COPY TESTSQS /TESTSQS

WORKDIR /

EXPOSE 80

ENTRYPOINT ["./testsqs", "--port", "80"]