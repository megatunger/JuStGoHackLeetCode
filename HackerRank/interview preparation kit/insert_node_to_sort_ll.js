const node1 = {
  data: 1,
  prev: null,
}
const node2 = {
  data: 3,
  prev: node1,
}
const node3 = {
  data: 2,
  prev: node2,
}
const node4 = {
  data: 3,
  prev: node3,
}
const node5 = {
  data: 4,
  prev: node4,
  next: null
}
const ll = {
  ...node1,
  next: {
    ...node2,
    next: {
      ...node3,
      next: {
        ...node4,
        next: node5
      }
    }
  },
}


function sortedInsert(head, data) {
  if(head == null) {
      return new DoublyLinkedListNode(data);
  }
  var current = head;
  if(data < current.data) {
      var newNode = new DoublyLinkedListNode(data);
      newNode.next = current;
      current.prev = newNode;
      return newNode;        
  }
  
  while(current != null) {
      if(data >= current.data) {
          if(current.next == null) {
              var newNode = new DoublyLinkedListNode(data);
              current.next = newNode;
              newNode.prev = current;
              
              return head;
          }
          
          if(current.next.data >= data) {
              var newNode = new DoublyLinkedListNode(data);
              newNode.next = current.next;
              newNode.prev = current;
              current.next = newNode;
              newNode.next.prev = newNode;
              
              return head;
          }
      }
      
      current = current.next;
  }
  
  return head;
  
}

console.log(sortedInsert(ll, 1))