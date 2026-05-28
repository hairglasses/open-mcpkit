package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/hairglasses/open-mcpkit/internal/openmcpkit"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return usage()
	}
	switch args[0] {
	case "manifest":
		return printJSON(openmcpkit.SampleGateway().Manifest())
	case "call":
		return call(args[1:])
	case "help", "-h", "--help":
		return usage()
	default:
		return fmt.Errorf("unknown command %q\n\n%w", args[0], usage())
	}
}

func call(args []string) error {
	if len(args) == 0 {
		return errors.New("usage: open-mcpkit call <tool> [--param key=value]")
	}
	tool := args[0]
	fs := flag.NewFlagSet("call", flag.ContinueOnError)
	params := multiFlag{}
	fs.Var(&params, "param", "key=value parameter")
	if err := fs.Parse(args[1:]); err != nil {
		return err
	}
	resp, err := openmcpkit.SampleGateway().Call(context.Background(), openmcpkit.Request{Tool: tool, Params: params.Map()})
	if err != nil {
		return err
	}
	return printJSON(resp)
}

func usage() error {
	return errors.New("usage: open-mcpkit manifest | call <tool> [--param key=value]")
}

type multiFlag []string

func (m *multiFlag) String() string { return strings.Join(*m, ",") }
func (m *multiFlag) Set(value string) error {
	if !strings.Contains(value, "=") {
		return fmt.Errorf("param %q must be key=value", value)
	}
	*m = append(*m, value)
	return nil
}
func (m multiFlag) Map() map[string]string {
	out := map[string]string{}
	for _, item := range m {
		key, value, _ := strings.Cut(item, "=")
		out[strings.TrimSpace(key)] = strings.TrimSpace(value)
	}
	return out
}

func printJSON(v any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}
