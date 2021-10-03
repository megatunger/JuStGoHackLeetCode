class DoublyLinkedList {
  constructor() {
    this.head = null;
    this.tail = null;
    this.length = 0;
  }

  insert(value) {
    this.length++;
    let newNode = createNode(value);
  
    if (this.tail) {
      // list is not empty
      this.tail.next = newNode;
      newNode.previous = this.tail;
      this.tail = newNode;
      return newNode;
    }
  
    this.head = this.tail = newNode;
    return newNode;
  }
}

function reverse(llist) {
  let current = llist
  let arr = []
  while (current) {
    arr.push(current.data)
    current = current.next
  }
  arr = arr.reverse()

  var head = new DoublyLinkedList()
     
  arr.forEach(v => {
    head.insert(v)
  })
  return v.head
}
