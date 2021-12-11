#!/usr/bin/node

const items = [
  { id: 1, parent: 0 }, // 0
  { id: 2, parent: 1 }, // 1
  { id: 4, parent: 0 }, // 0
  { id: 6, parent: 2 }, // 2
  { id: 7, parent: 4 }, // 1
  { id: 8, parent: 6 }, // 3
  { id: 9, parent: 6 }, // 3
  { id: 10, parent: 9 }, // 4
  { id: 11, parent: 10 }, // 5
  { id: 12, parent: 11 }, // 6
];

const getParentNode = (searchItem) => {
  const parents = items.filter((item) => item.id == searchItem.parent);

  if (parents && parents.length > 0) {
    return parents[0];
  }
  return null;
};
const findParentNodes = (item, parents) => {
  // add current item to parents group and find
  const parent = getParentNode(item);
  if (!parent) {
    return parents;
  }
  // otherwise
  return findParentNodes(parent, [...parents, parent]);
};

const main = () => {
  let itemWithLevels = items.map((item) => {
    const parentNodes = findParentNodes(item, []);
    return {
      ...item,
      // level count from 0 => level = length(parentNodes)
      level: parentNodes.length,
    };
  });
  console.error(itemWithLevels);
};

main();
