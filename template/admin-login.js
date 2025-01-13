function handleLoginButtonClick(event) {
    const container = event.target.parentNode;
    const login = container.querySelector('#login');
    const password = container.querySelector('#password');

    if (login && password) {
        const loginValue = login.value;
        const passwordValue = password.value;

        if (loginValue !== "" && passwordValue !== "") {
            const formData = new FormData();
            formData.append("login", loginValue);
            formData.append("password", passwordValue);

            fetch("/admin-login", {
                method: "POST",
                body: formData,
                credentials: "include",
                redirect: "follow"
            })
                .then(response => {
                        if (response.redirected) {
                            window.location.replace(response.url);
                        }
                        else {
                            return response.text()
                        }
                    }
                )
                .then((html) => {
                    if (html) {
                        const newUrl = window.location.href;
                        history.pushState({ html: html }, '', newUrl);
                        document.body.innerHTML = html;
                    }
                })
        }
    }
}

function attachEventListeners() {
    const loginButton = document.querySelectorAll('.input_button');
    loginButton.forEach(button => {
        button.addEventListener("click", handleLoginButtonClick);
    });
}

document.addEventListener("DOMContentLoaded", function () {
    history.replaceState({ html: document.body.innerHTML }, '', window.location.href);

    attachEventListeners();
});

window.onpopstate = function(event) {
    if (event.state && event.state.html) {
        document.body.innerHTML = event.state.html;
    }

    attachEventListeners()
};
