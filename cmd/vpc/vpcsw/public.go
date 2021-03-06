package vpcsw

import (
	"github.com/joyent/freebsd-vpc/cmd/vpc/vpcsw/create"
	"github.com/joyent/freebsd-vpc/cmd/vpc/vpcsw/destroy"
	"github.com/joyent/freebsd-vpc/cmd/vpc/vpcsw/list"
	"github.com/joyent/freebsd-vpc/cmd/vpc/vpcsw/port"
	"github.com/joyent/freebsd-vpc/internal/command"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const cmdName = "switch"

var Cmd = &command.Command{
	Name: cmdName,

	Cobra: &cobra.Command{
		Use:     cmdName,
		Aliases: []string{"sw"},
		Short:   "VPC switch management",
	},

	Setup: func(self *command.Command) error {
		subCommands := command.Commands{
			create.Cmd,
			destroy.Cmd,
			list.Cmd,
			port.Cmd,
		}

		if err := self.Register(subCommands); err != nil {
			return errors.Wrapf(err, "unable to register sub-commands under %s", cmdName)
		}

		return nil
	},
}
