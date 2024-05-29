# TerminalPad

A simple golang application, where the users can write notes in the terminal

## Architecture

For this to work, I want it to be a client-server architecture.
The server would be responsible to store and manage all the notes.
The client would take in the notes that the user has, show them to the user,
and be responsible for the communication with the server.

- The communication for the server should be in gRPC
- I want to manage the notes for each user. I want to be able to perform a
connection to the server using user/pass, and then access all the notes
- The client is to be done in a beautiful console application

*Login*
Would be nice for the connection to have a "cookie" which would explain if
the user is authenticated

### gRPC calls

1. Get message content
2. Get all messages titles and ids
3. Post/Update note
4. Login and get a cookie/token

### Notes screen
_____________________
|                   |
|note 1             |
|note 2             |
|note 3             |
|note 4             |
|               new |
|___________________|

### Note selected
_____________________
|<-   Note X        |
|                   |
|content            |
|                   |
|                   |
|___________________|


### Login
_____________________
|                   |
|                   |
|       User:       |
|       Password:   |
|       Login       |
|___________________|

## TODO - Server

1. Restructure logging. Is not being applied properly.
Values passed should have values such as metrics (key-pair attributes)
Pass down the context through the functions so that the same logger is
always used
2 Restructure error handling. Does not seem right
