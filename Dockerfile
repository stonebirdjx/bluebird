FROM centos:centos7.9.2009

LABEL maintainer="stonebirdjx <1245863260@qq.com>"

WORKDIR /bluebird

COPY ./output /bluebird/

CMD [ "bash", "bootstrap.sh" ]
