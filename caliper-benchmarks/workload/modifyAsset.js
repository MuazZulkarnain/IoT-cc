"use strict";

const { WorkloadModuleBase } = require("@hyperledger/caliper-core");

class MyWorkload extends WorkloadModuleBase {
  constructor() {
    super();
  }

  async initializeWorkloadModule(
    workerIndex,
    totalWorkers,
    roundIndex,
    roundArguments,
    sutAdapter,
    sutContext
  ) {
    await super.initializeWorkloadModule(
      workerIndex,
      totalWorkers,
      roundIndex,
      roundArguments,
      sutAdapter,
      sutContext
    );

    for (let i = 0; i < this.roundArguments.assets; i++) {
      const assetID = `${this.workerIndex}${i}`;
      console.log(`Worker ${this.workerIndex}: Creating asset ${assetID}`);
      // Use the correct contractFunction and contractArguments for creating assets
      const request = {
        contractId: "IoT-cc",
        contractFunction: "createAsset",
        invokerIdentity: "User1",
        contractArguments: [
          `{"asset": [{"@assetType":"owner","name":"${assetID.toString()}"}]}`,
        ],
        readOnly: false,
      };

      await this.sutAdapter.sendRequests(request);
      // console.log("Request Object:", request);
    }
  }

  async submitTransaction() {
    const randomId = Math.floor(Math.random() * this.roundArguments.assets);
    const assetID = `${this.workerIndex}${randomId}`;
    console.log(`Worker ${this.workerIndex}: Reading asset ${assetID}`);
    // Use the correct contractFunction and contractArguments for reading assets
    const readAssetRequest = {
      contractId: "IoT-cc", // Corrected variable name
      contractFunction: "readAsset", // Corrected contract function
      invokerIdentity: "User1",
      contractArguments: [
        `{"key": {"@assetType":"owner","name":"${assetID.toString()}"}}`,
      ],
      readOnly: true,
    };

    await this.sutAdapter.sendRequests(readAssetRequest);
  }

  async cleanupWorkloadModule() {
    for (let i = 0; i < this.roundArguments.assets; i++) {
      const assetID = `${this.workerIndex}${i}`;
      console.log(`Worker ${this.workerIndex}: Deleting asset ${assetID}`);
      // Use the correct contractFunction and contractArguments for deleting assets
      const deleteAssetRequest = {
        contractId: "IoT-cc", // Corrected variable name
        contractFunction: "deleteAsset", // Replace with the correct function for deleting assets
        invokerIdentity: "User1",
        contractArguments: [
          `{"key": {"@assetType":"owner","name":"${assetID.toString()}"}}`,
        ],
        readOnly: false, // Update to false for deletion
      };

      await this.sutAdapter.sendRequests(deleteAssetRequest);
    }
  }
}

/**
 * Create a new instance of the workload module.
 * @return {WorkloadModuleBase}
 */
function createWorkloadModule() {
  return new MyWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;
