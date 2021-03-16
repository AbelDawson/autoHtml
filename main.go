package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func main() {
	f, err := ioutil.ReadFile("/home/abel/Downloads/Pesto Prosciutto Bread.txt")
	check(err)

	inFile, err := os.Create("/home/abel/programming/autohtml/Pesto.html")
	check(err)

	defer inFile.Close()
	temp := getRecipeArray(string(f))
	recipeArray := temp[1:]

	var data string
	data = "<br /><br />\n" + recipeArray[0][:len(recipeArray[0]) - 4] + "\n"
	data += "=================================\n" + "<div class=\"rw-ui-container\"></div><br /><br /><br>===========<br /><br /><br>\n"
	data += recipeArray[1][:len(recipeArray[1]) - 4] + " | " + recipeArray[2][:len(recipeArray[2]) - 4] + "<br /><br /><br>\n\n"

	for i := 3; i <= len(recipeArray) - 4; i++ {
		data += "* " + recipeArray[i][:len(recipeArray[i]) - 2] + "<br /><br />\n\n"
	}
	data += "<div class=\"rw-ui-container\"></div><br /><br /><br>===========<br /><br /><br>\n"
	data += "<p>" + "**" + recipeArray[len(recipeArray)-3][:len(recipeArray[len(recipeArray)-3]) - 4] + "**" + "</p>\n\n"

	data += recipeArray[len(recipeArray)-2] + "<p>" + recipeArray[len(recipeArray)-1] + "--<br /><br />i <p>"
	data += "<script type=\"text/javascript\">(function(d, t, e, m){ window.RW_Async_Init = function(){RW.init({huid: \"473336\",uid: \"e7a14aa63aeab1844cd78f81639c3190\",source: \"website\",options: {\"size\": \"medium\",\"style\": \"oxygen\",\"isDummy\": false}});RW.render();};var rw, s = d.getElementsByTagName(e)[0], id = \"rw-js\",l = d.location, ck = \"Y\" + t.getFullYear() +\"M\" + t.getMonth() + \"D\" + t.getDate(), p = l.protocol,f = ((l.search.indexOf(\"DBG=\") > -1) ? \"\" : \".min\"),a = (\"https:\" == p ? \"secure.\" + m + \"js/\" : \"js.\" + m);if (d.getElementById(id)) return;rw = d.createElement(e);rw.id = id; rw.async = true; rw.type = \"text/javascript\";rw.src = p + \"//\" + a + \"external\" + f + \".js?ck=\" + ck;s.parentNode.insertBefore(rw, s);}(document, new Date(), \"script\", \"rating-widget.com/\"));</script></p>"

	out, err := inFile.WriteString(data)
	check(err)
	fmt.Printf("Wrote %d bytes\n", out)

	inFile.Sync()

}

func getRecipeArray(data string) []string {
	var ret []string

	temp := strings.Split(data, ";")
	for _, element := range temp {
		switch {
		case strings.Contains(element, "Instructions"):
			ret = append(ret, newBr(element))
		case strings.Contains(element, "Ingredients"):
			ret = append(ret, newBr(element))
		default:
			ret = append(ret, element)
		}
	}
	return ret
}

func newBr(data string) string{
	var ret string
	temp := strings.Split(data, "\\")
	for i, element := range temp {
		if i == 0{
			ret += "**" + element + "**" + "<br />"
		}else {
			ret += element + "<br />"
		}
	}
	return ret
}
