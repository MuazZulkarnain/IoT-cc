const axios = require("axios");
const { faker } = require("@faker-js/faker");
const _ = require("lodash");

async function initializeLedger() {
  try {
    const payloads = [];

    for (let i = 1; i <= 40; i++) {
      const payload = {
        "@assetType": "project",
        project: faker.company.name(),
        amount: _.random(50, 200), // Adjust the range as needed
        claimAmount: _.random(100, 300), // Adjust the range as needed
      };

      payloads.push(payload);
    }

    const response = await axios.post(
      "http://localhost:80/api/invoke/createAsset",
      {
        asset: payloads,
      },
      {
        headers: {
          "Content-Type": "application/json",
          "cache-control": "no-cache",
        },
      }
    );

    console.log("API response:", response.data);
  } catch (error) {
    console.error("Error calling API:", error.message);
  }
}

initializeLedger();
