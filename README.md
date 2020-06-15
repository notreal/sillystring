# Silly String
Add visual flourish to text via CLI or web

Usage:
```bash
$ ./sillystring -h
Usage of ./sillystring:
  -a	acute
  -c	caron
  -d	diaeresis
  -dg
    	double_grave
  -g	grave
  -p int
    	port of server, default 8080 (default 8080)
  -s	run http server
  -t	tilde

$ ./sillystring -a Acute
Áćúté

$ ./sillystring -s
Serving at localhost:8080
```

Web interface looks like:
![Web interface screenshot](sillystring_screen.png "Web interface screenshot")