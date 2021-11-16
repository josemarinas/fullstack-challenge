# Server

The server exposes an Rest api in the spcified port and location. The APi has 3 main enpoints

* `POST /image`: allows to create an image in the database

* `GET /image/:id`: gets the image with the specified id
* `GET /ping`: returns a message to check if the service is working

The main code can be found in `cmd/root.go`. First it tries to read the configuration file, after tries to connect to the database and if everithing has gone well it will expose the REST api in the port specified in the config file.

The errors returned by the api can be found in tthe folder `httperrors` and the CORS middleware allows request from all domains.

## Dependencies

* `rs/zerologg`
* `gin`
* `sql`
* `satori/go.uuid`
* `spf13/cobra`
* `spf13/viper`


## Run

`go mod init server && go mod tidy`

`go run main.go`
or 
`go build main.go && ./main`