var fs = require("fs");
var text = fs.readFileSync("../internship.txt");
var textByLine = text.split("\n")
intern = $('#internship')
for(i = 0; i < textByLine.length; i++) {
    intern.innerHTML += "<p>" +`${i} `+ arr[i] + "</p><br>";
}

