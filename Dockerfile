FROM maven

ENV JAVA_VERSION 8 
RUN \
    echo "deb http://ppa.launchpad.net/webupd8team/java/ubuntu trusty main" > /etc/apt/sources.list.d/webupd8team-java.list  && \
    echo "deb-src http://ppa.launchpad.net/webupd8team/java/ubuntu trusty main" >> /etc/apt/sources.list.d/webupd8team-java.list  && \
    apt-key adv --keyserver keyserver.ubuntu.com --recv-keys EEA14886  && \
    apt-get update  && \
    echo debconf shared/accepted-oracle-license-v1-1 select true | debconf-set-selections  && \
    echo debconf shared/accepted-oracle-license-v1-1 seen true | debconf-set-selections  && \
    apt-get install -y --force-yes oracle-java$JAVA_VERSION-installer oracle-java$JAVA_VERSION-set-default  && \
    rm -rf /var/cache/oracle-j*-installer  && \
    apt-get clean  && \
    rm -rf /var/lib/apt/lists/*
ENV JAVA_HOME=/usr/lib/jvm/java-8-oracle 
############### 
## Sleuthkit ##
############### 
RUN \
    apt-get update  && \
    apt-get install -y --force-yes \
      build-essential automake autoconf libafflib-dev libtool ant libewf-dev git && \
    apt-get clean  && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /usr/local/src/
ENV SLEUTH_VER sleuthkit-4.4.0-patch02 
RUN git clone https://github.com/deepcase/sleuthkit.git
WORKDIR /usr/local/src/sleuthkit/
RUN git checkout $SLEUTH_VER
RUN ./bootstrap
RUN ./configure --prefix=/usr/
RUN make
RUN make install


VOLUME /data
WORKDIR /data

CMD /bin/bash
