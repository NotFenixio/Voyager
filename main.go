package main

import "github.com/gocolly/colly"
import "fmt"
import "strings"

var categoryIds = map[string]string{
	"Suggestions":                                                "1",
	"Bugs and Glitches":                                          "3",
	"Questions about Scratch":                                    "4",
	"Announcements":                                              "5",
	"New Scratchers":                                             "6",
	"Help with Scripts":                                          "7",
	"Show and Tell":                                              "8",
	"Project Ideas":                                              "9",
	"Collaboration":                                              "10",
	"Requests":                                                   "11",
	"Moderators":                                                 "12",
	"Deutsch":                                                    "13",
	"Español":                                                    "14",
	"Français":                                                   "15",
	"中文":                                                        "16",
	"Polski":                                                     "17",
	"日本語":                                                      "18",
	"Nederlands":                                                 "19",
	"Português":                                                  "20",
	"Italiano":                                                   "21",
	"עברית":                                                     "22",
	"한국어":                                                      "23",
	"Norsk":                                                      "24",
	"Türkçe":                                                     "25",
	"Ελληνικά":                                                   "26",
	"Pусский":                                                    "27",
	"Translating Scratch":                                        "28",
	"Things I'm Making and Creating":                             "29",
	"Things I'm Reading and Playing":                             "30",
	"Advanced Topics":                                            "31",
	"Connecting to the Physical World":                           "32",
	"Català":                                                     "33",
	"Other Languages":                                            "34",
	"Mentors Forum":                                              "35",
	"Bahasa Indonesia":                                           "36",
	"Scratch Day 2014":                                           "37",
	"Spam Dustbin":                                               "38",
	"Scratch Helper Groups":                                      "39",
	"Camp Counselor Forum":                                       "40",
	"Extension Developer's Forum":                                "41",
	"Scratch Stability Team Forum":                               "42",
	"Scratch Day 2015":                                           "44",
	"Scratch Design Studio Forum":                                "46",
	"Developing Scratch Extensions":                              "48",
	"Open Source Projects":                                       "49",
	"Welcoming Committee":                                        "50",
	"Community Blocks Forum":                                     "51",
	"Scratch Day 2016":                                           "52",
	"Scratch Day 2017":                                           "54",
	"Africa":                                                     "55",
	"Scratch Day 2018":                                           "56",
	"Scratch 3.0 Beta":                                           "57",
	"Camp Counselors 2020":                                       "58",
	"فارسی":                                                      "59",
	"Project Save & Level Codes":                                 "60",
	"April Fools Day - Suggest-Show-Question-Bugs-Help-Glitch-Tell-Etc": "61",
}

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("h3", func(e *colly.HTMLElement) {
		name := strings.TrimSpace(e.DOM.Text())
		if name == "Sign in" {
			return
		}
		fmt.Println("Post name:", name)
		parts := strings.Split(e.ChildAttr("a", "href"), "/")
		fmt.Println("Post ID:", parts[3])
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://corsproxy.josueart40.workers.dev/?https://scratch.mit.edu/discuss/1")
}
