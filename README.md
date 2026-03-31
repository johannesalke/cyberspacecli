### Cyberspace CLI Client Prototype

At present, this client only has 3 real functions: Browse the feed, look at the replies to individual posts & write posts of your own.

To download it, you need to have Go installed. So long as that is given, you can simply clone the Git repo onto your machine (or download it via github), then while inside the project directory, execute the following commands: 

```go
go build -o cyberspace-client .
./cyberspace-client
```

Btw, it you're a programmer, please don't look at the contents of this too closely, especially main.go. You don't want to see what's going on in there. 

The client has the following commands:

- feed: Load 5 posts from the cyberspace feed. Each use loads the next 5 after the last one.
- write: Opens your default text editor (or if you have non, vim) and lets you write a post. Be aware that it might fail to post, so don't invest too much effort into it without copying the contents elsewhere before saving and closing the editor. After closing the editor, you'll have a chance to choose topics for the post.
- replies <post_id>: This command requires an argument, the post id shown in the top-line of each post.
- note <note_id>: Lets you open an existing note from your journal, edit it, and save the edited version.