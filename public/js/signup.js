let currentUserLink = document.querySelector("#currentUser")
let signup = document.querySelector("#signup")

currentUserLink.addEventListener('click', e => {
    // don't follow the link
    e.preventDefault()

    if (signup.childElementCount > 1) {
        // an error box is already there
    } else {
        let error = document.createElement('p')
        signup.prepend(error)
        error.textContent = "You must sign up and/or login before accessing the Current User Page"
        error.classList.add("bg-danger")
        error.classList.add("text-black")
        error.classList.add("p-2")
        error.classList.add("bg-opacity-50")
        error.classList.add("text-center")
    }
})
