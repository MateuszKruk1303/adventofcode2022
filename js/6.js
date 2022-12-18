import fs from "fs";

const data = fs.readFileSync("./input6.txt", { encoding: "utf8" }).split("");

const getMarkerPosition = (input) => {
  const processedCharacters = [];

  for (let i = 0; i <= input.length - 1; i++) {
    // part 2: 4 -> 14
    const subsquence = input.slice(i, 14 + i);

    if (subsquence.length === new Set(subsquence).size) {
      processedCharacters.push(...subsquence);
      break;
    }

    processedCharacters.push(input[i]);
  }

  return processedCharacters;
};

console.log(getMarkerPosition(data).length);
