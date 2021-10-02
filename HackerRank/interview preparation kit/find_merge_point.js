function findMergeNode(headA, headB) {
  let arr = []
  while (a || b) {
     if (a && arr.find(n => n === a)) return a.data;
     else if(a) {
         arr.push(a)
         a = a.next;
     }
     if (b && arr.find(n => n === b)) return b.data;
     else if (b) {
         arr.push(b);
         b = b.next;
     }
 }
}