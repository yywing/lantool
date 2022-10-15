package cmd

import (
	"fmt"
	"lantool/ddns"
	"os"
	"strings"

	"github.com/urfave/cli"
)

var DDNSFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "ak",
		Value: "",
		Usage: "AccessKeyID",
	},
	cli.StringFlag{
		Name:  "sk",
		Value: "",
		Usage: "AccessKeyID",
	},
	cli.StringFlag{
		Name:  "domain",
		Value: "",
		Usage: "domain",
	},
	cli.StringFlag{
		Name:  "rr",
		Value: "",
		Usage: "",
	},
}

func DDNSAction(c *cli.Context) error {
	if c.String("ak") == "" || c.String("sk") == "" || c.String("domain") == "" || c.String("rr") == "" {
		exeName := os.Args[0]
		commonUsage := []string{
			fmt.Sprintf(`%s ddns --ak xxxx --sk xxxx --domain xxx --rr xxx`, exeName),
		}
		return fmt.Errorf("you can use like \n%s", strings.Join(commonUsage, "\n"))
	}

	ddns := ddns.NewDDNS(c.String("ak"), c.String("sk"))

	return ddns.Run(c.String("domain"), c.String("rr"))
}
