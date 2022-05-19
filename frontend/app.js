$(document).ready(function(){
    var apiUrl = "http://localhost:9876/api/inventory/";
    var request = new XMLHttpRequest();
    request.open("GET", apiUrl, false);
    request.send(null);
    var items = JSON.parse(request.responseText);
    for (var i = 0; i < items.length; i++) {
        var display = "Name: " + items[i].itemName + " Description: " + items[i].itemDescription;
        $("#items").append('<div><li>'+display+'</li></div>');
    }

    var request = new XMLHttpRequest();
    request.open("GET", apiUrl + 'recover/', false);
    request.send(null);
    var deletedItems = JSON.parse(request.responseText);
    for (var i = 0; i < deletedItems.length; i++) {
        var display = "Name: " + deletedItems[i].itemName + " Description: " + deletedItems[i].itemDescription + " Comment: " + deletedItems[i].deletionComment;
        $("#deletedItems").append('<div><li>'+display+'</li></div>');
    }

    $("#editButton").click(function() {
        var $inputs = $('#editForm :input');
        var payload = {};
        var itemId = "";
        $inputs.each(function() {
            if (this.name == "itemIndex") {
                itemId = items[parseInt($(this).val())-1].itemId
            } else {
                payload[this.name] = $(this).val();
            }
        });
        var request = new XMLHttpRequest();
        request.open("PATCH", apiUrl + itemId, false);
        request.setRequestHeader('Content-Type', 'application/json');
        request.send(JSON.stringify(payload));
        location.reload();
    });


    $("#deleteButton").click(function() {
        var $inputs = $('#deleteForm :input');
        var payload = {};
        var itemId = "";
        $inputs.each(function() {
            if (this.name == "itemIndex") {
                itemId = items[parseInt($(this).val())-1].itemId
            } else {
                payload[this.name] = $(this).val();
            }
        });
        console.log(itemId);
        console.log(payload);
        var request = new XMLHttpRequest();
        request.open("DELETE", apiUrl + itemId, false);
        request.setRequestHeader('Content-Type', 'application/json');
        request.send(JSON.stringify(payload));
        location.reload();
    });

    $("#undeleteButton").click(function() {
        var $inputs = $('#undeleteForm :input');
        var itemId = "";
        $inputs.each(function() {
            if (this.name == "itemIndex") {
                itemId = deletedItems[parseInt($(this).val())-1].itemId
            }
        });
        var request = new XMLHttpRequest();
        request.open("POST", apiUrl + "recover/" + itemId, false);
        request.send(null);
        location.reload();
    });

    $("#insertButton").click(function() {
        var $inputs = $('#insertForm :input');
        var payload = {};
        var itemId = "";
        $inputs.each(function() {
            if (this.name == "itemIndex") {
                itemId = items[parseInt($(this).val())-1].itemId
            } else {
                payload[this.name] = $(this).val();
            }
        });
        console.log(payload);
        var request = new XMLHttpRequest();
        request.open("POST", apiUrl + itemId, false);
        request.setRequestHeader('Content-Type', 'application/json');
        request.send(JSON.stringify(payload));
        location.reload();
    });
});

