document.getElementById('getPacks').addEventListener('click', function() {
    fetch('/available_packs')
        .then(response => {
            if (!response.ok) {
                return response.json().then(err => { throw new Error(err.message); });
            }

            return response.json();
        })
        .then(data => {
            const obj = JSON.parse(JSON.stringify(data, null, 2));
            let text = "<table>"
            for (let x in obj) {
                text += "<tr><td>" + obj[x] + "</td></tr>";
            }
            text += "</table>"
            document.getElementById("packsOutput").innerHTML = text;
        })
        .catch(error => {
            alert(error.message);
        });
});

document.getElementById('packItems').addEventListener('click', function() {
    const items = document.getElementById('packItemsInput').value;
    if (items == "") {
        alert('please enter a valid number of items');
        return;
    }

    fetch('/pack_items/' + items)
        .then(response => {
            if (!response.ok) {
                return response.json().then(err => { throw new Error(err.message); });
            }

            return response.json();
        })
        .then(data => {
            const obj = JSON.parse(JSON.stringify(data, null, 2));
            let text = "<table>"
            Object.keys(obj).forEach(function(key) {
                // console.log('Key : ' + key + ', Value : ' + data[key])
                text += "<tr><td>" + key + "</td><td>" + data[key] + "</td></tr>";
            })
            // for (let x in obj) {
            //     text += "<tr><td>" + obj + "</td></tr>";
            // }
            text += "</table>"
            document.getElementById("packItemsOutput").innerHTML = text;
        })
        .catch(error => {
            alert(error.message);
        });
});

document.getElementById('addPack').addEventListener('click', function() {
    const items = document.getElementById('addPackInput').value;
    if (items == "") {
        alert('please enter a valid number of items');
        return;
    }

    fetch('/pack/' + items, { method: 'POST' })
        .then(response => {
            if (!response.ok) {
                return response.json().then(err => { throw new Error(err.message); });
            }

            return null;
        })
        .then(() => {
            alert('pack size added');
        })
        .catch(error => {
            alert(error.message);
        });
});

document.getElementById('deletePack').addEventListener('click', function() {
    const items = document.getElementById('deletePackInput').value;
    if (items == "") {
        alert('please enter a valid number of items');
        return;
    }

    fetch('/pack/' + items, { method: 'DELETE' })
        .then(response => {
            if (!response.ok) {
                return response.json().then(err => { throw new Error(err.message); });
            }

            return null;
        })
        .then(() => {
            alert('pack size removed');
        })
        .catch(error => {
            alert(error.message);
        });
});
