FROM golang:latest
RUN mkdir /app
ADD main /app/
EXPOSE 8000
ENTRYPOINT ["/app/main"]
