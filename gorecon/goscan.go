package main

// imports
import (
	"fmt"
	"os"
)

//main logic
func main() {

	//if no cmd line args, -h or --help is supplied print help menu, else print the banner to start tool
	if (len(os.Args) == 1) || (os.Args[1] == "-h") || (os.Args[1] == "--help") {
		// TODO: change to function once help menu design is done
		fmt.Println("help menu :)")
		return
	} else {
		//print banner
		banner()
	}

} //end main

// OS commands to execute

// func execLs() {
// 	//try to exec ls -l??
// 	app := "ls"                    //application to us
// 	arg0 := "-l"                   //app args
// 	cmd := exec.Command(app, arg0) //exec the application w/ args and pas it as cmd
// 	stdout, err := cmd.Output()    //if success, std,out get output. if fail, err gets output

// 	//get cmd error output
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Println(string(stdout)) // print successful output
// }

// func execPingsweep() {
// 	app := "nmap"
// 	arg0 := "-sn"
// 	arg1 := "-vv"

// }

//func to print ascii banner
func banner() {

	// color variables
	colorReset := "\033[0m" //reset term color
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorCyan := "\033[36m"
	colorWhite := "\033[37m"
	colorPurple := "\033[35m"
	bold := "\u001b[1m"
	underline := "\u001b[4m"
	reset := "\u001b[0m"
	whitebg := "\u001b[47;1m"

	fmt.Println(colorYellow + "         *                 *                  *              *        ")
	fmt.Println("                             " + bold + whitebg + colorGreen + "Go" + colorYellow + "Scan" + reset + colorYellow + "                    *             *")
	fmt.Println("                        *            *                             " + colorReset + "___")
	fmt.Println(colorReset + "  \033[33m*               \033[33m*                                          \033[31m|     " + colorReset + "| |")
	fmt.Println(colorReset + "        \033[33m*" + "              \033[37m" + colorReset + "_________\033[31m" + colorRed + "##" + "                 \033[33m*" + "        \033[31m/" + colorRed + " \\" + "    " + colorReset + "| |" + reset + "") //working
	fmt.Println(colorReset + "                      " + "\033[33m@" + colorReset + "\\\\\\\\\\\\\\\\\\" + colorRed + "##    \033[33m*     |              " + colorRed + "|" + colorReset + bold + "--o" + reset + colorRed + "|" + colorReset + "===|-|")
	fmt.Println(colorReset + "  \033[33m*                  " + colorYellow + "@@@" + colorReset + "\\\\\\\\\\\\\\\\" + colorRed + "##" + colorReset + "\\" + colorGreen + "       \\" + colorYellow + "|" + colorGreen + "/" + colorYellow + "|" + colorGreen + "/            " + colorRed + "|" + colorReset + bold + "---" + reset + colorRed + "|   " + colorReset + "|" + colorCyan + "j" + colorReset + "|")
	fmt.Println(colorYellow + "                    @@ @@" + colorReset + "\\\\\\\\\\\\\\\\\\\\\\    " + colorGreen + "\\" + colorYellow + "|" + colorGreen + "\\\\" + colorYellow + "|" + colorGreen + "//" + colorYellow + "|" + colorGreen + "/     " + colorYellow + "*   " + colorRed + "/     \\  " + colorReset + "|" + colorCyan + "a" + colorReset + "| ")
	fmt.Println(colorYellow + "             *     @@@@@@@" + colorReset + "\\\\\\\\\\\\\\\\\\\\\\    " + colorGreen + "\\" + colorYellow + "|" + colorGreen + "\\" + colorYellow + "|" + colorGreen + "/" + colorYellow + "|" + colorGreen + "/         " + colorRed + "|  " + "U" + colorRed + "    | " + colorReset + "|" + colorCyan + "k" + colorReset + "| ")
	fmt.Println(colorYellow + "                  @@@@@@@@@----------|    " + colorGreen + "\\\\" + colorYellow + "|" + colorGreen + "//          " + colorRed + "|  " + colorReset + bold + "S    " + reset + colorRed + "|" + colorReset + "=|" + colorCyan + "e" + colorReset + "| ")
	fmt.Println(colorBlue + "       __         " + colorYellow + "@@ @@@ @@__________|     " + colorGreen + "\\" + colorYellow + "|" + colorGreen + "/           " + colorRed + "|  " + colorBlue + "A    " + colorRed + "| " + colorReset + "| | ")
	fmt.Println(colorReset + "  " + colorBlue + "____|_" + colorReset + "@" + colorBlue + "|_       " + colorYellow + "@@@@@@@@@__________|     " + colorGreen + "\\" + colorYellow + "|" + colorGreen + "/           " + colorRed + "|_______| " + colorReset + "|" + colorCyan + "S" + colorReset + "| ")
	fmt.Println(colorReset + "=" + colorBlue + "|__ _____ |" + colorReset + "=     " + colorYellow + "@@@@ " + colorReset + "." + colorYellow + "@@@__________|      |             " + colorRed + "|" + colorYellow + "@" + colorRed + "| |" + colorYellow + "@" + colorRed + "|  " + colorReset + "|_| ")
	fmt.Println(colorReset + colorGreen + "____" + colorReset + "0" + colorGreen + "_____" + colorReset + "0" + colorGreen + "__\\|/__" + colorYellow + "@@@@" + colorReset + "__" + colorYellow + "@@@" + colorYellow + "__________|" + colorGreen + "_\\|/__" + colorYellow + "|" + colorGreen + "___\\|/__\\|/___________" + colorReset + "|" + colorGreen + "_" + colorReset + "|" + colorGreen + "_" + colorReset)
	fmt.Println(colorGreen + "  \\|/                " + colorReset + "/ /       " + colorGreen + "\\|/                              \\|/    ")
	fmt.Println(colorYellow + "+" + colorBlue + "---------------------------------------------------------------------" + colorYellow + "+")
	fmt.Println(colorBlue + "|   " + colorGreen + "Go" + colorYellow + "Scan" + colorReset + "- fast network enumeration (∩｀-´)⊃" + colorYellow + "━   " + colorPurple + "☆" + colorCyan + "ﾟ" + colorYellow + "." + colorPurple + "*" + colorGreen + "･" + colorYellow + "｡" + colorCyan + "ﾟ   " + colorRed + "           " + colorBlue + "|")
	fmt.Println(colorBlue + "|   Written by: " + bold + underline + colorWhite + "Jake Scheetz" + reset + "          " + colorGreen + "<you>          " + colorPurple + "☆" + colorCyan + "ﾟ" + colorYellow + "." + colorPurple + "*" + colorGreen + "･" + colorYellow + "｡" + colorCyan + "ﾟ          " + colorBlue + "|")
	fmt.Println(colorBlue + "|   " + colorCyan + "Twitter" + colorWhite + ": " + colorYellow + "@" + bold + colorWhite + "FindingUrPasswd" + reset + "                 " + colorPurple + "<cyberMagic>    " + colorReset + "(" + colorCyan + "ᵟຶ︵ ᵟຶ" + colorReset + ") " + colorBlue + "|")
	fmt.Println(colorBlue + "|   " + colorRed + "You" + colorReset + "Tube: " + bold + "youtube.com/c/FindingUrPasswd" + reset + "                   " + colorRed + "<badGuys>" + colorBlue + "|")
	fmt.Println(colorYellow + "+" + colorBlue + "---------------------------------------------------------------------" + colorYellow + "+" + colorReset)
}

func helpMenu() {

}
