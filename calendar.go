package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type CalendarCmd struct {
}

func (*CalendarCmd) Name() string {
	return "calendar"
}

func (*CalendarCmd) Synopsis() string {
	return "print JRA horse race schedule"
}

func (*CalendarCmd) Usage() string {
	return "TBD"
}

func (cmd *CalendarCmd) SetFlags(f *flag.FlagSet) {
}

func (cmd *CalendarCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	// TBD
	return subcommands.ExitSuccess
}
