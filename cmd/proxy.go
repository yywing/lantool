package cmd

import (
	"fmt"
	"os"
	"strings"

	"lantool/rotateproxy"

	"github.com/urfave/cli"
)

const (
	rule      = `protocol=="socks5" && "Version:5 Method:No Authentication(0x00)" && after="2022-02-01" && country="CN"`
	pageCount = 5
)

var ProxyFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "email",
		Value: "",
		Usage: "fofa username email",
	},
	cli.StringFlag{
		Name:  "token",
		Value: "",
		Usage: "fofa token",
	},
	cli.StringFlag{
		Name:  "check",
		Value: "https://www.google.com",
		Usage: "proxy check url",
	},
	cli.StringFlag{
		Name:  "listen",
		Value: ":8899",
		Usage: "proxy server listen addr",
	},
	cli.StringFlag{
		Name:  "username",
		Value: "",
		Usage: "proxy auth username",
	},
	cli.StringFlag{
		Name:  "password",
		Value: "",
		Usage: "proxy auth password",
	},
}

func ProxyAction(c *cli.Context) error {
	if c.String("email") == "" || c.String("token") == "" {
		exeName := os.Args[0]
		commonUsage := []string{
			fmt.Sprintf(`%s proxy --email xxxx --token xxxx`, exeName),
		}
		return fmt.Errorf("you can use like \n%s", strings.Join(commonUsage, "\n"))
	}

	baseCfg := rotateproxy.BaseConfig{
		ListenAddr:     c.String("listen"),
		Username:       c.String("username"),
		Password:       c.String("password"),
		SelectStrategy: 3,
	}

	rotateproxy.StartRunCrawler(c.String("token"), c.String("email"), rule, pageCount, "")
	rotateproxy.StartCheckProxyAlive(c.String("check"))
	client := rotateproxy.NewRedirectClient(rotateproxy.WithConfig(&baseCfg))
	client.Serve()
	select {}
}
