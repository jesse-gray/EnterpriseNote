//Loads notes
function loadRSS() {
    //Use CORS API website as proxy to retrieve XML file
    var proxy = 'http://';
    var url = "localhost:8000/api/notes";

    //Declare XMLHttpRequest Object
    var xmlhttp = new XMLHttpRequest();
    //Send a request from Client side to Server to retrieve the xml document
    xmlhttp.open("GET", proxy + url, true);
    xmlhttp.send();
    //Check if the entire xml document has been received? If so, process it.
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
            yourOutput += '<div class="noteCard"><h3>' + arr[i].noteid + '</h3><p>' + arr[i].notetext + '</p></div><button id="updateBtn" onclick="location.href="createNote" type="button">Update</button><button id="deleteButton">Delete</button>';
        } else {
            shareOutput += '<div class="noteCard"><h3>' + arr[i].noteid + '</h3><p>' + arr[i].notetext + '</p></div><button id="updateBtn">Update</button>';
        }
    }
    document.getElementById("yourNotes").innerHTML = yourOutput;
    document.getElementById("sharedNotes").innerHTML = shareOutput;
}