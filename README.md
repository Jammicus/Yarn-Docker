# Yarn-Docker
Docker files for testing applications that use  Node and Yarn on different operating systems


## Alpine

From the root, run one of the following:

### Builder

`docker build -f Alpine/BuilderDockerfile .`

### Firefox

`docker build -f Alpine/Firefox .`

### Chromium

`docker build -f Alpine/Chromium .`

## Debian

From the root, run one of the following:

### Builder

`docker build -f Debian/BuilderDockerfile .`

### Firefox

`docker build -f Debian/Firefox .`

### Chromium

`docker build -f Debian/Chromium .`

## Centos

From the root, run one of the following:

### Builder

`docker build -f Centos/BuilderDockerfile .`

### Firefox

`docker build -f Centos/Firefox .`

### Chromium

`docker build -f Centos/Chromium .`

## Ubuntu

From the root, run one of the following:

### Builder

`docker build -f Ubuntu/BuilderDockerfile .`

### Firefox

`docker build -f Ubuntu/Firefox .`

### Chromium

`docker build -f Ubuntu/Chromium .`
