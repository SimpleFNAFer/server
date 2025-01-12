document.addEventListener("DOMContentLoaded", function () {

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
                    credentials: "include"
                })
            }
        }
    }

    const loginButton = document.querySelectorAll('.input_button');
    loginButton.forEach(button => {
        button.addEventListener("click", handleLoginButtonClick);
    });
});
