// gopost - Post messages to different social networks
// (c) 2018 Alexander Kulbartsch
// License GPL

// TODO: static URL (as an parameter)
// TODO: placeholder in message and note templates for variables like static URL and time
// TODO: add a frontend

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var verbose bool
var test bool

var socialLogins = map[string]string{
	"twitter_id": "bla",
	"twitter_":   "blub",
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

func mastodon(message []string) {
	// TODO: native implementation

	cmd := "madonctl toot \""
	for i, l := range message {
		if i > 0 {
			cmd = cmd + "\n" + l
		} else {
			cmd = cmd + l
		}
	}
	cmd = cmd + "\""
	if verbose {
		fmt.Fprintf(os.Stderr, "The madonctl command is: %s\n", cmd)
	}
	if !test {
		out, err := exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		if verbose {
			fmt.Fprintf(os.Stderr, "The madonctl out is: %s\n", out)
		}
	}
}

func tweet(message []string) {
	fmt.Fprintln(os.Stderr, "Twitter not yet implemented")
	return

	// TODO: implement ...

}

func main() {

	// parse vars
	var out_note = flag.Bool("note", true, "Output of message as microformat minimal note to stdout")
	var out_mastodon = flag.Bool("mastodon", false, "Post message to mastodon")
	var out_twitter = flag.Bool("twitter", false, "Post message to twitter")
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.BoolVar(&test, "test", false, "testing, no external call")

	flag.Parse()

	// read message
	var message []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message = append(message, scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// send messages

	if *out_note {
		note(message)
	}

	if *out_mastodon {
		mastodon(message)
	}

	if *out_twitter {
		tweet(message)
	}

	// TODO: add Diaspora*

	// TODO: add Matrix

}
