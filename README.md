**gopost** is a tool to make posts to several social networks at once and also generates a web html snippet to embed in your (static) website.

The use cases are:

* Generate some HTML snippet (or markdown, or what ever) to embed it in a website (or for logging reasons), and
* To post to several social media services at once, either
   * interactive with direct user input, or
   * automatically when a new note is created on a (personal) website.

This is work in progress, currently Mastodon toots and file snippets are working.

### Usage

    gopost [options] < text2post > websnippet

    **gopost** takes the text to post from ```stdin``` and sends the websnippet to ```stdout```.
    (Error) messages are send to ```stderr```.

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

## Precondition

For the use of mastodon the [madonctl](https://github.com/McKael/madonctl) has to be installed and [configured](https://lilotux.net/%7Emikael/pub/madonctl/).

## License

GPL something

## Libraries

The following non standard Libraries are used:

* [madonctl](https://github.com/McKael/madonctl) which uses
** [madon](https://github.com/McKael/madon)

## Other

### Primary Goal

On the first hand the idea was to have tool to post simple text messages to different social networks at once and also offer an export this message in a way to respect the principles of the IndieWeb.
I also like the idea to have an universal tool without dependencies, that can be used as a frontend to post things, as well as having this integrated into a build process for (static) websites.

### Further improvements

A list of improvements, going further than the *todos* in the source code.

* Sending Pictures
* not just sending simple messages (aka notes in IndieWeb terminology), but more complex messages respectively blog-posts. An abstract could be used to send as short notes to twitter and alike services.

### Meta

The initial release was developed on the [IndieWebCamp 2018](https://indieweb.org/2018/D%C3%BCsseldorf)
and is mentioned [here](https://indieweb.org/projects#gopost).
