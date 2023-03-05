FROM scratch
MAINTAINER Jazee
WORKDIR /go/src/
COPY . .
EXPOSE 8081
CMD ["/bin/bash", "/go/src/build.sh"]