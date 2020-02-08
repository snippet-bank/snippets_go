First, run the server

## "40x" client errors

In HTTP, client error codes are in the range `400-499`.

I.e., the server believes that the client sent an invalid request.

Hit the server at the path `/client-error-path` with `curl` using the `-i` option. 

```sh
$ curl -i localhost/client-error-path
HTTP/1.1 400 Bad Request
Date: Sat, 08 Feb 2020 18:56:11 GMT
Content-Length: 10
Content-Type: text/plain; charset=utf-8

test error
```

^^^ Note the response code is `400 Bad Request`.

## "50x" client errors

In HTTP, server error codes are in the range `500-599`.

I.e, the server believes that it was at fault (not the client), for not being able to serve a valid response.

Try hitting the path `/server-error-path`:

```sh
$ curl -i localhost/server-error-path
HTTP/1.1 500 Internal Server Error
Date: Sat, 08 Feb 2020 18:56:37 GMT
Content-Length: 17
Content-Type: text/plain; charset=utf-8

test server error
```

^^^ Note the response code is `500 Internal Server Error`.
