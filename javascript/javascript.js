//Loads notes
function loadNotes() {
    var url = "http://localhost:8000/api/notes";

    //Declare XMLHttpRequest Object
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.open("GET", url, true);
    xmlhttp.send();
    xmlhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var myArr = JSON.parse(this.responseText);
            //Load XML document as XML format and process
            processNotes(myArr);
        }
    };
}

//Finds relevant data in document and displays it on the webpage
function processNotes(arr) {
    var yourOutput = "<h2>Your Notes</h2>";
    var shareOutput = "<h2>Notes that have been shared with you</h2>";
    for (var i = 0; i < arr.length; i++) {
        //Display extracted article into the divs
        if (arr[i].authorid == 1) { //TODO replace with currently logged in
            yourOutput += '<div class="container-fluid"><div class="card text-white bg-secondary mb-3"><div class="card-body"><h3>NoteID: ' + arr[i].noteid + '</h3><p class="card-text">Note Text: ' + arr[i].notetext + '</p><button class="btn btn-light mr-1" id="updateBtn" value="' + arr[i].noteid + '" onclick="updateNote(this.value)" type="button">Update</button><button class="btn btn-light mr-1" id="deleteButton" value="' + arr[i].noteid + '" onclick="deleteNote(this.value)" type="button">Delete</button><button class="btn btn-light" id="updatePerms" value="' + arr[i].noteid + '" onclick="updatePerms(this.value)" type="button">Update Permissions</button></div></div></div>';
        } else {
            shareOutput += '<div class="container-fluid"><div class="card text-white bg-secondary mb-3"><div class="card-body"><h3 class="card-title">NoteID: ' + arr[i].noteid + '</h3><p class="card-text">Note Text: ' + arr[i].notetext + '</p><button class="btn btn-light" id="updateBtn" value="' + arr[i].noteid + '" onclick="updateNote(this.value)" type="button">Update</button></div></div></div>';
        }
    }
    document.getElementById("yourNotes").innerHTML = yourOutput;
    document.getElementById("sharedNotes").innerHTML = shareOutput;
}

//Load users
function loadUsers() {
    var url = "http://localhost:8000/api/users";

    //Declare XMLHttpRequest Object
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.open("GET", url, true);
    xmlhttp.send();
    xmlhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var myArr = JSON.parse(this.responseText);
            //Load XML document as XML format and process
            processUsers(myArr);
        }
    };
}

//Finds relevant data in document and displays it on the webpage
function processUsers(arr) {
    var output = "<h2>All Users:</h2>";
    for (var i = 0; i < arr.length; i++) {
        //Display extracted article into the divs
        output += '<div class="container-fluid"><div class="card text-white bg-secondary mb-3"><div class="card-body"><h5 class="card-title">User ID: ' + arr[i].userid + '</h5><p class="card-text">Name: ' + arr[i].firstname + ' ' + arr[i].lastname + '</p></div></div></div>';
    }
    document.getElementById("userList").innerHTML = output;
}

function loadPage() {
    document.getElementById("formnoteid").value = localStorage.getItem("noteid");
}

function updateNote(noteID) {
    localStorage.setItem("noteid", noteID)
    location.href = 'updateNote';
}

function deleteNote(noteID) {
    if (window.confirm("Are you sure you want to delete note " + noteID + "?")) {
        $.ajax({
            type: 'DELETE',
            url: 'http://localhost:8000/api/notes/' + noteID
        });
        location.href = 'home';
    }
}

function updatePerms(noteID) {
    localStorage.setItem("noteid", noteID)
    location.href = 'updatePerms';
}

//Login Form
$('#loginForm').on('submit', function(event) {
    event.preventDefault();

    var userid = $('#userid').val()
    var password = $('#password').val()
    var json = JSON.stringify({ "userid": userid, "password": password })
    console.log(json)
    $.ajax({
        type: 'POST',
        url: 'http://localhost:8000/api/login',
        dataType: 'json',
        data: json,
        contentType: 'application/json',
        success: function(data) {
            localStorage.setItem("currentlyloggedin", data)
        }
    });
});