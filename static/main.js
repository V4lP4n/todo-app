import {List, Task} from './js/entities.js'

const host = "http://localhost"
const lists = "/api/lists"
const tasks = "/api/tasks"


//returns array of todo lists
async function get_lists(){
    let url = host + lists;
    let list_data
    let response = await fetch(url);
    if (response.ok){
        list_data = await response.json();
    } else{
        console.error(response.status)
    }
    return list_data
}
//returns array of todo tasks
async function get_tasks(){
    let url = host + tasks;
    let list_data
    let response = await fetch(url);
    if (response.ok){
        list_data = await response.json();
    } else{
        console.error(response.status)
    }
    return list_data
}
//implement data into the page
async function render(){
    //getting json data 
    let tasks = await get_tasks()
    let lists = await get_lists()
    //prepearing vars
    let app = document.querySelector('.app')
    let obj_tasks = []
    let obj_lists = []

    //generating objects based on our json
    for (let i in  lists) {
        let o = new List(lists[i].name, lists[i].id)
        obj_lists.push(o)
    }
    for (let i in tasks){
        let o = new Task(tasks[i].id,tasks[i].data,tasks[i].status,tasks[i].list_id)
        obj_tasks.push(o)  
    }

    //pushing tasks to lists as a childs
    for (let i in obj_lists){
        let l = obj_lists[i]
        for (let j in obj_tasks){
            let t = obj_tasks[j]
            if (t.list_id == l.id){
                l.addChild(t)
            }
        } 

    }

    //rendering
    for (let i in obj_lists){
        obj_lists[i].genSelf()
        app.appendChild(obj_lists[i].html)
    }

    


}
render()








/*
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

    menu = document.querySelector(".actions")
    prev_id = menu.id
    
    actions.onclick=function(){
        menu.className="visible"
        //menu.id+=num

    }
        
    return tr
}

//adding table with taskbars in html
async function load_table(){
    let app = document.querySelector('.app')
    let table = app.querySelector('table')
    let tbody = table.querySelector('tbody')
    let data = await get_json_list()

    for ( let i=0; i<data.length; i++){
    let row = prep_row(data[i].id, data[i].data, data[i].status)
    tbody.appendChild(row)

    }   
}
load_table()

*/