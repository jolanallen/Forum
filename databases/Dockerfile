FROM mysql:8.0

EXPOSE 3306



ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=forum
ENV MYSQL_USER=user
ENV MYSQL_PASSWORD=password1234

COPY forum.sql /docker-entrypoint-initdb.d/forum.sql
