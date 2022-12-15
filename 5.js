import fs from "fs";

const unprocessedCommands = fs
  .readFileSync("input5.txt", { encoding: "utf8" })
  .split("\n");

const data = [
  ["L", "N", "W", "T", "D"],
  ["C", "P", "H"],
  ["W", "P", "H", "N", "D", "G", "M", "J"],
  ["C", "W", "S", "N", "T", "Q", "L"],
  ["P", "H", "C", "N"],
  ["T", "H", "N", "D", "M", "W", "Q", "B"],
  ["M", "B", "R", "J", "G", "S", "L"],
  ["Z", "N", "W", "G", "V", "B", "R", "T"],
  ["W", "G", "D", "N", "P", "L"],
];

const testData = [["Z", "N"], ["M", "C", "D"], ["P"]];

const testCommands = [
  "move 1 from 2 to 1",
  "move 3 from 1 to 3",
  "move 2 from 2 to 1",
  "move 1 from 1 to 2",
];

const getCommandValues = (command) => {
  if (!command) return null;

  const pattern = /\d+/g;

  const result = command.match(pattern).map((item) => parseInt(item));

  return result;
};

const applyCommand = (commandValues) => {
  const [numberOfBlocks, from, dest] = commandValues;

  const blocksTaken = data[from - 1].slice(
    data[from - 1].length - numberOfBlocks,
    data[from - 1].length
  );

  data[from - 1] = data[from - 1].slice(
    0,
    data[from - 1].length - numberOfBlocks
  );
  //task 2 - don't use reverse on blocksTaken
  data[dest - 1] = [...data[dest - 1], ...blocksTaken];

  return data;
};

unprocessedCommands.forEach((item) => {
  applyCommand(getCommandValues(item));
});

console.log(data);
