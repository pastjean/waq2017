package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//  https://api.tronalddump.io/random/quote
var (
	version = "v2"
	quotes  = trumpQuotes
	r       = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func getRandomQuote() string {
	return quotes[r.Intn(len(quotes))]
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("thequoter-version", version)
		w.Header().Set("content-type", "text/plain; charset=utf-8")
		w.Write([]byte(fmt.Sprintf("You want a nice quote, here is one:\n%s", getRandomQuote())))
		//w.Header().Set("content-type", "text/html; charset=utf-8")
		//tmpl, _ := template.New("t").Parse(htmlTemplate)
		//tmpl.Execute(w, map[string]string{"Quote": getRandomQuote()})
	})

	port := "0.0.0.0:8080"
	log.Println("Listening on: ", port)
	log.Panic(http.ListenAndServe(port, nil))
}

var (
	norrisQuotes = []string{
		"Chuck Norris does not hunt because the word hunting infers the probability of failure. Chuck Norris goes killing",
		"Paper beats rock, rock beats scissors, and scissors beats paper, but Chuck Norris beats all 3 at the same time.",
		"The Chuck Norris Eclipse plugin made alien contact.",
		"A guy was furious at Chuck Norris. He fought with Chuck Norris with all of his rage. Chuck Norris happily fed on his rage until the guy became a hippie. Then Chuck Norris obliterated him, because Chuck Norris doesn't like hippies.",
		"Chuck Norris can sleep standing up...while laying down......awake.",
		"If someone says, nobodys perfect, Chuck Norris takes that as a personal insult.",
		"Chuck Norris taught lions and tigers how to chase down, kill and eat their prey. However, despite his best efforts on proper carnivorous table fare etiquette, they continue to start lavishly eating the assholes and genitals first.",
	}

	trumpQuotes = []string{
		"Hillary was involved in the e-mail scandal because she is the only one with judgement so bad that such a thing could have happened!",
		"Why did Clinton supporter Alison L. Grimes declare Crooked Hillary winner in KY when AP hasn't even called the race?",
		"I am somewhat surprised that Bernie Sanders was not true to himself and his supporters. They are not happy that he is selling out!",
		"Crooked Hillary Clinton is soft on crime, supports open borders, and wants massive tax hikes. A formula for disaster!",
		"JEB is a hypocrite! Used massive private \"Eminent Domain\" --- Just another clueless politician! \nhttps://t.co/yZT3lSutt9",
		"Hypocrite! Hillary Clinton claims she needs a \"public and a private stance\" in discussions with Wall Street banks.",
		"Crooked Hillary just can't close the deal with Bernie. It will be the same way with ISIS, and China on trade, and Mexico at the border. Bad!",
		"\"jeff shick: Donald J. Trump did you hear the mayor nutter of philadelphias comments towards you? Ignorance\" Yes, he is a crude dope!",
		"Crooked Hillary wants to get rid of all guns and yet she is surrounded by bodyguards who are fully armed. No more guns to protect Hillary!",
	}

	htmlTemplate = `<html>
		<head><title>waqquoter</title>
		<link href="https://fonts.googleapis.com/css?family=Palanquin+Dark" rel="stylesheet">
		</head>
		<body style="font-family: 'Palanquin Dark', sans-serif;font-size:2em;display:flex;align-items: center;flex-direction:column;justify-content: center;background-color:#efefef"><h1>The WAQquoter</h1><div style="display:flex;align-items: center;justify-content: center;background-color:white;min-height:50%;min-width:50%;max-width:300px;padding:2em 2em">{{.Quote}}</div></body>
	</html>`
)
