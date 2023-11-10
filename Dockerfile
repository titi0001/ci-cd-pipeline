FROM ubuntu:latest

EXPOSE 8000

ENV HOST=localhost 
ENV PORT=5432
ENV USER=root 
ENV PASSWORD=root 
ENV DBNAME=root

WORKDIR /app

COPY ./main main

CMD ["./main"]

