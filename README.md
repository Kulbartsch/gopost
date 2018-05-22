**gopost** is a tool to make posts to several social networks at once and also generates a web html snippet to embed in your (static) website. 

The use cases are:

* Generate some HTML snippet (or markdown, or what ever) to embed it in a website (or for logging reasons), and
* To post to several social media services at once, either
   * interactive with direct user input, or
   * automatically when a new note is created on a (personal) website.

This is work in progress, currently Mastodon toots, Twitter tweets and web snippet export are working.

## Usage

    gopost [options] < text2post > websnippet

**gopost** takes the text to post from ```stdin``` and sends the websnippet to ```stdout```.
(Error) messages are send to ```stderr```. 

The options are: 

      -config
        	path to config file
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
 
## Configuration

The configuration file is expected in ```~/.config/gopost/config.json``` .
An alternative file can be chosen with the commandline parameter ``.config``.
It's a plain json file with the format:

    {
       "key": "value",
       ...
    }

The configuration for the different features are as follows:

### Note config

* ```noteHead``` - A text string which represents the first part of the web snippet which will send to standard out. 
	It must contain a ```%s``` which will be replaced by the current time. The default value is:  
        ```"<div class=\"h-entry\">\n  <time class=\"dt-published\">%s</time>: \n  <p class=\"p-content\">\n"```
        After the note header the message lines will be inserted.
 * ```noteFoot```- A text string representing the last part of the web snippet.
 	Default is: ```"  </p>\n</div>"```

### Twitter config

* ```twitter``` - by default no message is sent to Twitter. You can change the Default to ```true```

The following informations must be gained from Twitter. Here is how: https://dev.twitter.com/web/sign-in/implementing

* ```TwitterConsumerKey```
* ```TwitterConsumerSecret```
* ```TwitterAccessToken```
* ```TwitterAccessSecret```

### Mastodon config

No config inside **gopost** now, because the external tool *madonctl* is used. This it's configuration in the next chapter "Precondition". 
(Yes, it's a todo to integrate this feature into **gopost** and remove the precondition.)

## Precondition

For the use of Mastodon the [madonctl](https://github.com/McKael/madonctl) has to be installed and [configured](https://lilotux.net/%7Emikael/pub/madonctl/).

## License

GPL something

## Libraries 

The following non standard Libraries are used:

* [madonctl](https://github.com/McKael/madonctl) which uses 
   * [madon](https://github.com/McKael/madon)
* [go-twitter](https://github.com/dghubble/go-twitter)

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
