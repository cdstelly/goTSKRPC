FROM debian

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
