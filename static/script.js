function sendOrder() {
    fetch("/order", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            item: "Widget",
            quantity: 3,
            price: 19.99
        })
    })
    .then(res => res.json())
    .then(data => alert("Response: " + JSON.stringify(data)))
    .catch(err => console.error(err));
}

