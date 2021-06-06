## [Websockets](https://stackoverflow.com/a/29933428)

With Websockets you have a stateful message passing system where messages can be sent either way and sending a message
has a lower overhead than with a RESTful HTTP request/response. Primary advantages of a connected websocket are:

1. **Two way communication.**
    - The server can notify the client of anything at any time.
    - Instead of polling a server on some regular interval to see if there is something new, a client can establish a
      webSocket and just listen for any messages coming from the server.
    - From the server's point of view, when an event of interest for a client occurs, the server simply sends a message
      to the client. The server cannot do this with plain HTTP.
2. **Lower overhead per message.**
    - If you anticipate a lot of traffic flowing between client and server, then there's a lower overhead per message
      with a webSocket.
    - This is because the TCP connection is already established and you just have to send a message on an already open
      socket.

      With an HTTP REST request, you have to first establish a TCP connection which is several back and forths between
      client and server. Then, you send HTTP request, receive the response and close the TCP connection. The HTTP
      request will necessarily include some overhead such as all cookies that are aligned with that server even if those
      are not relevant to the particular request.

      HTTP/2 (newest HTTP spec) allows for some additional efficiency in this regard if it is being used by both client
      and server because a single TCP connection can be used for more than just a single request/response.
3. **Stateful connections.**  
   Without resorting to cookies and session IDs, you can directly store state in your program for a given connection.
   While a lot of development has been done with stateless connections to solve most problems, sometimes it's just
   simpler with stateful connections.
