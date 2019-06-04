FROM mysql:5.7.25

ENV MYSQL_ROOT_PASSWORD=mytest
ADD test.sql /docker-entrypoint-initdb.d/