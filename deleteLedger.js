const axios = require("axios");

async function searchProjects() {
  const url = "http://localhost/api/query/search";
  const requestData = {
    query: {
      selector: {
        "@assetType": "project",
      },
    },
  };

  try {
    const response = await axios.post(url, requestData);
    const keysArray = response.data.result.map((project) => project["@key"]);
    return keysArray;
  } catch (error) {
    throw error.response ? error.response.data : error.message;
  }
}

async function deleteLedger() {
  try {
    const keysArray = await searchProjects();

    // Construct payloads for each key
    const payloads = keysArray.map((key) => ({
      key: {
        "@assetType": "project",
        "@key": key,
      },
    }));

    // Make DELETE requests for each payload
    await Promise.all(
      payloads.map(async (payload) => {
        const response = await axios.delete(
          "http://localhost/api/invoke/deleteAsset",
          {
            data: payload,
            headers: {
              "Content-Type": "application/json",
              accept: "*/*",
            },
          }
        );
        console.log("API response:", response.data);
      })
    );
  } catch (error) {
    console.error("Error calling API:", error.message);
  }
}

deleteLedger();
