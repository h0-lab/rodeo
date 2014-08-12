package redis

import "strings"
import "fmt"
import "regexp"

// CommandZRem provides TCP communication of `ZREM`.
type CommandZRem struct {
	key string
	val string
	commandDefault
}

func (cmd CommandZRem) build() []byte {
	words := []string{
		"*3",
		cmd.strlen(cmdZREM),
		cmdZREM, // TODO: DRY
		cmd.strlen(cmd.key),
		cmd.key,
		cmd.strlen(cmd.val),
		cmd.val,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}

func (cmd CommandZRem) parse(res []byte) (result string, e error) {

	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match("\\$.+\\r\\n", res); ok {
		lines := strings.Split(string(res), "\r\n")
		// TODO: validate
		result = lines[1]
		return
	}
	e = fmt.Errorf("Response to `ZREM` is `%v`", string(res))
	return
}