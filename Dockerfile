FROM debian:stable-slim
ENV PORT=8080
RUN apt-get update && apt-get -y install ca-certificates

COPY notely /usr/bin/notely

CMD ["/usr/bin/notely"]
