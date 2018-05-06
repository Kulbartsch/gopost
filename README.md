**gopost** is a tool to make posts to several social networks at once and also generates a web html snippet to embed in your (static) website. 

This is work in progress, currently Mastodon toots and file snippets are working.

### Usage

    gopost [options] < text2post > websnippet

      -mastodon
        	Post message to mastodon
      -note
        	Output of message as microformat minimal note to stdout (default true)
      -test
        	testing, no external call to social networks
      -twitter
        	Post message to twitter
      -verbose
        	verbose output
 
### Precondition

For the use of mastodon the [madonctl](https://github.com/McKael/madonctl) has to be installed and [configured](https://lilotux.net/%7Emikael/pub/madonctl/).

### License

GPL

### Libraries 

The following non standard Libraries are used:

* [madonctl](https://github.com/McKael/madonctl) which uses 
** [madon](https://github.com/McKael/madon)

### Other

The initial release was developed on the 
and is mentioned [here](https://indieweb.org/projects#gopost).
