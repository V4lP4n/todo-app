

let host = "http://localhost"
let list = "/api/list"


//returns full list of taskbars via json
async function get_json_list(){
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
//takes task info and returns html row for it
    function prep_row(num, data, status){
    let tr = document.createElement('tr')
    let data_td = document.createElement('td')    
    let actions_td = document.createElement('td') 
    let actions = document.createElement('img')
    actions.src="/images/gear.svg"
    actions.className="button"+num
    tr.className="todo_tr "+status
    data_td.className="todo_td"    
    data_td.innerHTML+= data
    actions_td.appendChild(actions)
    tr.appendChild(data_td)
    tr.appendChild(actions_td)

    menu = document.querySelector("#actions")
    
    actions.onclick=function(){
        menu.className="visible"
    }
        
    return tr
}

//adding table with taskbars in html
async function load_table(){
    let app = document.querySelector('#app')
    let table = app.querySelector('table')
    let tbody = table.querySelector('tbody')
    let data = await get_json_list()

    for ( let i=0; i<data.length; i++){
    let row = prep_row(data[i].id, data[i].data, data[i].status)
    tbody.appendChild(row)

    }
}
load_table()