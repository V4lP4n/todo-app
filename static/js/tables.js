class TB{
    constructor (name){
        this.name = name
    }
    generate (){
        let tb = document.createElement('table')
        let tr = document.createElement('tr')
        let td = document.createElement('td')
        tr.innerHTML+=this.name
        tr.appendChild(td)
        tb.appendChild(tr)
        return tb

    }
}

export {TB}