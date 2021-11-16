# Client

The project was generated usring the `create-react-app` with the typescript flag. It has a really simple router ant the src has 2 main folders `components` and `pages`
The components folder contains 2 components that are used in the `Home` page:

## Image Uploader

The image uploader uploads an image to IPFS and send an POST request to the backend to store the generated CID, if everything goes fine the interface will show the UUID associated to the CID.

## Image Viewer

The UUID can be used to retrieve the image in the imagge viewer. The image viewer will retrieve the image CID from the backend by doing a GET request. If the request goes fine it will generate a link to view the image in the IPFS gateway and it will also show the image by decoding it into a base64 string.

## Configuration

The base urls for both the ipfs and rest api con be configured with `.env` files or variables:

* `REACT_APP_API_URL`
* `REACT_APP_IPFS_URL`

The default values are `http://localhost:8000/api` and `/ip4/0.0.0.0/tcp/5001`

## Dependencies

* `ipfs-http-client`
* `uint8arrays`
* `react-query`
* `@mui/material`

## Run

`yarn start`