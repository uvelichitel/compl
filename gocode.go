package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"code.google.com/p/goplan9/plan9/acme"
)



var insertfirst []byte
var q0exp, q1exp int

var (
	g_is_server = flag.Bool("s", false, "run a server instead of a client")
	g_input     = flag.String("in", "", "use this file instead of stdin input")
	g_sock      = create_sock_flag("sock", "socket type (unix | tcp)")
	g_addr      = flag.String("addr", "localhost:37373", "address for tcp socket")
)

func get_socket_filename() string {
	user := os.Getenv("USER")
	if user == "" {
		user = "all"
	}
	return filepath.Join(os.TempDir(), fmt.Sprintf("gocode-daemon.%s", user))
}


func show_usage() {
	fmt.Fprintf(os.Stderr,
		"Usage: %s [-s] [-in=<path>] [-sock=<type>] [-addr=<addr>]\n"+
			"       <command> [<args>]\n\n",
		os.Args[0])
	fmt.Fprintf(os.Stderr,
		"Flags:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr,
		"\nCommands:\n"+
			"  autocomplete [<path>] <offset>     main autocompletion command\n"+
			"  close                              close the gocode daemon\n"+
			"  status                             gocode daemon status report\n"+
			"  drop-cache                         drop gocode daemon's cache\n"+
			"  set [<name> [<value>]]             list or set config options\n")
}

func main() {
	flag.Usage = show_usage
	flag.Parse()
	


	var retval int
	if *g_is_server {
		retval = do_server()
		os.Exit(retval)
	} else {
		retval = do_client()
	
	
if needinsert {
	carrent_win.Addr( "#%v,#%v", q0exp ,q1exp)
	carrent_win.Write("data", insertfirst)
	_, q1exp, _ = carrent_win.ReadAddr()
	


for e := range compl_win.EventChan() {
	if e.C1 == 'K' {
	inp := e.Text
//	compl_win.CloseFiles()
	ctl := carrent_win.Getctl()
	carrent_win, _  = acme.Open(77, ctl)
	carrent_win.Addr( "#%v,#%v", q1exp ,q1exp)
	carrent_win.Write("data", inp)
	carrent_win.Ctl("dot=addr")
	os.Exit(retval)
	}

		switch e.C2 {
		case 'x': // execute
				if string(e.Text) == "Del" {
				compl_win.Ctl("delete")
			}
		case 'L':
		insert := e.Text

	carrent_win.Addr( "#%v,#%v", q0exp ,q1exp)
	carrent_win.Write("data", insert)
	_, q1exp, _ = carrent_win.ReadAddr()
		
		
		}
	
		compl_win.WriteEvent(e)
		if e.C2 == 'D' {
		compl_win.WriteEvent(e)
		os.Exit(retval)
		}	
	}
}
	os.Exit(retval)
}
}

