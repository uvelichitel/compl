package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/rpc"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"code.google.com/p/goplan9/plan9/acme"
)

var compl_win *acme.Win
var compl_id int
var afile *acmeFile
var carrent_win *acme.Win
var carrent_win_id int
var needinsert bool

func do_client() int {
	addr := *g_addr
	if *g_sock == "unix" {
		addr = get_socket_filename()
	}

	// client
	client, err := rpc.Dial(*g_sock, addr)
	if err != nil {
		if *g_sock == "unix" && file_exists(addr) {
			os.Remove(addr)
		}

		err = try_run_server()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return 1
		}
		client, err = try_to_connect(*g_sock, addr)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return 1
		}
	}
	defer client.Close()

	if flag.NArg() > 0 {
		switch flag.Arg(0) {
		case "autocomplete":
			cmd_auto_complete(client)
		case "cursortype":
			cmd_cursor_type_pkg(client)
		case "close":
			cmd_close(client)
		case "status":
			cmd_status(client)
		case "drop-cache":
			cmd_drop_cache(client)
		case "set":
			cmd_set(client)
		}
	}else{
	cmd_auto_complete(client)
	}
	return 0
}

func try_run_server() error {
	path := get_executable_filename()
	args := []string{os.Args[0], "-s", "-sock", *g_sock, "-addr", *g_addr}
	cwd, _ := os.Getwd()
	procattr := os.ProcAttr{Dir: cwd, Env: os.Environ(), Files: []*os.File{nil, nil, nil}}
	p, err := os.StartProcess(path, args, &procattr)

	if err != nil {
		return err
	}
	return p.Release()
}

func try_to_connect(network, address string) (client *rpc.Client, err error) {
	t := 0
	for {
		client, err = rpc.Dial(network, address)
		if err != nil && t < 1000 {
			time.Sleep(10 * time.Millisecond)
			t += 10
			continue
		}
		break
	}

	return
}

func prepare_file_filename_cursor() ([]byte, string, int) {
	var file []byte
	var err error

	if *g_input != "" {
		file, err = ioutil.ReadFile(*g_input)
	} else {
		file, err = ioutil.ReadAll(os.Stdin)
	}

	if err != nil {
		panic(err.Error())
	}

	var skipped int
	file, skipped = filter_out_shebang(file)

	filename := *g_input
	cursor := -1

	offset := ""
	switch flag.NArg() {
	case 2:
		offset = flag.Arg(1)
	case 3:
		filename = flag.Arg(1) // Override default filename
		offset = flag.Arg(2)
	}

	if offset != "" {
		if offset[0] == 'c' || offset[0] == 'C' {
			cursor, _ = strconv.Atoi(offset[1:])
			cursor = char_to_byte_offset(file, cursor)
		} else {
			cursor, _ = strconv.Atoi(offset)
		}
	}

	cursor -= skipped
	if filename != "" && !filepath.IsAbs(filename) {
		cwd, _ := os.Getwd()
		filename = filepath.Join(cwd, filename)
	}
	return file, filename, cursor
}

//-------------------------------------------------------------------------
// commands
//-------------------------------------------------------------------------

func cmd_status(c *rpc.Client) {
	fmt.Printf("%s\n", client_status(c, 0))
}

func cmd_auto_complete(c *rpc.Client) {
        var env gocode_env
        env.get()
	var args, reply int
	var err error
	args = 0
	err = c.Call("RPC.RPC_setid", &args, &reply)
	compl_id = reply
	compl_win, err = acme.Open(compl_id, nil)
	if err != nil {
	compl_win, _ = acme.New()
	args = compl_win.GetId() 
	err = c.Call("RPC.RPC_setid", &args, &reply)}
//for acme

	var src []byte
	var searchpos  int
	var fname string

			if afile, err = acmeCurrentFile(); err != nil {
				fmt.Printf("%v", err)
			}
			fname, src, searchpos = afile.name, afile.body, afile.offset 
			compl_win.Name("%v+completions", fname)
			compl_win.Addr(",")
			compl_win.Write("data", nil)
//for acme

	
	write_candidates(client_auto_complete(c, src, fname, searchpos, env))
}

func cmd_cursor_type_pkg(c *rpc.Client) {
	var args, reply int
	var err error
	args = 0
	err = c.Call("RPC.RPC_setid", &args, &reply)
	compl_id = reply
	compl_win, err = acme.Open(compl_id, nil)
	if err != nil {
	compl_win, _ = acme.New()
	args = compl_win.GetId() 
	err = c.Call("RPC.RPC_setid", &args, &reply)}	
	//for acme
	
	var src []byte
	var searchpos  int
	var fname string

			if afile, err = acmeCurrentFile(); err != nil {
				fmt.Printf("%v", err)
			}
			fname, src, searchpos = afile.name, afile.body, afile.offset  
//for acme
	typ, pkg := client_cursor_type_pkg(c, src, fname, searchpos)
	fmt.Printf("%s,,%s\n", typ, pkg)
}

func cmd_close(c *rpc.Client) {
	client_close(c, 0)
}

func cmd_drop_cache(c *rpc.Client) {
	client_drop_cache(c, 0)
}

func cmd_set(c *rpc.Client) {
	switch flag.NArg() {
	case 1:
		fmt.Print(client_set(c, "\x00", "\x00"))
	case 2:
		fmt.Print(client_set(c, flag.Arg(1), "\x00"))
	case 3:
		fmt.Print(client_set(c, flag.Arg(1), flag.Arg(2)))
	}
}

func write_candidates(candidates []candidate, num int) {
var warnings, completions, messages string

	compl_win.Ctl("cleartag")

	if candidates != nil {
		
	messages	= fmt.Sprintf("Found %d candidates:\n", len(candidates))
	
	insertfirst	= []byte(candidates[0].Name)

	for _, c := range candidates {
		abbr := fmt.Sprintf("%s %s %s", c.Class, c.Name, c.Type)
		if c.Class == decl_func {
			abbr = fmt.Sprintf("%s %s%s", c.Class, c.Name, c.Type[len("func"):])
		}
		completions = completions + fmt.Sprintf("  %s\n", abbr)
	}
needinsert = true
compl_win.Write("body", []byte(messages))
compl_win.Write("body", []byte(completions))

} else {
needinsert = false
warnings = fmt.Sprintf("Nothing to complete.\n")
compl_win.Write("body", []byte(warnings))
}
compl_win.Ctl("clean")
}
