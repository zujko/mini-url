document.querySelector('#shortenbtn').onclick = function() {
    shortenUrl();
}

function shortenUrl() {
    var urldata = document.querySelector('#url-field').value;
    var xhr = new XMLHttpRequest();

    if(urldata != "") { 
        xhr.open("POST","/shorten");
        xhr.send(JSON.stringify({
            url: urldata,
        }));
        xhr.onreadystatechange = function() {
            if(xhr.readyState === 4) {
                getShortUrl(xhr);
            }
            
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