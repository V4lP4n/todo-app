//class of tasks list
class List{
    constructor (name, id){
        this.name = name
        this.id = id
        this.childs = new Array
    }
    //generate html of obj
    genSelf(){
        let tb = document.createElement('table')
        let tr = document.createElement('tr')
        let td = document.createElement('td')
        let empty_td = document.createElement('td')
        td.innerHTML+=this.name
        tr.appendChild(td)
        tr.appendChild(empty_td)
        tb.appendChild(tr)
        this.childs.forEach(c => {
            tb.appendChild(c.genSelf())
        });

        this.html = tb
        return tb
    
    }    
    // pushing tasks
    addChild(child){
        this.childs.push(child)
    }
}
// class of tasks
class Task{
    constructor (id, data, status, list_id){
        this.id = id
        this.data = data
        this.status = status
        this.list_id = list_id
        
    }
    //generate html of rows
    genSelf(){
        let row = document.createElement('tr')
        let tdat = document.createElement('td')
        let tact = document.createElement('td')
        let img = document.createElement('img')

        tdat.innerHTML += this.data

        img.src="/images/gear.svg"
        img.className="button"+this.id 

        tact.appendChild(img)
        
        row.appendChild(tdat)
        row.appendChild(tact)
        this.html = row
        return row
    }

}

export {List, Task}