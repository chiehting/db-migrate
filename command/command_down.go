package command

import (
	"flag"
	"fmt"
	"strings"

	migrate "github.com/heartz2o2o/db-migrate/migrate"
)

type DownCommand struct {
}

func (c *DownCommand) Help() string {
	helpText := `
Usage: sql-migrate down [options] ...

  Undo a database migration.

Options:

  -config=dbconfig.yml   Configuration file to use.
  -env="development"     Environment.
  -limit=1               Limit the number of migrations (0 = unlimited).
  -dryrun                Don't apply migrations, just print them.

`
	return strings.TrimSpace(helpText)
}

func (c *DownCommand) Synopsis() string {
	return "Undo a database migration"
}

func (c *DownCommand) Run(args []string) int {
	if err := c.RunProcess(args); err != nil {
		fmt.Print(err.Error())
		return 1
	}

	return 0
}

func (c *DownCommand) RunProcess(args []string) (err error) {
	var limit int
	var dryrun bool

	cmdFlags := flag.NewFlagSet("down", flag.ContinueOnError)
	cmdFlags.Usage = func() { fmt.Print(c.Help()) }
	cmdFlags.IntVar(&limit, "limit", 1, "Max number of migrations to apply.")
	cmdFlags.BoolVar(&dryrun, "dryrun", false, "Don't apply migrations, just print them.")
	ConfigFlags(cmdFlags)

	if err = cmdFlags.Parse(args); err != nil {
		return
	}

	err = ApplyMigrations(migrate.Down, dryrun, limit)
	if err != nil {
		return
	}

	return
}
