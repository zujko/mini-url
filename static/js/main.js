document.querySelector('#shortenbtn').onclick = function() {
    shortenUrl();
}

function shortenUrl() {
    var urldata = document.querySelector('#url-field').value;
    var xhr = new XMLHttpRequest();

    if(urldata != "") { 
        toggleSpinner();
        xhr.open("POST","/shorten");
        xhr.send(JSON.stringify({
            url: urldata,
        }));
        xhr.onreadystatechange = function() {
            if(xhr.readyState === 4) {
                toggleSpinner();
                getShortUrl(xhr);
            }
        }
    }
}

function toggleSpinner() {
    var element = document.getElementById('spinner-main')
    if(element) {
        var display = element.style.display;
        if(display === "none") {
            element.style.display = "inline-block";
        } else {
            element.style.display = "none";
        }
    }
}

function getShortUrl(xhrReq) {
    console.log(xhrReq.responseText)
    var data = JSON.parse(xhrReq.responseText);
    if(data.data === "invalid") {
        alert("Invalid URL");
    } 
}