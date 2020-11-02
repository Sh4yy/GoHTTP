# GoHTTP Client  
  > A basic CURL clone written in Golang.
  
#### Basic Use Case  
`make build`\
`./client --url http://google.com --method POST --header "Content-Type: application/json" --body "{\"ok\": true}"`

#### Supported Flags
`url` target url, required.\
`method` http request method, i.e. `GET`, `POST`, `PATCH`, `DELETE`. Defaults to `GET`.\
`header` http request header in `"Key: Value"` format, can be repeated for multiple headers.\
`body` http request body.


###### Note: This is very shitty code, please don't use in production. (as if anyone would lol)
