
/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     int data;
 *     SinglyLinkedListNode next;
 * }
 *
 */

function insertNodeAtPosition(llist, data, position) {
  let i=0;
  let _prev = null;
  let _next = llist
  while (true) {
    if (position === i) {
      let insertLL = {
        data,
        next: _next,
      }
      _prev.next = insertLL
      return llist;
    }
    _prev = _next;
    _next = _next.next;
    i++;
  }
}

const ll = {data: 11, next: {data: 9, next: {data: 19, next: {data: 10, next: {data: 4, next: null}}}}}
console.log(insertNodeAtPosition(ll, 20, 3))