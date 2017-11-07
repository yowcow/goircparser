package goircparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	type Case struct {
		input           string
		expectedPrefix  string
		expectedCommand string
		expectedParams  []string
		expectedSuffix  string
	}
	cases := []Case{
		Case{
			":test!test@test.test PRIVMSG #Test :This is a test case",
			"test!test@test.test",
			"PRIVMSG",
			[]string{"#Test"},
			"This is a test case",
		},
		Case{
			"PING :hogefuga",
			"",
			"PING",
			[]string{},
			"hogefuga",
		},
		Case{
			":irc.foobar.org 372 yowcow0 :- *             H    E    L    L    O              *",
			"irc.foobar.org",
			"372",
			[]string{"yowcow0"},
			"- *             H    E    L    L    O              *",
		},
		Case{
			":irc.foobar.org 005 yowcow0 CHANNELLEN=50 NICKLEN=31 TOPICLEN=490 AWAYLEN=127 KICKLEN=400 MODES=5 MAXLIST=beI:50 EXCEPTS=e INVEX=I PENALTY :are supported on this server",
			"irc.foobar.org",
			"005",
			[]string{
				"yowcow0",
				"CHANNELLEN=50",
				"NICKLEN=31",
				"TOPICLEN=490",
				"AWAYLEN=127",
				"KICKLEN=400",
				"MODES=5",
				"MAXLIST=beI:50",
				"EXCEPTS=e",
				"INVEX=I",
				"PENALTY",
			},
			"are supported on this server",
		},
	}

	for _, c := range cases {
		row := Parse(c.input)

		assert.Equal(t, c.expectedPrefix, row.Prefix)
		assert.Equal(t, c.expectedCommand, row.Command)
		assert.Equal(t, c.expectedSuffix, row.Suffix)

		for i, p := range c.expectedParams {
			assert.Equal(t, p, row.Params[i])
		}
	}
}
