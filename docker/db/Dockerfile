FROM mysql:5.7
CMD ["mysqld", "--character-set-server=utf8", "--collation-server=utf8_unicode_ci", "--general_log=1"]
ADD *.sql /docker-entrypoint-initdb.d/