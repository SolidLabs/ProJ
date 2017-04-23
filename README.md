ProJ
====

ProJ (pronounced Pro-J) is a local development project management tool.

It aims to simplify local development setup for new projects, removing the tedious bootstrapping process (creating virtual hosts, setting up the hosts file etc).

*Note*: The project is still in very early alpha, so some featues might be broken/missing.

## Usage

To setup a new project, just run the `proj init` command from inside the project directory.


Init config:

```
Usage:
   init [flags]

Flags:
  -h, --help               help for init
  -n, --host stringArray   Project hostnames
  -i, --ip string          IP address to use for hostname (default "127.0.0.1")
  -p, --port int           Hostname port (default 80)
```

The config takes a hostname, ip address (defaults to `127.0.0.1`) and a port (defaults to `80`).

When running the command, ProJ will create new hosts entries mapping the hostname to the ip address, and create a webserver configuration.

E.G

```bash
$ proj init -n my.project.dev
```

The will add the entry `127.0.0.1 my.project.dev` to your /etc/hosts file, as well as a webservice configuration that binds to the domain `my.project.dev`

### Using multiple hostnames

You can specify the --host (or -n) flag multiple times to add multiple hostnames

```bash
$ proj init -n my.project.dev -n sub.my.project.dev
```

## Supported Webservers

- [x] Nginx
- [ ] Apache

*Note:* Currently only Nginx is supported with a default PHP-FPM configuration. Support for more webservers and configurations will be added soon.