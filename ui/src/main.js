const baseUrl = "/s"

document.addEventListener("DOMContentLoaded", function main() {
    const urlForm = document.getElementById("urlform");
    urlForm.addEventListener("submit", async (submitEvent) => {
        submitEvent.preventDefault();

        const url = document.getElementById("url-input").value;
        const isHumanReadable = document.getElementById("human-readable-input").checked;
        const data = {
            url,
            format: isHumanReadable ? "mnemonic" : "sid"
        };

        const response = await fetch(
            `${baseUrl}/`,
            data
        );
    })
});
