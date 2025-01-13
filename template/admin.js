function handleFrequentButtonClick(event) {
    const container = event.target.parentNode;
    const frequentIpInput = container.querySelector('.frequent_ip');

    if (frequentIpInput) {
        const ipValue = frequentIpInput.textContent;

        if (ipValue !== "") {
            const formData = new FormData();
            formData.append("ip", ipValue);

            fetch("/block", {
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

function handleBlockedButtonClick(event) {
    const container = event.target.parentNode;
    const blockedIpInput = container.querySelector('.blocked_ip');

    if (blockedIpInput) {
        const ipValue = blockedIpInput.textContent;

        if (ipValue !== "") {
            const formData = new FormData();
            formData.append("ip", ipValue);

            fetch("/unblock", {
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
    const frequentButtons = document.querySelectorAll('.frequent_button');
    frequentButtons.forEach(button => {
        button.addEventListener("click", handleFrequentButtonClick);
    });

    const blockedButtons = document.querySelectorAll('.blocked_button');
    blockedButtons.forEach(button => {
        button.addEventListener("click", handleBlockedButtonClick);
    });
}

window.onpopstate = function(event) {
    if (event.state && event.state.html) {
        document.body.innerHTML = event.state.html;
    }

    attachEventListeners()
};

document.addEventListener("DOMContentLoaded", function () {
    history.replaceState({ html: document.body.innerHTML }, '', window.location.href);

    attachEventListeners();
});
