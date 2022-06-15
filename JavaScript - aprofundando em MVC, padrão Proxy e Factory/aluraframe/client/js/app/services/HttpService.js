class HttpService{

    get(url){

        return new Promise((resolve, reject)=>{

            let xhr = new XMLHttpRequest()
            xhr.open("GET", url)
            xhr.onreadystatechange = () => {
                if(xhr.readyState == 4){
    
                    if(xhr.status == 200){
    
                        resolve(JSON.parse(xhr.responseText))
    
                    }else{
                        reject(`Erro ao importar as negociações | ${xhr.status} | ${xhr.responseText}`)
                    }
                } 
            }
            xhr.send()
        })
    }
}