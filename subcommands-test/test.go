package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
)

type config struct {
	cap string
}

func (*config) Name() string     { return "config" }
func (*config) Synopsis() string { return "Print args to stdout." }
func (*config) Usage() string {
	return `config [-test] <some text>:
	  Print args to stdout.
	  `
}

func (c *config) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.cap, "cap", "hello", "cap example")
}

func (c *config) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if c.cap != "" {
		fmt.Println(c.cap)
	}
	return subcommands.ExitSuccess
}

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&config{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
