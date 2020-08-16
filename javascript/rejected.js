var fs = require("fs");
var text = fs.readFileSync("../rejected.txt");
var textByLine = text.split("\n")
rejected= $('#rejected')
for(i = 0; i < textByLine.length; i++) {
    added.innerHTML += "<p>" +`${i} `+ arr[i] + "</p><br>";
}
