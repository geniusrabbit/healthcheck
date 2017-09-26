package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

var (
	errInvalidResponseStatus = errors.New("Invalid response status")
)

func main() {
	app := cli.NewApp()
	app.Action = checkAction
	app.Run(os.Args)
}

func checkAction(c *cli.Context) error {
	resp, err := http.Get(c.Args()[0])
	if err != nil {
		log.Println("> [Error]", err)
		return err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	if len(data) > 0 {
		log.Println("> [Response]", string(data))
	}
	resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		log.Printf("> [Error] Response status is %d %s\n", resp.StatusCode, resp.Status)
		return errInvalidResponseStatus
	}

	return nil
}
