FROM centos:latest
RUN yum install -y docker unzip
COPY docker-builder /home/docker-builder
WORKDIR /home
CMD ./docker-builder
