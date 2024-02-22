FROM centos:centos7.6.1810

COPY ./linux-timezone /
CMD ["bash","-c","trap 'exit 1' 15; echo $TZ; /linux-timezone; sleep 9999; "]
