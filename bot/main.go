package bot

import (
	"fmt"
	"os"
	"os/user"
	"regexp"

	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

const dir = "/bot_logs"
const fileName = "log.txt"

const (
	Comment   string = "t1_"
	Account   string = "t2_"
	Link      string = "t3_"
	Message   string = "t4_"
	Subreddit string = "t5_"
	Award     string = "t6_"
)

func main() {

	getCreateDir(dir)
	getComments()
}

type limbBot struct {
	bot reddit.Bot
}

func (b *limbBot) Comment(comment *reddit.Comment) error {
	fmt.Printf(`%s posted "%s"\n`, comment.Author, comment.Body)
	// b.limbretrievalbot.Reply()
	path := getCreateDir(dir)
	appendToLog(fmt.Sprintln(comment.Author, ":", comment.Body), path+"/"+fileName)
	return nil
}

func getComments() {
	bot, err := reddit.NewBotFromAgentFile("limbretrievalbot.agent", 0)
	if err != nil {
		fmt.Printf("Failed to init agent: %s", err)
	}
	cfg := graw.Config{
		SubredditComments: []string{"LimbRetrievalBotTest"},
		// SubredditComments: []string{"all"},
	}
	// limbretrievalbot.Reply("t1_homaek7", "Some test comment")
	handler := &limbBot{bot: bot}
	if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
		fmt.Println("Failed to start graw run: ", err)
	} else {
		fmt.Println("graw run failed: ", wait())
	}

	// harvest, err := limbretrievalbot.Listing("/r/all", "")
	// if err != nil {
	// 	fmt.Printf("Failed to load subreddit: %s", err)
	// }

	// for index, comment := range harvest.Comments[:100] {
	// 	fmt.Printf("Comment (#%d, %s", index, comment.Body)
	// }
}

func getCreateDir(dirName string) string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	path := usr.HomeDir + dirName
	dirErr := os.MkdirAll(path, os.ModePerm)
	if dirErr != nil {
		fmt.Printf("Failed to create dir: %s", err)
		os.Exit(1)
	}

	return path
}

func appendToLog(text string, filepath string) {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Printf("Failed to create file: %s", err)
		os.Exit(1)
	}

	f.WriteString(text + "\n")
	f.Close()
}

func checkContainsShrug(body string, bodyHtml string) Shrug {
	for _, s := range invalidShrugBodies {
		match, _ := regexp.Match(string(s), []byte(body))
		fmt.Println(match)

		if match {
			return s
		}
	}

	return NoShrug
}
