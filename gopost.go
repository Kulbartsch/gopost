// gopost - Post messages to different social networks
// (c) 2018 Alexander Kulbartsch
// License GPL

// TODO: static URL (as an parameter)
// TODO: placeholder in message and note templates for variables like static URL and time
// TODO: add a frontend

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"os"
	"os/exec"
	"os/user"
	"time"
	"path/filepath"
)

var verbose bool
var test bool

type Configuration struct {
	noteHead		string
	noteFoot		string
	twitter			bool
        TwitterConsumerKey	string
        TwitterConsumerSecret	string
        TwitterAccessToken	string
        TwitterAccessSecret	string
	Variable1		string
	Variable2		string
	Variable3		string
	Variable4		string
}

func note(message []string) {
	// TODO: use a template
	t := time.Now()
	fmt.Printf("<div class=\"h-entry\">\n  <time class=\"dt-published\">%s</time>: \n  <p class=\"p-content\">\n", t.String())
	for _, l := range message {
		fmt.Printf("    %s\n", l)
	}
	fmt.Println("  </p>\n</div>")
}

func mastodon(messagestring string) {
	// TODO: native implementation / check https://github.com/mattn/go-mastodon or https://github.com/McKael/madon)

	cmd := "madonctl toot \"" + messagestring + "\""
	if verbose {
		fmt.Fprintf(os.Stderr, "mastodon: The madonctl command is: %s\n", cmd)
	}
	if !test {
		out, err := exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			fmt.Fprintln(os.Stderr, "mastodon: reading standard input: ", err)
		}
		if verbose {
			fmt.Fprintf(os.Stderr, "mastodon: The madonctl out is: %s\n", out)
		}
	}
}

func tweet(messagestring string) {
	// refers to https://github.com/dghubble/go-twitter
	// TODO: optional remove "@twitter.com" from mentions like "@username@twitter.com"

	config := oauth1.NewConfig("consumerKey", "consumerSecret")
	token := oauth1.NewToken("accessToken", "accessSecret")
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Home Timeline
	// tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
	//     Count: 20,
	// })

	// Send a Tweet
	if !test {
		tweet, resp, err := client.Statuses.Update(messagestring, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Twitter: error posting: ", err)
		}
		if verbose {
			fmt.Fprintln(os.Stderr, "Twitter: tweet: ", tweet)
			fmt.Fprintln(os.Stderr, "Twitter: response: ", resp)
		}
	}

}

func main() {

	// parse vars
	var out_note = flag.Bool("note", true, "Output of message as microformat minimal note to stdout")
	var out_mastodon = flag.Bool("mastodon", false, "Post message to mastodon")
	var out_twitter = flag.Bool("twitter", false, "Post message to twitter")
	var configfile = flag.String("config", "~/.config/gopost/config.json", "configuration file name")
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.BoolVar(&test, "test", false, "testing, no external call to social networks")

	flag.Parse()

	// read config file
	var fqconfigfile string
	fqconfigfile = *configfile
	if fqconfigfile[:2] == "~/" {
		usr, _ := user.Current()
		fqconfigfile = filepath.Join(usr.HomeDir, fqconfigfile[2:])
	}
	file, err := os.Open(fqconfigfile)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "main: error reading configuration file: ", err)
	} else {
		decoder := json.NewDecoder(file)
		configuration := Configuration{}
		err = decoder.Decode(&configuration)
		if err != nil {
			fmt.Fprintln(os.Stderr, "main: error decoding configuration file: ", err)
		}
	}
	if verbose {
		fmt.Fprintln(os.Stderr, "main: config from file: ", configuration)
	}

	// read message
	var message []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message = append(message, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "main: error reading standard input: ", err)
	}
	var messagestring string = ""
	for i, l := range message {
		if i == 0 {
			messagestring = l
		} else {
			messagestring = messagestring + "\n" + l
		}
	}

	// send messages

	if *out_note {
		go note(message)
	}

	if *out_mastodon {
		go mastodon(messagestring)
	}

	if *out_twitter {
		go tweet(messagestring)
	}

	// TODO: add Diaspora*

}
