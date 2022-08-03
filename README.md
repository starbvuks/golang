# GoLang Practice

## /go-server:

Basic web server built mainly using `net\http` package
Some notable functions used are:

- ResponseWriter: <br>
  > Mechanism used for sending responses to connected HTTP clients and how `res` headers are set
- Request:
  > Allows data to be retrieved from web `req` allowing access of POST*ed* data
- Error:
  > Replies to the request with the specified error message and HTTP code
- StatusNotFound:
  > Throw 404 error case
- ParseForm:
  > Parses form values to struct
- FormValue:
  > Self explanatory
- FileSever:
  > Returns handler to serve HTTP requests with contents of file system at root
- Handle:
  > Registers handler for given pattern in DefaultServeMux
- HandleFunc:
  > Registers handler func for given pattern in DefaultServeMux
- ListenAndServe:
  > Listens on the TCP network address addr and then calls Serve with handler requests on incoming connections
- Dir:
  > Implements FileSystem using native file system restricted to specific dir tree
- Fatal:
  > Fatal is equivalent to Print() followed by Exit call
