# Fullstack Dev Challenge

## Aditional notes

Both client and server are provided with a Dockerfile to build iomages to both services. The docker compose was modified to  assign static ips to the services. In addition a `docker-compose.prod.yml` file was added with the intention of building and deploying both new services.

A makefile was added:

* `make up`: will deploy the base docker compose
* `make up prod`: will build and deploy the prod docker compose

### Possible imporvments:

* Add enpoint to search all the files uploaded
* Add delete capabilities
* Add automated testing


## Description

The objective is to develop a web application (no UX design necessary) where the following flows are supported:

### Flow 1 (Upload)
* User acceses the web and can upload an image to IPFS
* the IPFS hash produced is send to a go backend via a rest API
* the backend stores the IPFS hash associated to a simple SQL datababase
* the backend returns a UID of the image

### Flow 2 (Retrieve)
* User acceseses the web and submits a UID provided by a previous image upload
* The UI retrieves the IPFS hash (if existing) through the REST API of the backend
* The UI retrieves the IPFS images using the hash and demonstrates the image

To facilitate development a docker-compose is provided including a local IPFS node (should be enough) and a Postgres db but feel free not to use it :-)

## Requirements
* Frontend : ReactJS + Typescript
* Backend : Golang + PostgreSQL
