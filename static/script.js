document
  .getElementById("additionForm")
  .addEventListener("submit", async function (event) {
    event.preventDefault();
    const a = document.getElementById("a").value;
    const b = document.getElementById("b").value;
    const resultDiv = document.querySelector(".result");
    resultDiv.textContent = "Loading...";
    resultDiv.className = "loading";

    try {
      const res = await fetch(`http://localhost:8080/add-form`, {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
        body: `a=${encodeURIComponent(a)}&b=${encodeURIComponent(b)}`,
      });

      const response = await fetch(
        `http://localhost:8080/add-form-with-parse`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ a: parseFloat(a), b: parseFloat(b) }),
        }
      );
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      const data = await response.json();
      resultDiv.textContent = `Result: ${data.result}`;
      resultDiv.className = "";
    } catch (error) {
      resultDiv.textContent = `Error: ${error.message}`;
      resultDiv.className = "error";
    }
  });
