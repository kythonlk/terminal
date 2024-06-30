const userInput = document.getElementById("user-prompt");
const submitButton = document.getElementById("submit-button");
const responseDiv = document.getElementById("result");
const loadingIndicator = document.createElement("div");

loadingIndicator.classList.add("loading");

submitButton.addEventListener("click", async () => {
  if (!userInput) {
    console.error("Error: User input element not found. Please check the ID.");
    return;
  }

  const drugName = userInput.value;
  const prompt = "give drug leaflet for pharmacy student to study this drug: " + drugName;

  responseDiv.appendChild(loadingIndicator);
  const response = await makeRequestToAiApi(prompt);

  responseDiv.removeChild(loadingIndicator);

  if (response) {
    const htmlContent = marked.parse(response);
    responseDiv.innerHTML = htmlContent;
    await saveDrugToDatabase(drugName, response);
  } else {
    responseDiv.textContent = "No response received from AI API.";
  }
  console.log(response);
});

async function makeRequestToAiApi(prompt) {
  try {
    const session = await ai.createTextSession();
    const api = await session;
    const response = await api.prompt(prompt);
    return response;
  } catch (error) {
    console.error("Error:", error);
  }
}

async function saveDrugToDatabase(name, result) {
  try {
    const response = await fetch("/foo", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name: name, result: result }),
    });

    if (!response.ok) {
      throw new Error("Failed to save data to database");
    }

    console.log("Data saved successfully");
  } catch (error) {
    console.error("Error:", error);
  }
}


