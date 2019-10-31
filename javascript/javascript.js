//Loads notes
function loadRSS() {
    var url = "http://localhost:8000/api/notes";

    //Declare XMLHttpRequest Object
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.open("GET", url, true);
    xmlhttp.send();
    xmlhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var myArr = JSON.parse(this.responseText);
            //Load XML document as XML format and process
            processJSON(myArr);
        }
    };
}

//Finds relevant data in document and displays it on the webpage
function processJSON(arr) {
    var yourOutput = "<h2>Your Notes</h2>";
    var shareOutput = "<h2>Notes that have been shared with you</h2>";
    for (var i = 0; i < arr.length; i++) {
        //Display extracted article into the divs
        if (arr[i].authorid == 1) { //TODO replace with currently logged in
            yourOutput += '<div class="noteCard"><h3>NoteID: ' + arr[i].noteid + '</h3><p>Note Text: ' + arr[i].notetext + '</p></div><button id="updateBtn" value="' + arr[i].noteid + '" onclick="location.href=\'updateNote\'" type="button">Update</button><button id="deleteButton" value="' + arr[i].noteid + '">Delete</button>';
        } else {
            shareOutput += '<div class="noteCard"><h3>NoteID: ' + arr[i].noteid + '</h3><p>Note Text: ' + arr[i].notetext + '</p></div><button id="updateBtn" value="' + arr[i].noteid + '" onclick="location.href=\'updateNote\'" type="button">Update</button>';
        }
    }
    document.getElementById("yourNotes").innerHTML = yourOutput;
    document.getElementById("sharedNotes").innerHTML = shareOutput;
}