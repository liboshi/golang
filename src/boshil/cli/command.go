package cli

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"text/tabwriter"
)

type HasFlags interface {
	Register(ctx context.Context, f *flag.FlagSet)
	Process(ctx context.Context) error
}

type Command interface {
	HasFlags
	Run(ctx context.Context, f *flag.FlagSet) error
}

func generalHelp() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

	cmds := []string{}
	for name := range commands {
		cmds = append(cmds, name)
	}

	sort.Strings(cmds)

	for _, name := range cmds {
		fmt.Fprintf(os.Stderr, " %s\n", name)
	}
}

func commandHelp(name string, cmd Command, f *flag.FlagSet) {
	type HasUsage interface {
		Usage() string
	}

	fmt.Fprintf(os.Stderr, "Usage: %s %s [OPTIONS]", os.Args[0], name)
	if u, ok := cmd.(HasUsage); ok {
		fmt.Fprintf(os.Stderr, " %s", u.Usage())
	}
	fmt.Fprintf(os.Stderr, "\n")

	type HasDescription interface {
		Description() string
	}

	if u, ok := cmd.(HasDescription); ok {
		fmt.Fprintf(os.Stderr, "%s\n", u.Description())
	}

	n := 0
	f.VisitAll(func(_ *flag.Flag) {
		n += 1
	})

	if n > 0 {
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		tw := tabwriter.NewWriter(os.Stderr, 2, 0, 2, ' ', 0)
		f.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(tw, "\t-%s=%s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		tw.Flush()
	}
}

func clientLogout(ctx context.Context, cmd Command) error {
	type logout interface {
		Logout(context.Context) error
	}

	if l, ok := cmd.(logout); ok {
		return l.Logout(ctx)
	}

	return nil
}

func Say(words string) {
	fmt.Println(words)
}

func Run(args []string) int {
	var err error

	if len(args) == 0 {
		generalHelp()
		return 1
	}

	// Look up real command name in aliases table.
	name, ok := aliases[args[0]]
	if !ok {
		name = args[0]
	}

	cmd, ok := commands[name]
	if !ok {
		generalHelp()
		return 1
	}

	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)

	ctx := context.Background()
	cmd.Register(ctx, fs)

	if err = fs.Parse(args[1:]); err != nil {
		goto error
	}

	if err = cmd.Process(ctx); err != nil {
		goto error
	}

	if err = cmd.Run(ctx, fs); err != nil {
		goto error
	}

	if err = clientLogout(ctx, cmd); err != nil {
		goto error
	}

	return 0

error:
	if err == flag.ErrHelp {
		commandHelp(args[0], cmd, fs)
	} else {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
	}

	_ = clientLogout(ctx, cmd)

	return 1
}
