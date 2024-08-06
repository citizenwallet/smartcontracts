import readline from "readline";
import { ethers, upgrades, run } from "hardhat";
import "@nomicfoundation/hardhat-toolbox";
import { config } from "dotenv";

/**
 * @notice This script deploys the following contracts:
 * - Profile (unless process.env.PROFILE_ADDR is set, or a custom address is provided at the prompt)
 * - Paymaster (unless process.env.PAYMASTER_ADDR is set)
 * - TokenEntryPoint (unless process.env.TOKEN_ENTRYPOINT_ADDR is set)
 * - AccountFactory (unless process.env.ACCOUNT_FACTORY_ADDR is set)
 */

async function main() {
  const whiteListedAddresses = [];
  const { TOKEN_ADDR } = process.env;
  if (TOKEN_ADDR && TOKEN_ADDR.startsWith("0x")) {
    whiteListedAddresses.push(TOKEN_ADDR);
  }

  console.log("⚙️ Running...\n");
  config();

  if (
    process.env.ENTRYPOINT_ADDR === undefined ||
    process.env.ENTRYPOINT_ADDR === ""
  ) {
    throw Error("ENTRYPOINT_ADDR is not set");
  }

  const profileResponse = async () => {
    return new Promise<string | undefined>((resolve) => {
      const rl = readline.createInterface({
        input: process.stdin,
        output: process.stdout,
      });

      rl.question(
        "Custom profile contract address (leave empty if you want to create a new one): ",
        (answer) => {
          if (
            answer === "" ||
            answer === undefined ||
            !answer.startsWith("0x")
          ) {
            resolve(undefined);
          }

          resolve(answer);
          rl.close();
        }
      );
    });
  };

  const customProfile = process.env.PROFILE_ADDR
    ? process.env.PROFILE_ADDR
    : await profileResponse();

  console.log("\n");

  if (whiteListedAddresses.length === 0) {
    if (process.env.TOKEN_ADDR) {
      whiteListedAddresses.push(process.env.TOKEN_ADDR);
    } else {
      const addressResponse = new Promise<string>((resolve) => {
        const rl = readline.createInterface({
          input: process.stdin,
          output: process.stdout,
        });

        rl.question("Community Token Address: ", (answer) => {
          resolve(answer.trim());
          rl.close();
        });
      });
      const contractAddress = await addressResponse;
      whiteListedAddresses.push(contractAddress);
    }
  }

  console.log("\n");

  const sponsor = ethers.Wallet.createRandom();

  let profileAddress: string = customProfile ?? "";
  if (customProfile === undefined) {
    console.log("⚙️ deploying Profile...");

    const profileFactory = await ethers.getContractFactory("Profile");

    const profile = await upgrades.deployProxy(
      profileFactory,
      [process.env.DEPLOYER_ADDRESS],
      {
        kind: "uups",
        initializer: "initialize",
        timeout: 999999,
      }
    );

    console.log("🚀 request sent...");
    await profile.deployed();

    console.log(`✅ Profile deployed to ${profile.address}`);

    profileAddress = profile.address;

    console.log("\n");
  }

  if (profileAddress === "") {
    throw Error("Something went wrong, profile address is empty");
  }

  whiteListedAddresses.push(profileAddress);

  if (whiteListedAddresses.length > 0) {
    console.log(`whitelisting contracts:`);
    whiteListedAddresses.forEach((addr) => console.log(addr));
    console.log("\n");
  }

  let paymaster;
  if (process.env.PAYMASTER_ADDR) {
    paymaster = {
      address: process.env.PAYMASTER_ADDR,
    };
    console.log(
      `✅ Using Paymaster contract deployed at: ${paymaster.address}`
    );
  } else {
    console.log("⚙️ deploying Paymaster...");

    const paymasterFactory = await ethers.getContractFactory("Paymaster");

    paymaster = await upgrades.deployProxy(
      paymasterFactory,
      [sponsor.address],
      {
        kind: "uups",
        initializer: "initialize",
        timeout: 999999,
      }
    );

    console.log("🚀 request sent...");

    await paymaster.deployed();

    console.log(`✅ Paymaster contract deployed to: ${paymaster.address}`);
  }

  console.log("\n");

  let tokenEntryPoint;
  if (process.env.TOKEN_ENTRYPOINT_ADDR) {
    tokenEntryPoint = { address: process.env.TOKEN_ENTRYPOINT_ADDR };
    console.log(
      `✅ Using TokenEntryPoint deployed at: ${tokenEntryPoint.address}`
    );
  } else {
    console.log("⚙️ deploying TokenEntryPoint...");

    const tokenEntryPointFactory = await ethers.getContractFactory(
      "TokenEntryPoint"
    );
    tokenEntryPoint = await upgrades.deployProxy(
      tokenEntryPointFactory,
      [
        sponsor.address,
        paymaster.address,
        [...whiteListedAddresses, profileAddress],
      ],
      {
        kind: "uups",
        initializer: "initialize",
        constructorArgs: [process.env.ENTRYPOINT_ADDR],
        timeout: 999999,
      }
    );

    console.log("🚀 request sent...");

    await tokenEntryPoint.deployed();

    console.log(`✅ TokenEntryPoint deployed to: ${tokenEntryPoint.address}`);
  }

  console.log("\n");

  let accFactory;
  if (process.env.ACCOUNT_FACTORY_ADDR) {
    accFactory = { address: process.env.ACCOUNT_FACTORY_ADDR };
    console.log(`✅ Account Factory deployed at: ${accFactory.address}`);
  } else {
    console.log("⚙️ deploying AccountFactory...");

    const accountFactory = await ethers.getContractFactory("AccountFactory");
    accFactory = await upgrades.deployProxy(accountFactory, [], {
      kind: "uups",
      initializer: "initialize",
      constructorArgs: [process.env.ENTRYPOINT_ADDR, tokenEntryPoint.address],
      timeout: 999999,
    });

    console.log("🚀 request sent...");

    await accFactory.deployed();

    console.log(`✅ Account Factory deployed to: ${accFactory.address}`);
  }

  console.log("\n");

  console.log("🧐 verifying...");

  // wait for etherscan to index the contract
  await new Promise((resolve) => setTimeout(resolve, 10000));

  try {
    await run("verify:verify", {
      address: paymaster.address,
      constructorArguments: [],
    });

    console.log("verified!");
  } catch (error: any) {
    if (
      error.message.includes(
        "does not seem to be verified. Please verify and publish the contract source before proceeding with this proxy verification."
      )
    ) {
      console.log("\n");
      console.log(
        "⚠️ We were unable to verify the contract fully. This can happen when deploying the same contract with a proxy multiple times.\n"
      );
      console.log(
        "You might need to go the contract's code page on Etherscan and click 'More Options' -> 'Is this a proxy?'"
      );
      console.log("\n");
    } else {
      console.log(error);
    }
  }

  try {
    await run("verify:verify", {
      address: tokenEntryPoint.address,
      constructorArguments: [process.env.ENTRYPOINT_ADDR],
    });

    console.log("verified!");
  } catch (error: any) {
    if (
      error.message.includes(
        "does not seem to be verified. Please verify and publish the contract source before proceeding with this proxy verification."
      )
    ) {
      console.log("\n");
      console.log(
        "⚠️ We were unable to verify the contract fully. This can happen when deploying the same contract with a proxy multiple times.\n"
      );
      console.log(
        "You might need to go the contract's code page on Etherscan and click 'More Options' -> 'Is this a proxy?'"
      );
      console.log("\n");
    } else {
      console.log(error);
    }
  }

  try {
    await run("verify:verify", {
      address: accFactory.address,
      constructorArguments: [
        process.env.ENTRYPOINT_ADDR,
        tokenEntryPoint.address,
      ],
    });

    console.log("verified!");
  } catch (error: any) {
    if (
      error.message.includes(
        "does not seem to be verified. Please verify and publish the contract source before proceeding with this proxy verification."
      )
    ) {
      console.log("\n");
      console.log(
        "⚠️ We were unable to verify the contract fully. This can happen when deploying the same contract with a proxy multiple times.\n"
      );
      console.log(
        "You might need to go the contract's code page on Etherscan and click 'More Options' -> 'Is this a proxy?'"
      );
      console.log("\n");
    } else {
      console.log(error);
    }
  }

  try {
    await run("verify:verify", {
      address: profileAddress,
      constructorArguments: [],
    });

    console.log("verified!");
  } catch (error: any) {
    if (
      error.message.includes(
        "does not seem to be verified. Please verify and publish the contract source before proceeding with this proxy verification."
      )
    ) {
      console.log("\n");
      console.log(
        "⚠️ We were unable to verify the contract fully. This can happen when deploying the same contract with a proxy multiple times.\n"
      );
      console.log(
        "You might need to go the contract's code page on Etherscan and click 'More Options' -> 'Is this a proxy?'"
      );
      console.log("\n");
    } else {
      console.log(error);
    }
  }

  console.log("\n");
  console.log("\n");
  console.log("\n");

  console.log("*************************************");
  console.log("DEPLOYMENT COMPLETE 🎉");
  console.log(" ");
  console.log(" ");
  console.log("Community token address: ", TOKEN_ADDR);
  console.log(" ");
  console.log("Paymaster contract address: ", paymaster.address);
  console.log(
    "Paymaster sponsor address (EOA to top up to sponsor gas fees): ",
    sponsor.address
  );
  console.log(
    "Paymaster sponsor private key: ",
    sponsor.privateKey.replace("0x", "")
  );
  console.log(" ");
  if (whiteListedAddresses.length > 0) {
    console.log(`Whitelisted contracts:`);
    whiteListedAddresses.forEach((addr) => console.log(addr));
    console.log(profileAddress);
    console.log("\n");
  }
  console.log(" ");
  console.log("Token Entry Point: ", tokenEntryPoint.address);
  console.log("Account Factory: ", accFactory.address);
  console.log("Profile: ", profileAddress);
  console.log("*************************************");

  console.log("\n");
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
