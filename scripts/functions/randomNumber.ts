import crypto from "crypto";

export function getRandomNumber(digits: number) {
  const min = Math.pow(10, digits - 1);
  const max = Math.pow(10, digits) - 1;
  return Math.floor(min + Math.random() * (max - min + 1));
}

export function randomUint256() {
  const randomBytes = crypto.randomBytes(32); // Generate 32 random bytes
  let hex = randomBytes.toString("hex"); // Convert bytes to hexadecimal
  return BigInt("0x" + hex); // Convert hexadecimal to BigInt
}
