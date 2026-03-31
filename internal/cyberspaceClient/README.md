## Cyberspace API Client

This is the repository for @7spires Cyberspace Client. The current 'main' part of it is the internal/cyberspaceClient directory, which contains a set of functions for communicating with the Cyberspace API, as well as a core Client struct on which these functions are methods and which keeps state. 

While I am also working on my own interface, the idea was that the Client library would be independent of it and could be easily used by other people to build their own interfaces. 
Each section of the API documentation has a corresponding file in the client library, containing one function per API action. Most files contain the structs relevant to their own operation, unless they act on objects from another section (e.g. replies to posts). 
The remaining two files are main_client, which contains the core struct that most functions are methods to, and utilities, which contains internal utility functions such as ones relating to http requests. 

Core Client struct contents:
- http.Client
- API base url
- Username
- UserID
- Last status code received: If a request failed due to expired auth tokens, this lets you identify that and refresh the tokens before retrying the request.
- Cached posts/replies, etc.: Not essential, but can be useful for optimization and working around API usage limits.
- Cursors: Similar to the above. Lets you store e.g. up to where in the feed you have loaded.
