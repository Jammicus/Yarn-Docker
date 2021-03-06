FROM ubuntu:18.04

ENV NODE_VERSION 12.18.4
ENV NODEJS_HOME /opt/nodejs/node-v${NODE_VERSION}/bin
ENV PATH /opt/nodejs/node-v${NODE_VERSION}/bin:$PATH

ENV YARN_VERSION 1.22.5
ENV PATH /opt/yarn/yarn-v${YARN_VERSION}/bin:`yarn global bin`:$PATH

RUN apt-get update
RUN apt install -y curl \
                   build-essential \
                   dirmngr \
                   gpg 

RUN apt-get remove -y nodejs                    
RUN mkdir /opt/nodejs

WORKDIR /opt/nodejs

# gpg keys listed at https://github.com/nodejs/node#release-team
RUN for key in \
    4ED778F539E3634C779C87C6D7062848A1AB005C \
    94AE36675C464D64BAFA68DD7434390BDBE9B9C5\
    B9AE9905FFD7803F25714661B63B535A4C206CA9 \
    77984A986EBC2AA786BC0F66B01FBB92821C587A \
    71DCFD284A79C3B38668286BC97EC7A07EDE3FC1 \
    FD3A5288F042B6850C66B31F09FE44734EB7990E \
    8FCCA13FEF1D0C2E91008E09770F7A9A5AE15600 \
    C4F0DFFF4E8C1A8236409D08E73BC641CC11F4C8 \
    DD8F2338BAE7501E3DD5AC78C273792F7D83545D \
    A48C2BEE680E841632CD4E44F07496B3EB3C1762 \
    B9E2F5981AA6E0CD28160D9FF13993A75599653C \
    9554F04D7259F04124DE6B476D5A82AC7E37093B \
    93C7E9E91B49E432C2F75674B0A78B0A6C481CF6 \
    56730D5401028683275BD23C23EFEFE93C4CFFFE \
    114F43EE0176B71C7BC219DD50A3051F888C628D \
    7937DFD2AB06298B2293C3187D33FF9D0246406D \
  ; do \
    gpg --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys "$key" || \
    gpg --keyserver hkp://ipv4.pool.sks-keyservers.net --recv-keys "$key" || \
    gpg --keyserver hkp://pgp.mit.edu:80 --recv-keys "$key" ; \
  done \
    && curl -SLO "https://nodejs.org/dist/v$NODE_VERSION/node-v$NODE_VERSION-linux-x64.tar.gz" \
    && curl -SLO --compressed "https://nodejs.org/dist/v$NODE_VERSION/SHASUMS256.txt.asc" \
    && gpg --batch --decrypt --output SHASUMS256.txt SHASUMS256.txt.asc \
    && grep " node-v$NODE_VERSION-linux-x64.tar.gz\$" SHASUMS256.txt | sha256sum -c - && \
    tar zvxf node-v${NODE_VERSION}-linux-x64.tar.gz -C /opt/nodejs && \
    mv /opt/nodejs/node-v${NODE_VERSION}-linux-x64 /opt/nodejs/node-v${NODE_VERSION} && \
    rm -rf /opt/nodejs/node-v$NODE_VERSION-linux-x64.tar.gz /opt/nodejs/SHASUMS256.txt /opt/nodejs/SHASUMS256.txt.asc
 
RUN mkdir /opt/yarn
WORKDIR /opt/yarn

RUN curl -OL https://github.com/yarnpkg/yarn/releases/download/v${YARN_VERSION}/yarn-v${YARN_VERSION}.tar.gz
 
RUN curl https://dl.yarnpkg.com/debian/pubkey.gpg | gpg --import && \
    curl -OL https://github.com/yarnpkg/yarn/releases/download/v${YARN_VERSION}/yarn-v${YARN_VERSION}.tar.gz.asc && \
    gpg --verify yarn-v${YARN_VERSION}.tar.gz.asc && \ 
    tar zvxf yarn-v${YARN_VERSION}.tar.gz -C /opt/yarn && \
    rm -rf /opt/yarn/yarn-v${YARN_VERSION}.tar.gz /opt/yarn/yarn-v${YARN_VERSION}.tar.gz.asc

WORKDIR /home

CMD ["/bin/bash"]