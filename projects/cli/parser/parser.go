package parser

import (
	"bytes"
	"strconv"

	"github.com/go-ptr/go-ptr/ptr"
	"github.com/luma-dev/fishis/projects/cli/gen/models"
	iptr "github.com/luma-dev/fishis/projects/cli/internal/ptr"
)

func ParseHistories(in []byte) (h []models.History) {
	p := &parser{
		in: in,
		i:  0,
	}
	return p.parseHistories()
}

type parser struct {
	in []byte
	i  int
}

func (p *parser) peekable(a int) bool {
	return p.i+a < len(p.in)
}

func (p *parser) peek(a int) byte {
	return p.in[p.i+a]
}

func (p *parser) read() (c byte) {
	c = p.in[p.i]
	p.i = p.i + 1
	return
}

func (p *parser) end() bool {
	return p.i >= len(p.in)
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\n' || c == '\r' || c == '\t'
}

func isPunc(c byte) bool {
	return c == ' ' || c == '\n' || c == '\r' || c == '\t' || c == '-' || c == ':'
}

func (p *parser) skipSpaces() {
	for !p.end() && isSpace(p.peek(0)) {
		p.read()
	}
}

func (p *parser) skipPuncs() {
	for !p.end() && isPunc(p.peek(0)) {
		p.read()
	}
}

func (p *parser) readTillPunc() []byte {
	from := p.i
	for !p.end() && !isPunc(p.peek(0)) {
		p.read()
	}
	return p.in[from:p.i]
}

func (p *parser) readString() (str []byte) {
	for !p.end() && p.peek(0) != '\n' {
		if p.peek(0) == '\\' && p.peek(1) == 'n' {
			p.read()
			p.read()
			str = append(str, '\n')
		} else {
			str = append(str, p.read())
		}
	}
	return
}

func (p *parser) parseHistories() (h []models.History) {
	for !p.end() {
		e := p.parseHistory()
		if e.Cmd != nil && *e.Cmd != "" {
			h = append(h, e)
		}
	}
	return
}

func (p *parser) parseHistory() (h models.History) {
	p.skipPuncs()
	for !p.end() {
		p.skipSpaces()
		if p.peek(0) == '-' {
			break
		}
		key := p.readTillPunc()
		p.read()
		if bytes.Equal(key, []byte("cmd")) {
			p.skipSpaces()
			h.Cmd = ptr.NewString(string(p.readString()))
			p.read()
		} else if bytes.Equal(key, []byte("when")) {
			p.skipSpaces()
			i, _ := strconv.Atoi(string(p.readString()))
			h.When = iptr.NewTimestamp(models.Timestamp(i))
			p.read()
		} else if bytes.Equal(key, []byte("paths")) {
			h.Paths = []string{}
			p.read()
			for !p.end() && p.peek(0) == ' ' {
				p.skipPuncs()
				h.Paths = append(h.Paths, string(p.readString()))
				p.read()
			}
		}
	}
	return
}
