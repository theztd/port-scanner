[![Release Go project](https://github.com/theztd/port-scanner/actions/workflows/release.yml/badge.svg)](https://github.com/theztd/port-scanner/actions/workflows/release.yml)
[![Build and deploj](https://github.com/theztd/port-scanner/actions/workflows/build_and_deploy.yml/badge.svg)](https://github.com/theztd/port-scanner/actions/workflows/build_and_deploy.yml)

# portScaner

Scan defined ip's for open ports

![./images/printscreen.png](./images/printscreen.png)

Look at [./status.html](status.html) file whih is example result...

## Usage

```bash
Usage of portscanner:
  -full
        Scan everything from 1-65535 (Super slow)
  -in string
        Path to the file with hosts (One line = one host)
  -out string
        Path to the output file.
  -quick
        Do only fast scan (predefined most common ports)
  -template string
        Name of the output template file (file have to be present under ./template directory).
```

## Examples


### Generate report as a html page

You can send this html page via email, put it to the cloudflare, ...

```bash
portscanner -quick -in servers.txt > /var/www/nginx/status/openports-$(date +%F).html
```


### Generate json output

```bash
portscanner -quick -in servers.txt -template json.json  > data.json
```
Than you can sent it via curl to any api endpoint or object storage.

### Generate prometheus output

```bash
portscanner -quick -in servers.txt -template prometheus.tpl  > textfile_portscanner.prom
```

### Own template

Simply add your template file to the templates directory. And provide your template name as a parametr ;-)
