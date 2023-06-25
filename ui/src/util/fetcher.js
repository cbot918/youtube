export {fetchGet, fetchPost}

const fetchGet = (url)=>{
  fetch(url,{
    method: "get",
    headers: {
      "Content-Type": "application/json"
    }
  })
  .then(res=>res.text())
  .then((data)=>{
    console.log(data)
  })
  .catch(err=>{ console.log(err) })
}

const fetchPost = (url, data)=>{
  fetch(url,{
    method: "post",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(data)
  })
  .then(res=>res.text())
  .then((data)=>{
    console.log(data)
  })
  .catch(err=>{ console.log(err) })
}

