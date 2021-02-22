package console

import "github.com/urfave/cli"

func StringFlag(name, value, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func BoolFlag(name string, value *bool, usage string) cli.BoolFlag {
	return cli.BoolFlag{
		Name:        name,
		Destination: value,
		Usage:       usage,
	}
}

var systemd = `[Unit]
Description={{.Name}}
After=syslog.target
After=network.target
After=mariadb.service mysqld.service postgresql.service memcached.service redis.service

[Service]
# Modify these two values and uncomment them if you have
# repos with lots of files and get an HTTP error 500 because
# of that
###
#LimitMEMLOCK=infinity
#LimitNOFILE=65535
Type=simple
User=root
Group=root
WorkingDirectory={{.Dir}}
ExecStart={{.Exec}}
Restart=always
Environment=PWD={{.Dir}}

# Some distributions may not support these hardening directives. If you cannot start the service due
# to an unknown option, comment out the ones not supported by your version of systemd.
ProtectSystem=full
PrivateDevices=yes
PrivateTmp=yes
NoNewPrivileges=true

[Install]
WantedBy=multi-user.target
`

type data struct {
	Name string
	Dir  string
	Exec string
}
