import fs from "fs";

const input = fs.readFileSync("input4.txt", { encoding: "utf8" }).split("\n");

const processInterval = (interval) =>
  interval.split("-").map((item) => parseInt(item));

const isIncluded = (pair) => {
  const [firstInterval, secondInterval] = pair.split(",");

  const [firstBottom, firstTop] = processInterval(firstInterval);
  const [secondBottom, secondTop] = processInterval(secondInterval);

  const isCovered =
    (firstBottom <= secondBottom && firstTop >= secondTop) ||
    (secondBottom <= firstBottom && secondTop >= firstTop);

  return isCovered;
};

const areSeparated = (pair) => {
  const [firstInterval, secondInterval] = pair.split(",");

  const [firstBottom, firstTop] = processInterval(firstInterval);
  const [secondBottom, secondTop] = processInterval(secondInterval);

  const firstArr = new Array(firstTop - firstBottom + 1)
    .fill(0)
    .map((_, i) => firstBottom + i);
  const secondArr = new Array(secondTop - secondBottom + 1)
    .fill(0)
    .map((_, i) => secondBottom + i);

  const combined = [...firstArr, ...secondArr];

  const uniqueValues = new Set(combined);

  return combined.length === uniqueValues.size;
};

const result = input.reduce(
  (acc, item) => (areSeparated(item) ? acc : acc + 1),
  0
);
// const result2 = areSeparated('8-9,5-7')

//

console.log(input.length);
console.log(result);
// console.log(result2)
