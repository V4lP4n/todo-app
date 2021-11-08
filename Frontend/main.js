
document.addEventListener("DOMContentLoaded", ready)

let host = "http://localhost"
let list = "/api/list"



async function get_list(){
let url = host + list;
let list_data
let response = await fetch(url);
if (response.ok){
    list_data = await response.json();
    

} else{
    console.error(response.status)
}
return list_data
}

function prep_row(num, data, status){
 let tr = document.createElement('tr')
 let id_td = document.createElement('td')
 let data_td = document.createElement('td')
 let status_td = document.createElement('td')
 
 tr.className="todo_tr_"+status
 id_td.className="todo_td"
 data_td.className="todo_td"

 
 id_td.innerHTML=num
 id_td.innerHTML+= data

tr.appendChild(id_td)
    tr.appendChild(data_td)
return tr
}


async function ready(){

let app = document.querySelector("#app")
let data = await get_list()
for ( let i=0; i<data.length; i++){
 let row = prep_row(data[i].id, data[i].data, data[i].status)
 app.appendChild(row)
}
}