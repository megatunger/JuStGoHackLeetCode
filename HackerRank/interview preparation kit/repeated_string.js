function repeatedString(s, n) {
  let numberOfChars = 0
  for (let i=0; i<s.length; i++) {
      if (s[i] === 'a') {
          numberOfChars++
      }
  }
  const k = Math.floor((n/s.length))
  let leftChars = 0
  if (n%s.length !== 0){
      for (let i=0; i<=n%s.length-1; i++) {
          if (s[i] === 'a') {
              leftChars++
          }
      }
  }
  return Math.floor((numberOfChars * k)+leftChars)
}