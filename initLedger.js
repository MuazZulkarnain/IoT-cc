const axios = require("axios");
const { faker } = require("@faker-js/faker");
const _ = require("lodash");

async function initializeLedger() {
  try {
    const payloads = [];

    for (let i = 1; i <= 1000000; i++) {
      let project = faker.company.name(); // Define project variable
      let date = new Date();
      let localtime = new Date(date.getTime() + 8 * 60 * 60000); // Adjust for Malaysia timezone (UTC+8)

      const payload = {
        "@assetType": "project",
        project: project,
        amount: _.random(50, 200),
        claimAmount: _.random(100, 300),
        lastAction: `Initialized ${project} at ${localtime.toLocaleString(
          "en-US",
          {
            timeZone: "Asia/Kuala_Lumpur", // Specify the timezone
          }
        )}`,
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
