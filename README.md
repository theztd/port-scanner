[![Release Go project](https://github.com/theztd/port-scanner/actions/workflows/release.yml/badge.svg)](https://github.com/theztd/port-scanner/actions/workflows/release.yml)
[![Build and deploj](https://github.com/theztd/port-scanner/actions/workflows/build_and_deploy.yml/badge.svg)](https://github.com/theztd/port-scanner/actions/workflows/build_and_deploy.yml)

# portScaner

Scan defined ip's for open ports

![./images/printscreen.png](./images/printscreen.png)

Look at [./example-status.htm](example-status.htm) file. It is an example of the result...

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
    	Name of the output template (build in are: json, prometheus, html). (default "html")
  -template-file string
    	Path to the custom template file.
```

## Examples


### Generate report as a html page

You can send this html page via email, put it to the cloudflare, ...

```bash
portscanner -quick -in servers.txt > /var/www/nginx/status/openports-$(date +%F).html
```


### Generate json output

```bash
portscanner -quick -in servers.txt -template json  > data.json
```
Than you can sent it via curl to any api endpoint or object storage.

### Generate prometheus output

```bash
portscanner -quick -in servers.txt -template prometheus  > textfile_portscanner.prom
```

### Own template

Simply add your template by parametr -template-file. And provide your template name as a parametr ;-)

```bash
portscanner -quick -in servers.txt -template template.json -template-file ./my_templates/template.json
```
