document.querySelector('#shortenbtn').onclick = function() {
    shortenUrl();
}

$('url-table').on('mouseover', 'td', function() {
    // Store the hovered cell's text in a variable
    var textToMatch = $(this).text();

    // Loop through every `td` element
    $('td').each(function() {
        // Pull selected `td` element's text
        var text = $(this).text();

        // Compare this with initial text and add matching class if it matches
        if (textToMatch === text)
            $(this).parent().addClass('matching');
    });
});

// Mouse out event handler
// This simply removes the matching styling
$('table').on('mouseout', 'td', function() {
    $('.matching').removeClass('matching');
});

/**
 * Makes an XMLHttpRequest to the server to shorten a URL.
 */
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

/**
 * Toggles spinner
 */
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

/**
 * Handle response from server and display data to user.
 * @param {*} xhrReq 
 */
function getShortUrl(xhrReq) {
    console.log(xhrReq.responseText)
    var data = JSON.parse(xhrReq.responseText);
    if(data.error) {
        $("#url-field").notify("Invalid URL!");
    } else {
        displayUrl(data);
    }
}

/**
 * Displays URL data in a table
 * @param {String} data 
 */
function displayUrl(data) {
    // First make sure the table is visible
    var table = document.getElementById('url-table');
    if(table.style.display == "none") {
        table.style.display = "table";
    }
    var tableRef = document.getElementById('url-table').getElementsByTagName('tbody')[0];

    // Insert a row in the table at the last row
    var newRow   = tableRef.insertRow(tableRef.rows.length);
    // Insert a cell in the row at index 0
    newRow.insertCell(0).innerHTML = data.longUrl;
    newRow.insertCell(1).innerHTML = window.location.host+"/"+data.shortUrl;
}