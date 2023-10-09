document.getElementById("input")
    .addEventListener("submit", async e => {
        e.preventDefault()

        const data = new FormData(e.target)
        const value = Object.fromEntries(data.entries())

        const answer = await fetch(`/orders/${value.orderuid}`,
            {
                method: "GET",
            })
            .then(res => res.json())
            .then(body => JSON.stringify(body, null, 2))

        document.getElementById("data").innerText = answer
    })