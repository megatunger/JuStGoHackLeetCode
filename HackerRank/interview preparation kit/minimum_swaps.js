function swap(arr, x, y) {
  let r = arr[x];
  arr[x] = arr[y];
  arr[y] = r;
  return arr;
}

function checkAsc(arr) {
  let result = true;
  for (let i=0; i<arr.length; i++) {
    if (arr[i+1] < arr[i]) {
      result = false;
      break;
    }
  }
  return result
}

function minimumSwaps(arr) {
  let newArr = arr;
  let steps = 0;
  while (!checkAsc(arr)) { 
    for (let i=0; i<arr.length; i++) {
      if (newArr[i] !== i+1) {
        newArr = swap(newArr, i, newArr[i]-1)
        steps++
        break
      }
    }
  }
  return steps
}

console.log('result', minimumSwaps([1,3,5,2,4,6,7]))