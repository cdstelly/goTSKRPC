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


ADD goTSK.go /usr/local/src/goTSK.go
RUN apt update
RUN apt install -y golang-go
WORKDIR /usr/bin/
RUN go build /usr/local/src/goTSK.go

# ADD goTSK /usr/bin/goTSK

# CMD /bin/bash
VOLUME /data
WORKDIR /data
CMD /usr/bin/goTSK
