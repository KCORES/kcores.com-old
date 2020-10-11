# Dockerfile

# base info
FROM nginx:latest
MAINTAINER karminski <code.karminski@outlook.com>
USER root

# init
RUN mkdir /data/tmp/nginx/ -p

# copy repo to /data/repo
COPY . /data/repo/kcores.com/

# copy config 
RUN rm -f /etc/nginx/conf.d/default.conf
RUN ls -alh /data/repo/kcores.com/config/nginx/
RUN ln -s /data/repo/kcores.com/config/nginx/kcores.com.public.conf /etc/nginx/conf.d/



# define health check
HEALTHCHECK --interval=5s --timeout=3s CMD curl -fs http://127.0.0.1:80/nginx-status?src=docker_health_check -H"Host:kcores.com" || exit 1


# run nginx
EXPOSE 80/tcp
STOPSIGNAL SIGTERM
CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
