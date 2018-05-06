**gopost** is a tool to make posts to several social networks and also generate a webtemplate-snippet to embed in your (static) website. 

This is work in progress, currently Mastodon toots and file snipets are working.

### Usage

    gopost [options] < text2post > websnippet

      -mastodon
        	Post message to mastodon
      -note
        	Output of message as microformat minimal note to stdout (default true)
      -test
        	testing, no external call
      -twitter
        	Post message to twitter
      -verbose
        	verbose output
 
## Precondition

For the use of mastodon the [madonctl](https://github.com/McKael/madonctl) has to be installed and configured.

## License

GPL

