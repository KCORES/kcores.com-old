# kcores.com.dockerfile

# base info
FROM nginx:latest
MAINTAINER karminski <code.karminski@outlook.com>
USER root

# set env
ENV SERVER_ENV="__SERVER_ENV__"

# copy repo to /data/repo
COPY . /data/repo/kcores.com/

# define health check
HEALTHCHECK --interval=5s --timeout=3s CMD netstat -an | grep 80 > /dev/null; if [ 0 != $? ]; then exit 1; fi;

# run php-fpm
EXPOSE 9000
STOPSIGNAL SIGTERM
ENTRYPOINT ["php-fpm"]