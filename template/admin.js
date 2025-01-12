document.addEventListener("DOMContentLoaded", function () {

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
                    credentials: "include"
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
                    credentials: "include"
                })
            }
        }
    }

    const frequentButtons = document.querySelectorAll('.frequent_button');
    frequentButtons.forEach(button => {
        button.addEventListener("click", handleFrequentButtonClick);
    });

    const blockedButtons = document.querySelectorAll('.blocked_button');
    blockedButtons.forEach(button => {
        button.addEventListener("click", handleBlockedButtonClick);
    });
});
