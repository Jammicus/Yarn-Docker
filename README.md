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

## Centos

From the root, run one of the following:

### Builder

`docker build -f Centos/BuilderDockerfile .`

### Firefox

`docker build -f Centos/Firefox .`

### Chrome

`docker build -f Centos/Chrome .`

## Ubuntu

From the root, run one of the following:

### Builder

`docker build -f Ubuntu/BuilderDockerfile .`

### Firefox

`docker build -f Ubuntu/Firefox .`

### Chromium

`docker build -f Ubuntu/Chromium .`
