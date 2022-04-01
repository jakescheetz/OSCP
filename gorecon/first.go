package main

// imports
import(
	"fmt"
)

//func to print ascii banner
func banner(){
    
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


    fmt.Println(colorYellow + "         *                 *                  *              *        ");
	fmt.Println("                             "+whitebg+colorGreen+"Go"+colorYellow+"Scan"+reset+colorYellow+"                    *             *");
	fmt.Println("                        *            *                             \033[37m___");
	fmt.Println(colorReset + "  \033[33m*               \033[33m*                                          \033[31m|     \033[37m| |");
	fmt.Println(colorReset + "        \033[33m*"+"              \033[37m_________\033[31m##"+"                 \033[33m*"+"        \033[31m/"+colorRed+" \\"+"    \033[37m| |"); //working
	fmt.Println(colorReset + "                      "+"\033[33m@"+ colorWhite + "\\\\\\\\\\\\\\\\\\" + colorRed + "##    \033[33m*     |              "+"\033[31m|\033[37m--o\033[31m|"+"\033[37m===|-|");
	fmt.Println(colorReset + "  \033[33m*                  " + colorYellow + "@@@" + colorWhite + "\\\\\\\\\\\\\\\\" + colorRed + "##" + colorWhite + "\\" + colorGreen + "       \\" + colorYellow + "|" + colorGreen + "/" + colorYellow + "|" + colorGreen + "/            " + colorRed + "|" + colorWhite + "---" + colorRed + "|   " + colorWhite + "|" + colorCyan + "j" + colorWhite + "|");
	fmt.Println(colorYellow + "                    @@ @@"+colorWhite+"\\\\\\\\\\\\\\\\\\\\\\    "+colorGreen+"\\"+colorYellow+"|"+colorGreen+"\\\\"+colorYellow+"|"+colorGreen+"//"+colorYellow+"|"+colorGreen+"/     "+colorYellow+"*   "+colorRed+"/     \\  "+colorWhite+"|"+colorCyan+"a"+colorWhite+"| ");
	fmt.Println(colorYellow + "             *     @@@@@@@"+colorWhite+"\\\\\\\\\\\\\\\\\\\\\\    "+colorGreen+"\\"+colorYellow+"|"+colorGreen+"\\"+colorYellow+"|"+colorGreen+"/"+colorYellow+"|"+colorGreen+"/         "+colorRed+"|  U    | "+colorWhite+"|"+colorCyan+"k"+colorWhite+"| ");
	fmt.Println(colorYellow + "                  @@@@@@@@@----------|    "+colorGreen+"\\\\"+colorYellow+"|"+colorGreen+"//          "+colorRed+"|  "+colorWhite+"S    "+colorRed+"|"+colorWhite+"=|"+colorCyan+"e"+colorWhite+"| ");
	fmt.Println(colorBlue + "       __         "+colorYellow+"@@ @@@ @@__________|     "+colorGreen+"\\"+colorYellow+"|"+colorGreen+"/           "+colorRed+"|  "+colorBlue+"A    "+colorRed+"| "+colorWhite+"| | ");
	fmt.Println(colorReset+"  "+colorBlue+"____|_"+colorWhite+"@"+colorBlue+"|_       "+colorYellow+"@@@@@@@@@__________|     "+colorGreen+"\\"+ colorYellow+"|"+colorGreen+"/           " + colorRed + "|_______| "+colorWhite+"|"+colorCyan+"S"+colorWhite+"| ");
	fmt.Println(colorWhite + "=" + colorBlue + "|__ _____ |" + colorWhite + "=     " + colorYellow + "@@@@ " + colorWhite+ "." + colorYellow + "@@@__________|      |             " + colorRed + "|" + colorYellow + "@" + colorRed + "| |" + colorYellow + "@"+colorRed+"|  "+colorWhite+"|_| ");
	fmt.Println(colorReset + colorGreen + "____" + colorWhite + "0" + colorGreen +"_____" + colorWhite + "0" + colorGreen+ "__\\|/__" + colorYellow + "@@@@" + colorWhite + "__" + colorYellow + "@@@" + colorYellow + "__________|" + colorGreen + "_\\|/__" + colorYellow + "|" + colorGreen+ "___\\|/__\\|/___________" + colorWhite + "|" + colorGreen + "_" + colorWhite+ "|" + colorGreen + "_"+ colorReset);
 	fmt.Println(colorGreen + "  \\|/                "+colorWhite+"/ /       "+colorGreen+"\\|/                              \\|/    ");
 	fmt.Println(colorYellow + "+"+colorBlue+"---------------------------------------------------------------------"+colorYellow+"+");
 	fmt.Println(colorBlue + "|   "+colorGreen+"Go"+colorYellow+"Scan"+colorWhite+"- fast network enumeration (∩｀-´)⊃"+colorYellow+"━   "+colorPurple+"☆"+colorCyan+"ﾟ"+colorYellow+"."+colorPurple+"*"+colorGreen+"･"+colorYellow+"｡"+colorCyan+"ﾟ   "+colorRed+"           "+colorBlue+"|");
 	fmt.Println(colorBlue + "|   Written by: "+bold+underline+colorWhite+"Jake Scheetz"+reset+"          "+colorGreen+"<you>          "+colorPurple+"☆"+colorCyan+"ﾟ"+colorYellow+"."+colorPurple+"*"+colorGreen+"･"+colorYellow+"｡"+colorCyan+"ﾟ          "+colorBlue+"|");
 	fmt.Println(colorBlue + "|   "+colorCyan+"Twitter"+colorWhite+": "+colorYellow+"@"+bold+colorWhite+"FindingUrPasswd"+reset+"                 "+colorPurple+"<cyberMagic>    "+colorWhite+"("+colorCyan+"ᵟຶ︵ ᵟຶ"+colorWhite+") "+colorBlue+"|");
 	fmt.Println(colorBlue + "|   "+ colorRed+"You"+colorWhite+"Tube: "+bold+"youtube.com/c/FindingUrPasswd"+reset+"                   "+colorRed+"<badGuys>"+colorBlue+"|");
 	fmt.Println(colorYellow+"+"+colorBlue+"---------------------------------------------------------------------"+colorYellow+"+"+colorReset);
}

func main (){

    banner();
	
	
} //end main



/*
fmt.Println(colorBlue + "|   "+colorGreen+"Go"+colorYellow+"Scan"+colorWhite+"- fast network enumeration (∩｀-´)⊃"+colorYellow+"━   "+colorPurple+"☆"+colorCyan+"ﾟ"+colorYellow+"."+colorPurple+"*"+colorGreen+"･"+colorYellow+"｡"+colorCyan+"ﾟ   "+colorRed+"<targets>  "+colorBlue+"|");
fmt.Println(colorBlue + "|   "+colorWhite+"Written by: "+bold+underline+"Jake Scheetz"+reset+"          "+colorGreen+"<you>      "+colorPurple+"<cyber>              "+colorBlue+"|");
fmt.Println(colorBlue + "|   "+colorCyan+"Twitter"+colorWhite+": "+colorYellow+"@"+colorWhite+bold+"FindingUrPasswd"+reset+"                                         "+colorBlue+"|");
fmt.Println(colorBlue + "|   "+ colorRed+"You"+colorWhite+"Tube: "+bold+"youtube.com/c/FindingUrPasswd"+reset+"                            "+colorBlue+"|");

cyber magic --> "+colorPurple+"☆"+colorCyan+"ﾟ"+colorYellow+"."+colorPurple+"*"+colorGreen+"･"+colorYellow+"｡"+colorCyan+"ﾟ
*/