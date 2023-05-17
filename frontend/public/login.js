$(document).ready(function() {
    $('form').submit(function(event) {
        event.preventDefault();

        var formData = $('form').serializeArray();
        var jsonData = {};
        formData.forEach(field => {
            jsonData[field.name] = field.value
        })
        console.log(JSON.stringify(jsonData))
        $.ajax({
            type: 'POST',
            url: 'http://localhost:8080/auth',
            data: JSON.stringify(jsonData),
            contentType: 'application/json',
            success: function(data) {
                console.log('Success: ', data);
            },
            error: function(xhr, textStatus, error) {
                console.log('Error: ', error);
            }
        }).done(function(json) {
            console.log("Receive response from server")
            console.log(json)
        }).always(function(xhr, status) {
            console.log("The request is complete")
            console.log(xhr, status)
        })
    })
})
