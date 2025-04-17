FROM ubuntu:latest

WORKDIR /app


RUN apt update && apt install -y iproute2 mysql-client




# Copier le binaire et les certificats SSL/TLS
COPY forum-app /app

COPY backend/server/ssl_tls /app/ssl_tls

# Donner les permissions d'exécution au binaire
RUN chmod +x /app/forum-app

# Exécuter l'application
CMD ["/app/forum-app"]

EXPOSE 443