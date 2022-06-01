function tocaSom(idElementoAudio){
    elemento = document.querySelector(idElementoAudio);
    if(elemento != null && elemento.localName === 'audio'){
        elemento.play();
    }
    else{
        console.log("Error 404: Sound not found");
    }
};

const teclas = document.querySelectorAll('.tecla');

for(let i = 0; i < teclas.length; i++){

    const tecla = teclas[i];
    const idAudio = `#som_${tecla.classList[1]}`;

    tecla.onclick = function(){
        tocaSom(idAudio);
    };

    tecla.onkeydown = function(evento){
        console.log(evento);
        if(evento.code === 'Space' || evento.code === "Enter"){
            tecla.classList.add('ativa');
        }
    }

    tecla.onkeyup = function(){
        tecla.classList.remove('ativa');
    }
}