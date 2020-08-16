var fs = require("fs");
var text = fs.readFileSync("../added.txt");
var textByLine = text.split("\n")
added = $('#added')
for(i = 0; i < textByLine.length; i++) {
    added.innerHTML += "<p>" +`${i} `+ arr[i] + "</p><br>";
}
