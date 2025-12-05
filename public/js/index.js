document.addEventListener("DOMContentLoaded", () => {
    const btn = document.getElementById("helloBtn");
    const responseDiv = document.getElementById("response");

    btn.addEventListener("click", async () => {
        try {
            const res = await fetch("/api/hello");
            const data = await res.json();
            responseDiv.innerHTML = `<p><strong>API says:</strong> ${data.message}</p>`;
        } catch (err) {
            responseDiv.innerHTML = `<p style="color:red;">Error: ${err}</p>`;
        }
    });
});