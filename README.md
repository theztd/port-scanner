# portScaner

Scan defined ip's for open ports


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

# Own template

Simply add your template file to the templates directory. And provide your template name as a parametr
