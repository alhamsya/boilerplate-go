FROM mysql:5.7.25

COPY schema/ /docker-entrypoint-initdb.d/