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
	f, err := ioutil.ReadFile("/home/abel/Downloads/Budha Bowl.txt")
	check(err)

	inFile, err := os.Create("/home/abel/Documents/recipeHtml/BudhaBowl.html")
	check(err)

	defer inFile.Close()
	recipeArray := getRecipeArray(string(f))
	var data []string
	data = append(data, "<br /><br />\n")
	data = append(data, recipeArray[0][:len(recipeArray[0]) - 1])
	data = append(data, "================================\n\n")
	data = append(data, "<div class=\"rw-ui-container\"></div><br /><br /><br>===========<br /><br /><br>\n")

	data = append(data, recipeArray[1][:len(recipeArray[1]) - 4] + " | " + recipeArray[2][:len(recipeArray[2]) - 4] + "<br /><br /><br>\n\n")

	data = append(data, "* " + recipeArray[3][:len(recipeArray[3]) - 3] + "<br /><br />\n\n")
	data = append(data, "* " + recipeArray[4][:len(recipeArray[4]) - 3] + "<br /><br />\n\n")
	data = append(data, "* " + recipeArray[5][:len(recipeArray[5]) - 4] + "<br /><br />\n\n")
	data = append(data, "<p>" + recipeArray[6][:len(recipeArray[6]) - 4] + "</p>\n\n")
	data = append(data, recipeArray[7])
	data = append(data, "<p>")
	data = append(data, recipeArray[8])
	data = append(data, "--<br /><br />")
	data = append(data, "<p>")
	data = append(data, "<script type=\"text/javascript\">(function(d, t, e, m){")
	data = append(data, "// Async Rating-Widget initialization. 25     window.RW_Async_Init = function(){")
	data = append(data, "RW.init({")
	data = append(data, "huid: \"473336\",")
	data = append(data, "uid: \"e7a14aa63aeab1844cd78f81639c3190\",")
	data = append(data, "source: \"website\",")
	data = append(data, "options: {")
	data = append(data, "\"size\": \"medium\",")
	data = append(data, "\"style\": \"oxygen\",")
	data = append(data, "\"isDummy\": false")
	data = append(data, "}")
	data = append(data, "});")
	data = append(data, "RW.render();")
	data = append(data, "};")
	data = append(data, "var rw, s = d.getElementsByTagName(e)[0], id = \"rw-js\",")
	data = append(data, "l = d.location, ck = \"Y\" + t.getFullYear() +")
	data = append(data, "\"M\" + t.getMonth() + \"D\" + t.getDate(), p = l.protocol,")
	data = append(data, "f = ((l.search.indexOf(\"DBG=\") > -1) ? \"\" : \".min\"),")
	data = append(data, "a = (\"https:\" == p ? \"secure.\" + m + \"js/\" : \"js.\" + m);")
	data = append(data, "if (d.getElementById(id)) return;")
	data = append(data, "rw = d.createElement(e);")
	data = append(data, "rw.id = id; rw.async = true; rw.type = \"text/javascript\";")
	data = append(data, "rw.src = p + \"//\" + a + \"external\" + f + \".js?ck=\" + ck;")
	data = append(data, "s.parentNode.insertBefore(rw, s);")
	data = append(data, "}(document, new Date(), \"script\", \"rating-widget.com/\"));</script>")
	data = append(data, "</p>")
	fmt.Printf("%s", strings.Join(data, ""))
}

func getRecipeArray(data string) []string {
	var ret []string

	temp := strings.Split(data, ";")

	for _, element := range temp {
		switch {
		case strings.Contains(element, "Title"):
			ret = append(ret, element)
		case strings.Contains(element, "Author"):
			ret = append(ret, element)
		case strings.Contains(element, "Date"):
			ret = append(ret, element)
		case strings.Contains(element, "Prep"):
			ret = append(ret, element)
		case strings.Contains(element, "Cook"):
			ret = append(ret, element)
		case strings.Contains(element, "Total"):
			ret = append(ret, element)
		case strings.Contains(element, "Servings"):
			ret = append(ret, element)
		case strings.Contains(element, "Instructions"):
			ret = append(ret, newBr(element))
		case strings.Contains(element, "Ingredients"):
			ret = append(ret, newBr(element))
		}
	}
	return ret
}

func newBr(data string) string{
	var ret []string
	temp := strings.Split(data, "\\")
	for _, element := range temp {
		ret = append(ret, element + "<br />")
	}
	return strings.Join(ret, "")
}
