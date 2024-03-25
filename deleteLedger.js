const axios = require("axios");
const { faker } = require("@faker-js/faker");
const _ = require("lodash");

async function initializeLedger() {
  try {
    const payloads = [];

    for (let i = 1; i <= 40; i++) {
      const payload = {
        "@assetType": "project",
        project: `0${i}`,
      };

      payloads.push(payload);
    }

    const response = await axios.post(
      "http://localhost:80/api/invoke/deleteAsset",
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
