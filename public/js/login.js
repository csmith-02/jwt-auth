let currentUserLink = document.querySelector("#currentUser")
let login = document.querySelector("#login")

currentUserLink.addEventListener('click', e => {
    // don't follow the link
    e.preventDefault()

    if (login.childElementCount > 1) {
        // an error box is already there
    } else {
        let error = document.createElement('p')
        login.prepend(error)
        error.textContent = "You must login before accessing the Current User Page"
        error.classList.add("bg-danger")
        error.classList.add("text-black")
        error.classList.add("p-2")
        error.classList.add("bg-opacity-50")
        error.classList.add("text-center")
    }
})

